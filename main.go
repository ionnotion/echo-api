package main

import (
	"net/http"

	"github.com/ionnotion/echo-api/controllers"
	"github.com/ionnotion/echo-api/models"
	"github.com/labstack/echo/v4"
)

func main()  {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	models.DbConnect()
	e.GET("/properties", controllers.GetAll)
	e.GET("/properties/:id", controllers.GetById)
	e.POST("/properties", controllers.PostPromotion)
	e.PUT("/properties/:id", controllers.UpdatePromotion)
	e.DELETE("/properties/:id",controllers.DeletePromotion)

	e.Logger.Fatal(e.Start(":1323"))
}