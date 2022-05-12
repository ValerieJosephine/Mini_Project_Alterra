package controller

import (
	"net/http"

	"MINI_PROJECT_ALTERRA/models"

	"github.com/labstack/echo/v4"
)

//ambil get smua data
func FetchAllproduct(c echo.Context) error {
	result, err := models.FetchAllproduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, result)
}

//perintah tulis post data
func StoreProduct(c echo.Context) error {
	Id_product := c.FormValue("Id")
	description_product := c.FormValue("description_product")
	status := c.FormValue("status")
	restock_at := c.FormValue("restock_at")
	price := c.FormValue("price")
	updated_at := c.FormValue("updated_at")

	result, err := models.StoreProduct(Id_product, description_product, status, restock_at, price, updated_at)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
