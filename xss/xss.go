package xss

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/Contrast-Security-OSS/go-test-bench/utils"
)

var templates = template.Must(template.ParseFiles(
	"./views/partials/safeButtons.gohtml",
	"./views/pages/xss.gohtml",
	"./views/partials/ruleInfo.gohtml",
))

func queryHandler(w http.ResponseWriter, r *http.Request, safety string) (template.HTML, bool) {
	var input string

	switch safety {
	case "safe":
		input = utils.GetUserInput(r)
		input = url.QueryEscape(input)
	case "unsafe":
		input = utils.GetUserInput(r)
	case "noop":
		return template.HTML("NOOP"), false
	default:
		log.Fatal("Error running queryHandler. No option passed")
	}
	//execute input script
	return template.HTML(input), false

}

// bufferedQueryHandler used as a handler which uses bytes.Buffer for source input ignoring the user input
func bufferedQueryHandler(w http.ResponseWriter, r *http.Request, safety string) (template.HTML, bool) {
	var buf bytes.Buffer
	buf.Write([]byte("<script>"))
	buf.WriteString("alert('buffered input - ")
	buf.WriteRune('©')
	buf.WriteString("');")
	buf.WriteString("</script")
	buf.WriteByte(byte('>'))

	var input string

	switch safety {
	case "safe":
		input = string(buf.Bytes())
		input = url.QueryEscape(input)
	case "unsafe":
		input = string(buf.Bytes())
	case "noop":
		return template.HTML("NOOP"), false
	default:
		log.Fatal("Error running bufferedQueryHandler. No option passed")
	}
	return template.HTML(input), false
}

func paramsHandler(w http.ResponseWriter, r *http.Request, safety string) (template.HTML, bool) {
	// since the attack mode as a last parameter in the query path - e.g. /unsafe, /safe, /noop
	// the user input is placed in the middle and it includes the "/" symbol so we need to combine two pieces
	// /xss/params/reflectedXss/<script>alert(1);</script>/unsafe
	// therefore we specify exact positions of the path to be considered as the user input value
	var input string

	switch safety {
	case "safe":
		input = utils.GetPathValue(r, 4, 5)
		input = url.QueryEscape(input)
	case "unsafe":
		input = utils.GetPathValue(r, 4, 5)
	case "noop":
		return template.HTML("NOOP"), false
	default:
		log.Fatal("Error running paramsHandler. No option passed")
	}

	return template.HTML(input), false
}

func bodyHandler(w http.ResponseWriter, r *http.Request, safety string) (template.HTML, bool) {
	if r.Method == http.MethodGet {
		return template.HTML("Cannot GET " + r.URL.Path), false
	}
	var input string

	switch safety {
	case "safe":
		input = utils.GetUserInput(r)
		input = url.QueryEscape(input)
	case "unsafe":
		input = utils.GetUserInput(r)
	case "noop":
		return template.HTML("NOOP"), false
	default:
		log.Fatal("Error running bodyHandler. No option passed")
	}

	return template.HTML(input), false
}

// bufferedBodyHandler used as a handler which uses bytes.Buffer for source input
func bufferedBodyHandler(w http.ResponseWriter, r *http.Request, safety string) (template.HTML, bool) {
	if r.Method == http.MethodGet {
		return template.HTML("Cannot GET " + r.URL.Path), false
	}
	var input string
	buf := bytes.NewBufferString(utils.GetUserInput(r))

	switch safety {
	case "safe":
		input = buf.String()
		input = url.QueryEscape(input)
	case "unsafe":
		input = buf.String()
	case "noop":
		return template.HTML("NOOP"), false
	default:
		log.Fatal("Error running bufferedBodyHandler. No option passed")
	}

	return template.HTML(input), false
}

func getSimpleHTTPResponse(userInput string) *http.Response {
	return &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Body:       ioutil.NopCloser(strings.NewReader(userInput)),
		Header: http.Header{
			"Content-Type": []string{"text/plain; charset=utf-8"},
			"input":        []string{userInput},
		},
	}
}

func responseHandler(w http.ResponseWriter, r *http.Request, safety string) (template.HTML, bool) {
	userInput := utils.GetUserInput(r)
	switch safety {
	case "safe":
		userInput = url.QueryEscape(userInput)
		response := getSimpleHTTPResponse(userInput)
		var buf bytes.Buffer
		response.Write(&buf)

		return template.HTML(userInput), false
	case "unsafe":
		response := getSimpleHTTPResponse(userInput)
		var buf bytes.Buffer
		response.Write(&buf)

		return template.HTML(userInput), false
	case "noop":
		return template.HTML("NOOP"), false
	default:
		log.Fatal("Error running responseHandler. No option passed")
	}

	return template.HTML(""), false
}

func xssTemplate(w http.ResponseWriter, r *http.Request, pd utils.Parameters) (template.HTML, bool) {
	var buf bytes.Buffer

	err := templates.ExecuteTemplate(&buf, "xss", pd.Rulebar[pd.Name])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return template.HTML(buf.String()), true
}

// Handler is the API handler for XSS
func Handler(w http.ResponseWriter, r *http.Request, pd utils.Parameters) (template.HTML, bool) {
	splitURL := strings.Split(r.URL.Path, "/")
	switch splitURL[2] {
	case "query":
		return queryHandler(w, r, splitURL[4])
	case "buffered-query":
		return bufferedQueryHandler(w, r, splitURL[4])
	case "params":
		return paramsHandler(w, r, splitURL[6])
	case "body":
		return bodyHandler(w, r, splitURL[4])
	case "buffered-body":
		return bufferedBodyHandler(w, r, splitURL[4])
	case "response":
		return responseHandler(w, r, splitURL[4])
	default:
		return xssTemplate(w, r, pd)
	}
}
