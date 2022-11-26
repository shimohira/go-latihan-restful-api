package main

import (
	"latihan-restful-api/helper"
	"latihan-restful-api/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	//server := InitializedServer()
	server := InitializedServerBinding()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
