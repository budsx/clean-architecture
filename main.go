package main

import (
	"go-clean/app"
	"go-clean/controllers"
	"go-clean/repository"
	"go-clean/services"
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/lib/pq"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controllers.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	server.ListenAndServe()

}
