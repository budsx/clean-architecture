package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-clean/helper"
	"go-clean/models/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
	Find(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
}

type categoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepositoryImpl{}
}

func (repository *categoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO category (name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.NewPanicError(err)

	id, err := result.LastInsertId()
	helper.NewPanicError(err)
	category.Id = int(id)

	return category

}

func (repository *categoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE category SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.NewPanicError(err)

	return category
}

func (repository *categoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "DELETE FROM category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.NewPanicError(err)
}

func (repository *categoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.NewPanicError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.NewPanicError(err)

		categories = append(categories, category)
	}
	return categories
}

func (repository *categoryRepositoryImpl) Find(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "SELECT id, name FROM category where id = ?"
	row, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.NewPanicError(err)

	category := domain.Category{}

	if row.Next() {
		err := row.Scan(&category.Id, &category.Name)
		helper.NewPanicError(err)
		return category, nil
	} else {
		return category, errors.New("Category Not Found")
	}

}
