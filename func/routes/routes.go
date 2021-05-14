package routes

import (
	"golangdemo/func/route"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetName(c echo.Context) error {

	return c.String(http.StatusOK, "Hello, World!")

}

func GetjsonName(c echo.Context) error {
	u := route.User{
		Name:  "Sanchai Pengboot",
		Email: "sanchaipengboot@gmail.com",
	}

	return c.JSON(http.StatusOK, u)
}
