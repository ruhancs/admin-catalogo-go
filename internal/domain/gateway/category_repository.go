package gateway

import (
	"admin-catalogo-go/internal/application/dto"
	"admin-catalogo-go/internal/domain/entity"
	"context"
)

type CategoryRepositoryInterface interface {
	Insert(ctx context.Context, category *entity.Category) error
	FindByID(ctx context.Context, id string) (entity.Category, error)
	ListCategory(ctx context.Context, params dto.ListCategoryInputDto) ([]entity.Category,error)
	Update(ctx context.Context, category *entity.Category) error
	Delete(ctx context.Context, id string) error
}
