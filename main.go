package main

import (
	"github.com/go-playground/validator/v10"
	"latihan-restful-api/app"
	"latihan-restful-api/controller"
	"latihan-restful-api/helper"
	"latihan-restful-api/middleware"
	"latihan-restful-api/repository"
	"latihan-restful-api/service"
	"net/http"
)

func main() {
	validate := validator.New()
	db := app.GetConnection()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)
	authMiddleware := middleware.NewAuthMiddleware(router)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
