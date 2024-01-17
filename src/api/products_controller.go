package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Api) configureProductRoutes() {
	group := a.echo.Group("/products")

	group.GET("/", a.listProducts)
}

func (a *Api) listProducts(c echo.Context) (err error) {
	listProductsResult := a.productServices.List()

	if !listProductsResult.Success {
		return c.JSON(http.StatusBadRequest, listProductsResult)
	}

	return c.JSON(http.StatusOK, listProductsResult)
}
