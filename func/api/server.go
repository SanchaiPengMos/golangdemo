package api

import (
	"golangdemo/func/routes"

	"github.com/labstack/echo/v4"
)

func New() {
	e := echo.New()
	e.GET("/item", routes.GetName)
	e.GET("/jsonitem", routes.GetjsonName)

	e.Logger.Fatal(e.Start(":1234"))
}
