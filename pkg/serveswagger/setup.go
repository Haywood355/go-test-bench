package serveswagger

import (
	"net/http"

	"github.com/Contrast-Security-OSS/go-test-bench/internal/common"
	"github.com/Contrast-Security-OSS/go-test-bench/internal/injection/cmdi"
	"github.com/go-openapi/runtime"
)

// This lets us implement the Responder interface from go-swagger
// so we can have access to the http.ResponseWriter when each route gets exercised.
type CustomResponder func(http.ResponseWriter, runtime.Producer)

func (c CustomResponder) WriteResponse(w http.ResponseWriter, p runtime.Producer) {
	c(w, p)
}

var Pd = common.ConstParams{
	Year:      2022,
	Logo:      "https://raw.githubusercontent.com/swaggo/swag/master/assets/swaggo.png",
	Framework: "Go-Swagger",
}

func Setup() error {
	if err := common.ParseViewTemplates(); err != nil {
		return err
	}

	cmdi.RegisterRoutes("go-swagger")
	Pd.Rulebar = common.PopulateRouteMap(common.AllRoutes)
	return nil
}
