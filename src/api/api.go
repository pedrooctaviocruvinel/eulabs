package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pedrooctaviocruvinel/eulabs/src/config"
	"github.com/pedrooctaviocruvinel/eulabs/src/services/product_services"
)

type Api struct {
	echo            *echo.Echo
	productServices *product_services.ProductServices
}

func NewApi(ps *product_services.ProductServices) (api *Api) {
	e := echo.New()

	api = &Api{
		echo:            e,
		productServices: ps,
	}

	api.configureRoutes()
	api.configureMiddlewares()

	return api
}

func (a *Api) configureRoutes() {
	a.configureProductRoutes()
}

func (a *Api) configureMiddlewares() {
	a.echo.Pre(middleware.AddTrailingSlash())
}

func (a *Api) Run() {
	panic(a.echo.Start(fmt.Sprintf(":%s", config.Api.Port)))
}
