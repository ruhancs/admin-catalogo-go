package usecase

import (
	"admin-catalogo-go/internal/application/dto"
	"admin-catalogo-go/internal/domain/gateway"
	"context"
)

type ListCategoryUseCase struct {
	CategoryRepository gateway.CategoryRepositoryInterface
}

func NewListCategoryUseCase(repository gateway.CategoryRepositoryInterface) *ListCategoryUseCase{
	return &ListCategoryUseCase{
		CategoryRepository: repository,
	}
}

func(usecase *ListCategoryUseCase) Execute(ctx context.Context, input dto.ListCategoryInputDto) (dto.ListCategoryOutputDto,error) {
	categories,err := usecase.CategoryRepository.ListCategory(ctx,input)
	if err != nil {
		return dto.ListCategoryOutputDto{},err
	}

	outputDto := dto.ListCategoryOutputDto{
		Items: categories,
		Total: len(categories),
		CurrentPage: input.Page,
		PerPage: input.PerPage,
	}

	return outputDto,nil
}