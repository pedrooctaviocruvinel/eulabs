package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pedrooctaviocruvinel/eulabs/src/services/product_services"
)

func (a *Api) configureProductRoutes() {
	group := a.echo.Group("/products")

	group.GET("/", a.listProducts)
	group.POST("/", a.createProduct)
}

func (a *Api) listProducts(c echo.Context) (err error) {
	listProductsResult := a.productServices.List()

	if !listProductsResult.Success {
		return c.JSON(http.StatusBadRequest, listProductsResult)
	}

	return c.JSON(http.StatusOK, listProductsResult)
}

func (a *Api) createProduct(c echo.Context) (err error) {
	convertBodyResult := convertBody[product_services.CreateProductDTORequest](c)
	if convertBodyResult.Success {
		return c.JSON(http.StatusBadRequest, convertBodyResult)
	}

	createProductResult := a.productServices.Create(&convertBodyResult.Data)

	if !createProductResult.Success {
		return c.JSON(http.StatusBadRequest, createProductResult)
	}

	return c.JSON(http.StatusOK, createProductResult)
}
