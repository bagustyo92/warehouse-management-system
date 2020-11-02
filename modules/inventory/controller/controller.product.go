package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bagustyo92/wms/middleware/logger"
	"github.com/bagustyo92/wms/modules/inventory/model"
	"github.com/bagustyo92/wms/utils"
	"github.com/labstack/echo"
)

func (ic inventoryController) createProduct(c echo.Context) error {
	product := new(model.Product)
	if err := c.Bind(product); err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	if err := ic.productServ.InsertProduct(product); err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	return c.JSON(utils.Response(http.StatusOK, "OK", &product))

}

func (ic inventoryController) getProduct(c echo.Context) error {
	productID := c.Param("id")
	id, err := strconv.ParseUint(productID, 10, 16)
	if err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	product, err := ic.productServ.GetProduct(uint(id))
	if err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	return c.JSON(utils.Response(http.StatusOK, "OK", &product))
}

func (ic inventoryController) getProducts(c echo.Context) error {
	query := new(model.Query)
	if err := c.Bind(query); err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	products, err := ic.productServ.GetProducts(query)
	if err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	return c.JSON(utils.Response(http.StatusOK, "OK", &products))
}

func (ic inventoryController) updateProducts(c echo.Context) error {
	product := new(model.Product)
	fmt.Println(product)
	if err := c.Bind(product); err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	if err := ic.productServ.UpdateProduct(product); err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	return c.JSON(utils.Response(http.StatusOK, "OK", "OK"))
}

func (ic inventoryController) deleteProduct(c echo.Context) error {
	productID := c.Param("id")
	id, err := strconv.ParseUint(productID, 10, 16)
	if err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	if err := ic.productServ.DeleteProduct(uint(id)); err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	return c.JSON(utils.Response(http.StatusOK, "OK", "OK"))
}
