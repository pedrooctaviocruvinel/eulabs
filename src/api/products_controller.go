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
	group.PUT("/:id/", a.updateProduct)
	group.DELETE("/:id/", a.deleteProduct)
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

func (a *Api) updateProduct(c echo.Context) (err error) {
	id := c.Param("id")

	convertBodyResult := convertBody[product_services.UpdateProductDTORequest](c)
	if !convertBodyResult.Success {
		return c.JSON(http.StatusBadRequest, convertBodyResult)
	}

	updateProductResult := a.productServices.Update(id, &convertBodyResult.Data)

	if !updateProductResult.Success {
		return c.JSON(http.StatusBadRequest, updateProductResult)
	}

	return c.JSON(http.StatusOK, updateProductResult)
}

func (a *Api) deleteProduct(c echo.Context) (err error) {
	id := c.Param("id")

	deleteProductResult := a.productServices.Delete(id)

	if !deleteProductResult.Success {
		return c.JSON(http.StatusBadRequest, deleteProductResult)
	}

	return c.JSON(http.StatusOK, deleteProductResult)
}
