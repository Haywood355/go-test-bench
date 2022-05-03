// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/Contrast-Security-OSS/go-test-bench/cmd/go-swagger/restapi/operations/cmd_injection"
	"github.com/Contrast-Security-OSS/go-test-bench/cmd/go-swagger/restapi/operations/swagger_server"
)

// NewSwaggerBenchAPI creates a new SwaggerBench instance
func NewSwaggerBenchAPI(spec *loads.Document) *SwaggerBenchAPI {
	return &SwaggerBenchAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		HTMLProducer: runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
			return errors.NotImplemented("html producer has not yet been implemented")
		}),
		JSONProducer: runtime.JSONProducer(),
		TxtProducer:  runtime.TextProducer(),

		CmdInjectionCmdInjectionFrontHandler: cmd_injection.CmdInjectionFrontHandlerFunc(func(params cmd_injection.CmdInjectionFrontParams) middleware.Responder {
			return middleware.NotImplemented("operation cmd_injection.CmdInjectionFront has not yet been implemented")
		}),
		CmdInjectionGetQueryExploitHandler: cmd_injection.GetQueryExploitHandlerFunc(func(params cmd_injection.GetQueryExploitParams) middleware.Responder {
			return middleware.NotImplemented("operation cmd_injection.GetQueryExploit has not yet been implemented")
		}),
		CmdInjectionPostCookiesExploitHandler: cmd_injection.PostCookiesExploitHandlerFunc(func(params cmd_injection.PostCookiesExploitParams) middleware.Responder {
			return middleware.NotImplemented("operation cmd_injection.PostCookiesExploit has not yet been implemented")
		}),
		SwaggerServerRootHandler: swagger_server.RootHandlerFunc(func(params swagger_server.RootParams) middleware.Responder {
			return middleware.NotImplemented("operation swagger_server.Root has not yet been implemented")
		}),
	}
}

/*SwaggerBenchAPI An API built with go-swagger to generate intentionally vulnerable endpoints */
type SwaggerBenchAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/io.goswagger.go-test-bench.v1+json
	JSONConsumer runtime.Consumer

	// HTMLProducer registers a producer for the following mime types:
	//   - text/html
	HTMLProducer runtime.Producer
	// JSONProducer registers a producer for the following mime types:
	//   - application/io.goswagger.go-test-bench.v1+json
	JSONProducer runtime.Producer
	// TxtProducer registers a producer for the following mime types:
	//   - text/plain
	TxtProducer runtime.Producer

	// CmdInjectionCmdInjectionFrontHandler sets the operation handler for the cmd injection front operation
	CmdInjectionCmdInjectionFrontHandler cmd_injection.CmdInjectionFrontHandler
	// CmdInjectionGetQueryExploitHandler sets the operation handler for the get query exploit operation
	CmdInjectionGetQueryExploitHandler cmd_injection.GetQueryExploitHandler
	// CmdInjectionPostCookiesExploitHandler sets the operation handler for the post cookies exploit operation
	CmdInjectionPostCookiesExploitHandler cmd_injection.PostCookiesExploitHandler
	// SwaggerServerRootHandler sets the operation handler for the root operation
	SwaggerServerRootHandler swagger_server.RootHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *SwaggerBenchAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *SwaggerBenchAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *SwaggerBenchAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *SwaggerBenchAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *SwaggerBenchAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *SwaggerBenchAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *SwaggerBenchAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *SwaggerBenchAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *SwaggerBenchAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the SwaggerBenchAPI
func (o *SwaggerBenchAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.HTMLProducer == nil {
		unregistered = append(unregistered, "HTMLProducer")
	}
	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}
	if o.TxtProducer == nil {
		unregistered = append(unregistered, "TxtProducer")
	}

	if o.CmdInjectionCmdInjectionFrontHandler == nil {
		unregistered = append(unregistered, "cmd_injection.CmdInjectionFrontHandler")
	}
	if o.CmdInjectionGetQueryExploitHandler == nil {
		unregistered = append(unregistered, "cmd_injection.GetQueryExploitHandler")
	}
	if o.CmdInjectionPostCookiesExploitHandler == nil {
		unregistered = append(unregistered, "cmd_injection.PostCookiesExploitHandler")
	}
	if o.SwaggerServerRootHandler == nil {
		unregistered = append(unregistered, "swagger_server.RootHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *SwaggerBenchAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *SwaggerBenchAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *SwaggerBenchAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *SwaggerBenchAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/io.goswagger.go-test-bench.v1+json":
			result["application/io.goswagger.go-test-bench.v1+json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *SwaggerBenchAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "text/html":
			result["text/html"] = o.HTMLProducer
		case "application/io.goswagger.go-test-bench.v1+json":
			result["application/io.goswagger.go-test-bench.v1+json"] = o.JSONProducer
		case "text/plain":
			result["text/plain"] = o.TxtProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *SwaggerBenchAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the swagger bench API
func (o *SwaggerBenchAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *SwaggerBenchAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/cmdInjection"] = cmd_injection.NewCmdInjectionFront(o.context, o.CmdInjectionCmdInjectionFrontHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/cmdInjection/{command}/query/{safety}"] = cmd_injection.NewGetQueryExploit(o.context, o.CmdInjectionGetQueryExploitHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/cmdInjection/{command}/cookies/{safety}"] = cmd_injection.NewPostCookiesExploit(o.context, o.CmdInjectionPostCookiesExploitHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"][""] = swagger_server.NewRoot(o.context, o.SwaggerServerRootHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *SwaggerBenchAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *SwaggerBenchAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *SwaggerBenchAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *SwaggerBenchAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *SwaggerBenchAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
