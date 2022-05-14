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
	productRepo := repository.NewRepository(db)
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(*productService)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "welcome to warehouse!")
	})

	e.GET("/product", productController.FetchAllproduct)
	e.POST("/product", productController.StoreAllProduct)
	// e.POST("/product", controller.StoreProduct)
	// e.PUT("/product", controller.UpdateProduct)
	// e.DELETE("/product", controller.DeleteProduct)

	e.GET("/generate-hash/:password", controller.GenerateHashPassword)

	return e
}
