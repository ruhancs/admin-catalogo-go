package usecase

import (
	"admin-catalogo-go/internal/domain/entity"
	"admin-catalogo-go/internal/domain/gateway"
	"context"
)

type GetCategoryUseCase struct {
	CategoryRepository gateway.CategoryRepositoryInterface
}

func NewGetCategoryUseCase(repository gateway.CategoryRepositoryInterface) *GetCategoryUseCase {
	return &GetCategoryUseCase{
		CategoryRepository: repository,
	}
}

func(usecase *GetCategoryUseCase) Execute(ctx context.Context, id string) (entity.Category,error) {
	category,err := usecase.CategoryRepository.FindByID(ctx,id)
	if err != nil {
		return entity.Category{},err
	}

	return category,nil
}