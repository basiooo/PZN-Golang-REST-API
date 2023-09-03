package app

import (
	"database/sql"
	"golang-restfull-api/helpers"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/golang_rest_api")
	helpers.PanicIfError(err)

	db.SetConnMaxIdleTime(5)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
