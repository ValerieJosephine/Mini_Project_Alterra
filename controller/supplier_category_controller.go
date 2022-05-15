package controller

import (
	// "MINI_PROJECT_ALTERRA/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ctrl *AllController) GetAllTypeSupplier(c echo.Context) error {
	data, err := ctrl.service.GetAllTypeSupplier()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}
