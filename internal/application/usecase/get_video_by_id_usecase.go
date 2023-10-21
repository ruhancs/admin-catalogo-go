package usecase

import (
	"admin-catalogo-go/internal/application/dto"
	"admin-catalogo-go/internal/domain/gateway"
	"context"
)

type GetVideoByIDUseCase struct {
	VideoRepository gateway.VideoRepositoryInterface
}

func NewGetVideoByIDUseCase(repository gateway.VideoRepositoryInterface) *GetVideoByIDUseCase {
	return &GetVideoByIDUseCase{
		VideoRepository: repository,
	}
}

func(usecase *GetVideoByIDUseCase) Execute(ctx context.Context, id string) (dto.GetVideoByIDOutputDto, error) {
	videoEntity,err := usecase.VideoRepository.GetVideoByID(ctx,id)
	if err != nil {
		return dto.GetVideoByIDOutputDto{},err
	}

	output := dto.GetVideoByIDOutputDto{
		Title: videoEntity.Title,
		Description: videoEntity.Description,
		YearLaunched: videoEntity.YearLaunched,
		Duration: videoEntity.Duration,
		BannerUrl: videoEntity.BannerUrl,
		VideoUrl: videoEntity.VideoUrl,
		CategoriesIDs: videoEntity.CategoriesID,
		CreatedAt: videoEntity.CreatedAt,
	}

	return output,nil
}