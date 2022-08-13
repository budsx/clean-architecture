package services

import (
	"context"
	"database/sql"
	"go-clean/helper"
	"go-clean/models/domain"
	"go-clean/models/web"
	"go-clean/repository"
	"log"

	"github.com/go-playground/validator"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	UpdateCategory(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	DeleteCategory(ctx context.Context, categoryId int)
	FindAllCategory(ctx context.Context) []web.CategoryResponse
	FindCategory(ctx context.Context, categoryId int) web.CategoryResponse
}

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) CreateCategory(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.NewPanicError(err)

	tx, err := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}

}

func (service *CategoryServiceImpl) UpdateCategory(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.NewPanicError(err)

	tx, err := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.Find(ctx, tx, request.Id)
	if err != nil {
		log.Println("error: ", err)
	}

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) DeleteCategory(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.NewPanicError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.Find(ctx, tx, categoryId)
	if err != nil {
		log.Println("error: ", err)
	}

	service.CategoryRepository.Delete(ctx, tx, category)

}

func (service *CategoryServiceImpl) FindAllCategory(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.NewPanicError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)
	return helper.ToCategoryResponses(categories)
}

func (service *CategoryServiceImpl) FindCategory(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.NewPanicError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.Find(ctx, tx, categoryId)
	if err != err {
		log.Println(err)
	}
	return helper.ToCategoryResponse(category)
}
