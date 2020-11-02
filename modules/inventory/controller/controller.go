package controller

import (
	"github.com/bagustyo92/wms/modules/inventory/service"
	"github.com/labstack/echo"
)

type inventoryController struct {
	productServ service.ProductInterface
}

func ApplyController(router *echo.Echo, ps service.ProductInterface) {
	handler := &inventoryController{
		productServ: ps,
	}

	inventory := router.Group("inventory")

	// product
	inventory.POST("/product", handler.createProduct)
	inventory.GET("/product/:id", handler.getProduct)
	inventory.GET("/products", handler.getProducts)
	// inventory.PATCH("/product/:id", handler.updateProduct)
	// inventory.DELETE("/product/:id", handler.deleteProduct)

	// stock

	// inbound

	// outbound
}
