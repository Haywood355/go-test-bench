// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Contrast-Security-OSS/go-test-bench/cmd/go-swagger/restapi/operations"
	"github.com/Contrast-Security-OSS/go-test-bench/cmd/go-swagger/restapi/operations/cmd_injection"
	"github.com/Contrast-Security-OSS/go-test-bench/cmd/go-swagger/restapi/operations/swagger_server"
	"github.com/Contrast-Security-OSS/go-test-bench/internal/common"
	"github.com/Contrast-Security-OSS/go-test-bench/pkg/serveswagger"
)

//go:generate swagger generate server --target ../../go-swagger --name SwaggerBench --spec ../swagger.yml --principal interface{}

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

	api.HTMLProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		var (
			t                  = common.Templates["underConstruction.gohtml"]
			params interface{} = serveswagger.Pd
		)
		if str, ok := data.(string); ok {
			for _, r := range common.AllRoutes {
				fmt.Println(str, r.Base)
				if str != r.Base {
					continue
				}
				tmpl, ok := common.Templates[r.TmplFile]
				if !ok {
					break
				}
				t = tmpl
				params = common.Parameters{
					ConstParams: serveswagger.Pd,
					Name:        r.Base,
				}
			}
		}

		t.ExecuteTemplate(w, "layout.gohtml", params)
		return nil
	})

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
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
