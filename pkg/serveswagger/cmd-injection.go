package serveswagger

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Contrast-Security-OSS/go-test-bench/cmd/go-swagger/restapi/operations/cmd_injection"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

func CommandInjectionHandler(params cmd_injection.CmdInjectionFrontParams) middleware.Responder {

	return CustomResponder(func(w http.ResponseWriter, producer runtime.Producer) {

		var t *template.Template
		// TODO set the htmlproducer to set this by default if missing info
		// t = servestd.Templates["underConstruction.gohtml"]

		// Do we need these headers to be Set this way?
		// TODO: move this to global middleware
		// w.Header().Set("Application-Framework", "Go-Swagger")

		//err := t.ExecuteTemplate(w, "layout.gohtml", Pd)
		err := t.ExecuteTemplate(w, "underConstruction.gohtml", Pd)

		if err != nil {
			log.Print(err.Error())
		}

		w.WriteHeader(200)
		payload := "A payload for the command injection handler"
		if err := producer.Produce(w, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	})
}

