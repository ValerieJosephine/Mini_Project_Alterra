package controller

import (
	"net/http"

	"MINI_PROJECT_ALTERRA/models"

	"github.com/labstack/echo/v4"
)

//>ambil get smua data
func (ctrl *AllController) FetchAllproduct(c echo.Context) error {
	data, err := ctrl.service.FetchAllproduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, data)

}

//>perintah tulis get data by id (only)
func (ctrl *AllController) FetchProductWithId(c echo.Context) error {
	var id int
	data, err := ctrl.service.FetchProductWithId(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, data)
}

//>perintah tulis post data
func (ctrl *AllController) StoreAllProduct(c echo.Context) error {
	var input models.Product
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	data, errData := ctrl.service.StoreAllProduct(input)
	if errData != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

//> perintah delete data
func (ctrl *AllController) DeleteAllProduct(c echo.Context) error {
	var id int
	err := ctrl.service.DeleteAllProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, err)
}

// > perintah update data
// func UpdateProduct(c echo.Context) error {
// 	Id_product := c.FormValue("Id")
// 	description_product := c.FormValue("description_product")
// 	status := c.FormValue("status")
// 	restock_at := c.FormValue("restock_at")
// 	price := c.FormValue("price")
// 	updated_at := c.FormValue("updated_at")

// 	conv_Id, err := strconv.Atoi(Id_product)
// 	conv_stock, err := strconv.Atoi(restock_at)
// 	conv_price, err := strconv.Atoi(price)
// 	conv_updated, err := strconv.Atoi(updated_at)

// 	result, err := models.UpdateProduct(conv_Id, description_product, status, conv_stock, conv_price, conv_updated)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, result)
// }

// > perintah DELETE data
// func DeleteProduct(c echo.Context) error {
// 	Id_product := c.FormValue("Id")

// 	conv_Id, err := strconv.Atoi(Id_product)

// 	result, err := models.DeleteProduct(conv_Id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, result)
// }
