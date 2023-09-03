//go:build wireinject
// +build wireinject

package main

import (
	"golang-restfull-api/app"
	"golang-restfull-api/controller"
	"golang-restfull-api/middleware"
	"golang-restfull-api/repository"
	"golang-restfull-api/service"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryServices,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializeServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet, //binding karna returnnya implementasi dari interface buka interfacenya langsung
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
