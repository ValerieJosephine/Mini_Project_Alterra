package controller

import (
	"MINI_PROJECT_ALTERRA/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := utils.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}
