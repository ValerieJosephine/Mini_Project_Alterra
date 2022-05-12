package routes

import (
	"net/http"

	"MINI_PROJECT_ALTERRA/controller"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Selamat datang dalam sistem warehouse!")
	})

	e.GET("product", controller.FetchAllproduct)

	return e
}
