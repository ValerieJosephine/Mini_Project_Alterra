package routes

import (
	"net/http"

	"MINI_PROJECT_ALTERRA/controller"
	db "MINI_PROJECT_ALTERRA/database"
	"MINI_PROJECT_ALTERRA/repository"
	"MINI_PROJECT_ALTERRA/service"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	db := db.GetConnection()
	AllRepo := repository.NewAllRepository(db)
	AllService := service.NewAllService(AllRepo)
	AllController := controller.NewAllController(*AllService)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "welcome to warehouse!")
	})

	e.GET("/product", AllController.FetchAllproduct)
	e.POST("/product", AllController.StoreAllProduct)
	// e.POST("/product", controller.StoreProduct)
	// e.PUT("/product", controller.UpdateProduct)
	// e.DELETE("/product", controller.DeleteProduct)
	e.GET("/user", AllController.GetAllUser)

	e.GET("/generate-hash/:password", controller.GenerateHashPassword)

	return e
}
