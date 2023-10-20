package gateway

import (
	"admin-catalogo-go/internal/domain/entity"
	"context"
)

type VideoRepositoryInterface interface {
	Insert(ctx context.Context,video *entity.Video) error
}