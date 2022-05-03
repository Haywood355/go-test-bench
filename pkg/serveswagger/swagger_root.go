package serveswagger

import (
	"net/http"

	"github.com/Contrast-Security-OSS/go-test-bench/cmd/go-swagger/restapi/operations/swagger_server"
	"github.com/Contrast-Security-OSS/go-test-bench/internal/common"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

func SwaggerRootHandler(params swagger_server.RootParams) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, p runtime.Producer) {
		t := common.Templates["index.gohtml"]

		w.Header().Set("Application-Framework", "Stdlib")
		err := t.ExecuteTemplate(w, "layout.gohtml", Pd)
		if err != nil {
			// panic(err)
		}
		// if err := p.Produce(w, )
		// fmt.Println(p)
		// w.Write([]byte("test"))
	})
	// return CustomResponder(func(w http.ResponseWriter, producer runtime.Producer) {
	//
	// 	var t *template.Template
	// 	//t = servestd.Templates["underConstruction.gohtml"]
	// 	t = servestd.Templates["underConstruction.gohtml"]
	//
	// 	//err := t.ExecuteTemplate(w, "layout.gohtml", Pd)
	// 	err := t.ExecuteTemplate(w, "underConstruction.gohtml", Pd)
	//
	// 	if err != nil {
	// 		log.Print(err.Error())
	// 	}
	//
	// 	payload := "Payload for the swagger root"
	// 	if err := producer.Produce(w, payload); err != nil {
	// 		panic(err) // let the recovery middleware deal with this
	// 	}
	// })
}
