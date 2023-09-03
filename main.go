package main

import (
	"golang-restfull-api/app"
	"golang-restfull-api/controller"
	"golang-restfull-api/middleware"
	"golang-restfull-api/repository"
	"golang-restfull-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryServices(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	server.ListenAndServe()
}
