package controllers

import (
	"go-clean/helper"
	"go-clean/models/web"
	"go-clean/services"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type CategoryControllerImpl struct {
	CategoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	// * Read the body of the request
	helper.ReadFromRequest(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.CreateCategory(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteResponse(writer, webResponse)

}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequest(request, &categoryUpdateRequest)

	// * Get Id from params
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.NewPanicError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.UpdateCategory(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Updated",
		Data:   categoryResponse,
	}

	helper.WriteResponse(writer, webResponse)

}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// * Get Id from params
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.NewPanicError(err)

	controller.CategoryService.DeleteCategory(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Deleted",
	}
	helper.WriteResponse(writer, webResponse)
}

func (controller *CategoryControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAllCategory(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}
	helper.WriteResponse(writer, webResponse)
}

func (controller *CategoryControllerImpl) GetByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.NewPanicError(err)

	categoryResponse := controller.CategoryService.FindCategory(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteResponse(writer, webResponse)
}
