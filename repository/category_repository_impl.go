package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restfull-api/helpers"
	"golang-restfull-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into category(name) value (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)

	helpers.PanicIfError(err)
	id, err := result.LastInsertId()
	helpers.PanicIfError(err)
	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helpers.PanicIfError(err)

	return category

}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helpers.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "select id,name from category where id= ?"

	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helpers.PanicIfError(err)
	category := domain.Category{}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helpers.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category tidak ada")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id,name from category"

	rows, err := tx.QueryContext(ctx, SQL)

	helpers.PanicIfError(err)
	defer rows.Close()
	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helpers.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
