package repository

import (
	"admin-catalogo-go/internal/application/dto"
	"admin-catalogo-go/internal/domain/entity"
	"admin-catalogo-go/internal/infra/db"
	"context"
	"database/sql"
	"errors"
	"log"
)

type CategoryRepository struct {
	DB *sql.DB
	Queries *db.Queries
}

func NewCategoryRepository(database *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		DB: database,
		Queries: db.New(database),
	}
}

func(repo *CategoryRepository) Insert(ctx context.Context, category *entity.Category) error {
	err := repo.Queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID: category.ID,
		Name: category.Name,
		Description: sql.NullString{String:category.Description, Valid: true},
		IsActive: category.IsActive,
		CreatedAt: category.CreatedAt,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (repo *CategoryRepository) FindByID(ctx context.Context, id string) (entity.Category, error) {
	categoryModel,err := repo.Queries.GetCategory(ctx,id)
	if err != nil {
		return entity.Category{},err
	}

	categoryEntity := entity.Category{
		ID: categoryModel.ID,
		Name: categoryModel.Name,
		Description: categoryModel.Description.String,
		IsActive: categoryModel.IsActive,
		CreatedAt: categoryModel.CreatedAt,
	}

	return categoryEntity,nil
}

func (repo *CategoryRepository) ListCategory(ctx context.Context, params dto.ListCategoryInputDto) ([]entity.Category,error) {
	offset := (params.Page - 1) * params.PerPage
	categoriesModel,err := repo.Queries.ListCategories(ctx, db.ListCategoriesParams{
		Limit: int32(params.PerPage),
		Offset: int32(offset),
	})
	if err != nil {
		return nil,err
	}

	var categoriesEntity = []entity.Category{}
	for _,categoryModel := range categoriesModel {
		category := entity.Category{
			ID: categoryModel.ID,
			Name: categoryModel.Name,
			Description: categoryModel.Description.String,
			IsActive: categoryModel.IsActive,
			CreatedAt: categoryModel.CreatedAt,
		}
		categoriesEntity = append(categoriesEntity, category)
	}

	return categoriesEntity,nil
}

func (repo *CategoryRepository) Update(ctx context.Context, category *entity.Category) error {
	return errors.New("not implemented")
}

func (repo *CategoryRepository) Delete(ctx context.Context, id string) error {
	err := repo.Queries.DeleteCategory(ctx,id)
	if err != nil {
		return err
	}
	return nil
}