package api

import (
	configDB "golangdemo/func/config"
	"golangdemo/func/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() {

	h := configDB.UserHandler{}

	h.Initialize()

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/item", routes.GetName)
	e.GET("/jsonitem", routes.GetjsonName)

	e.Logger.Fatal(e.Start(":1234"))
}
