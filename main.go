package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"latihan-restful-api/app"
	"latihan-restful-api/controller"
	"latihan-restful-api/exception"
	"latihan-restful-api/helper"
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

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
