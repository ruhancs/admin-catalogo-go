package gateway

import (
	"admin-catalogo-go/internal/application/dto"
	"admin-catalogo-go/internal/domain/entity"
	"context"
)

type VideoRepositoryInterface interface {
	Insert(ctx context.Context,video *entity.Video) error
	ListVideos(ctx context.Context, input dto.ListVideoInputDto) ([]entity.Video,error)
	GetVideoByID(ctx context.Context, id string) (entity.Video,error)
	GetVideoByCategoryID(ctx context.Context, categoryID string) ([]entity.Video,error)
	UpdatePublishState(ctx context.Context,id string, input dto.UpdateVideoPublishStateInputDto) error
}