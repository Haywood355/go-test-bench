// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"io"
	"os"
	"log"
	"html/template"
	"github.com/Contrast-Security-OSS/go-test-bench/pkg/servestd"
	//"github.com/Contrast-Security-OSS/go-test-bench/pkg/serveswagger"

	"github.com/Contrast-Security-OSS/go-test-bench/internal/injection/cmdi"
	"github.com/Contrast-Security-OSS/go-test-bench/internal/common"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Contrast-Security-OSS/go-test-bench/cmd/go-swagger/restapi/operations"
	"github.com/Contrast-Security-OSS/go-test-bench/cmd/go-swagger/restapi/operations/cmd_injection"
	"github.com/Contrast-Security-OSS/go-test-bench/cmd/go-swagger/restapi/operations/swagger_server"
)

//go:generate swagger generate server --target ../../go-swagger --name SwaggerBench --spec ../swagger.yml --principal interface{} --exclude-main

func configureFlags(api *operations.SwaggerBenchAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SwaggerBenchAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()
	api.TxtProducer = runtime.TextProducer()

	if api.CmdInjectionCmdInjectionFrontHandler == nil {
		api.CmdInjectionCmdInjectionFrontHandler = cmd_injection.CmdInjectionFrontHandlerFunc(func(params cmd_injection.CmdInjectionFrontParams) middleware.Responder {
			return middleware.NotImplemented("operation cmd_injection.CmdInjectionFront has not yet been implemented")
		})
	}
	if api.CmdInjectionGetQueryExploitHandler == nil {
		api.CmdInjectionGetQueryExploitHandler = cmd_injection.GetQueryExploitHandlerFunc(func(params cmd_injection.GetQueryExploitParams) middleware.Responder {
			return middleware.NotImplemented("operation cmd_injection.GetQueryExploit has not yet been implemented")
		})
	}
	if api.CmdInjectionPostCookiesExploitHandler == nil {
		api.CmdInjectionPostCookiesExploitHandler = cmd_injection.PostCookiesExploitHandlerFunc(func(params cmd_injection.PostCookiesExploitParams) middleware.Responder {
			return middleware.NotImplemented("operation cmd_injection.PostCookiesExploit has not yet been implemented")
		})
	}
	if api.SwaggerServerRootHandler == nil {
		api.SwaggerServerRootHandler = swagger_server.RootHandlerFunc(func(params swagger_server.RootParams) middleware.Responder {
			return middleware.NotImplemented("operation swagger_server.Root has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.


func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.

var swaggerPD = common.ConstParams{
	Year:      2022,
	Logo:      "https://raw.githubusercontent.com/swaggo/swag/master/assets/swaggo.png",
	Framework: "Go-Swagger",
}

//TODO: Well if this one turns out to be able to build the UI,
func addUI(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var t *template.Template
		// when should we move away from the std??
		err := servestd.ParseTemplates()
		if err != nil {
			log.Fatalln("Cannot parse Templates:" , err)
		}

		// =========== cmd injection construction =========
		// this has to be first populated so that the Pd
		// object contains a valid RuleBar
		cmdi.RegisterRoutes("go-swagger")

		swaggerPD.Rulebar = common.PopulateRouteMap(common.AllRoutes)

		var parms = common.Parameters{
			ConstParams: swaggerPD,
			Name:        "Command Injection Placeholder",
		}

		// GetUserInput might be doing something weird that
		// plays a more important role than what I was ascribing to it
		var cmdInjFileName = "cmdInjection.gohtml"
		t = servestd.Templates[cmdInjFileName]

		//w.Write([]byte("OK"))
		log.Println("t variable log: ", t)
		err = t.ExecuteTemplate(io.MultiWriter(w, os.Stdout),"layout.gohtml", &parms)
		//err T= t.ExecuteTemplate(w, "layout.gohtml", &parms)
		if err != nil {
			log.Print(err.Error())
		}

		// =========== cmd injection construction =========

		//err := servestd.ParseTemplates()
		//if err != nil {
		//	log.Fatalln("Cannot parse Templates:" , err)
		//}
		//var t *template.Template
		// cmdInjection.gohtml
		//t = servestd.Templates["index.gohtml"]

		//w.Header().Set("Application-Framework", "Go-Swagger")
		//err = t.ExecuteTemplate(w, "layout.gohtml", Pd)
		//if err != nil {
		//	log.Print(err.Error())
		//}

		log.Println("input request:", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	//return handler
	return addUI(handler)
}
