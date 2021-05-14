package api

import (
	"database/sql"
	"golangdemo/func/route"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

// DB
type UserHandler struct {
	DB *sql.DB
}

func (h *UserHandler) Initialize() {
	// connStr := "postgres://postgres:sanchai@34.87.1.245/demolab?sslmode=disable"
	conStr2 := "postgres://sanchaipengboot:sanchai@localhost/sanchaipengboot?sslmode=disable"
	db, err := sql.Open("postgres", conStr2)
	if err != nil {
		log.Println(err)
	}
	h.DB = db

}

func New() {

	h := UserHandler{}

	h.Initialize()

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	e.GET("/user", h.GetName)

	e.Logger.Fatal(e.Start(":1234"))
}

func (h *UserHandler) GetName(c echo.Context) (err error) {
	rows, err := h.DB.Query("SELECT id , email,firstname,lastname,tel,birthday FROM userhotel order by id asc")
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	var CheckUse []route.UserCheck

	for rows.Next() {
		var data route.UserCheck
		err := rows.Scan(&data.Id, &data.Email, &data.Firstname, &data.Lastname, &data.Tel, &data.Birthday)
		if err != nil {
			log.Println("Scan failed:", err.Error())
		}
		CheckUse = append(CheckUse, data)
	}

	return c.JSON(http.StatusOK, CheckUse)
}
