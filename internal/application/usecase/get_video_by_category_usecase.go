package usecase

import (
	"admin-catalogo-go/internal/application/dto"
	"admin-catalogo-go/internal/domain/gateway"
	"context"
)

type GetVideoByCategoryUseCase struct {
	VideoRepository gateway.VideoRepositoryInterface
}

func NewGetVideoByCategoryUseCase(repository gateway.VideoRepositoryInterface) *GetVideoByCategoryUseCase {
	return &GetVideoByCategoryUseCase{
		VideoRepository: repository,
	}
}

func (usecase *GetVideoByCategoryUseCase) Execute(ctx context.Context, categoryID string) ([]dto.GetVideoByCategoryOutputDto, error) {
	videosEntity, err := usecase.VideoRepository.GetVideoByCategoryID(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	var output []dto.GetVideoByCategoryOutputDto
	for _, video := range videosEntity {
		videoDto := dto.GetVideoByCategoryOutputDto{
			Title:         video.Title,
			Description:   video.Description,
			YearLaunched:  video.YearLaunched,
			Duration:      video.Duration,
			BannerUrl:     video.BannerUrl,
			VideoUrl:      video.VideoUrl,
			CategoriesIDs: video.CategoriesID,
			CreatedAt:     video.CreatedAt,
		}
		output = append(output, videoDto)
	}

	return output,nil
}
