package usecase

import (
	"admin-catalogo-go/internal/application/dto"
	"admin-catalogo-go/internal/domain/gateway"
	"context"
)

type ListVideoUseCase struct {
	VideoRepository gateway.VideoRepositoryInterface
}

func NewListVideoUseCase(repository gateway.VideoRepositoryInterface) *ListVideoUseCase {
	return &ListVideoUseCase{
		VideoRepository: repository,
	}
}

func (usecase *ListVideoUseCase) Execute(ctx context.Context, input dto.ListVideoInputDto) (dto.ListVideoOutputDto,error) {
	videos,err := usecase.VideoRepository.ListVideos(ctx,input)
	if err != nil {
		return dto.ListVideoOutputDto{}, err
	}

	output := dto.ListVideoOutputDto{
		Items: videos,
		Total: len(videos),
		CurrentPage: input.Page,
		PerPage: input.PerPage,
	}

	return output,nil
}