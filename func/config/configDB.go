package configDB

import (
	"database/sql"
	"log"
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
