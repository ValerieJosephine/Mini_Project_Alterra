package routes

import (
	"net/http"

	"MINI_PROJECT_ALTERRA/controller"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "welcome to warehouse!")
	})

	e.GET("product", controller.FetchAllproduct)
	e.POST("product", controller.StoreProduct)
	e.POST("product", controller.UpdateProduct)
	e.DELETE("product", controller.DeleteProduct)

	return e
}
