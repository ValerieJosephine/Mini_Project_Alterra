package controller

import (
	"net/http"

	"MINI_PROJECT_ALTERRA/models"
	// "MINI_PROJECT_ALTERRA/service"

	"github.com/labstack/echo/v4"
)

// type ProductController struct {
// 	service service.Service
// }

// func NewProductController(service service.Service) *ProductController {
// 	return &ProductController{service: service}
// }

//>ambil get smua data
func (ctrl *AllController) FetchAllproduct(c echo.Context) error {
	data, err := ctrl.service.FetchAllproduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
	// 	result, err := repository.FetchProducts()
	// 	if err != nil {
	// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	// 	}
	// 	return c.JSON(http.StatusOK, result)
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

// func StoreProduct(c echo.Context) error {
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

// 	result, err := models.StoreProduct(conv_Id, description_product, status, conv_stock, conv_price, conv_updated)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, result)
// }

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
