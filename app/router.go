package app

import (
	"go-clean/controllers"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controllers.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.GetAll)
	router.GET("/api/categories/:categoryId", categoryController.GetByID)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories", categoryController.Delete)

	return router

}
