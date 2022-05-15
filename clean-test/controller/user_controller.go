package controller

import (
	"MINI_PROJECT_ALTERRA/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AllController struct {
	service service.Service
}

func NewAllController(service service.Service) *AllController {
	return &AllController{service: service}
}

func (ctrl *AllController) GetAllUser(c echo.Context) error {
	data, err := ctrl.service.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}
