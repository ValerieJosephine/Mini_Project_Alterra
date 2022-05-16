package controller

// import (
// 	"MINI_PROJECT_ALTERRA/service"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// type ControllerUser struct {
// 	service service.Service
// }

// func NewUserController(service service.Service) *ControllerUser {
// 	return &ControllerUser{service: service}
// }

// func (ctrl *ControllerUser) GetAllUser(c echo.Context) error {
// 	data, err := ctrl.service.GetAllUser()
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}
// 	return c.JSON(http.StatusOK, data)
// }
