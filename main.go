package main

import (
	"golang-restfull-api/helpers"
	"golang-restfull-api/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8000",
		Handler: authMiddleware,
	}
}
func main() {
	server := InitializeServer()
	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
