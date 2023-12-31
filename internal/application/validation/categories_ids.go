package validation

import (
	"admin-catalogo-go/internal/domain/gateway"
	"context"
	"fmt"
)

type CategoryIDsValidator struct {
	CategoryRepo gateway.CategoryRepositoryInterface
}

func NewCategoryIDsValidator(categoryRepo gateway.CategoryRepositoryInterface) *CategoryIDsValidator{
	return &CategoryIDsValidator{
		CategoryRepo: categoryRepo,
	}
}

func (vc *CategoryIDsValidator) ValidateCategoriesIDs(ctx context.Context,categoriesID []string) ([]string, error) {
	var validIDs []string
	for _,categoryId := range categoriesID {
		_,err := vc.CategoryRepo.FindByID(ctx,categoryId)
		if err != nil {
			fmt.Println("ERROR TO FIND CATEGORY ID")
			continue
		}
		validIDs = append(validIDs, categoryId)
	}
	return validIDs,nil
}