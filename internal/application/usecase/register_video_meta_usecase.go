package usecase

import (
	"admin-catalogo-go/internal/application/dto"
	"admin-catalogo-go/internal/application/validation"
	"admin-catalogo-go/internal/domain/entity"
	"admin-catalogo-go/internal/domain/gateway"
	"admin-catalogo-go/pkg/events"
	"context"

	"github.com/aws/aws-sdk-go/service/s3"
)

type RegisterVideoMetaUseCase struct {
	VideoRepository      gateway.VideoRepositoryInterface
	VideoRegisteredEvent events.EventInterface
	EventDispatcher      events.EventDispatcherInterface
	CategoryIDValidator  *validation.CategoryIDsValidator
}

func NewRegisterVideoMetaUseCase(
	repository gateway.VideoRepositoryInterface,
	event events.EventInterface,
	evDispatcher events.EventDispatcherInterface,
	s3Client *s3.S3,
	categoryIDValidator  *validation.CategoryIDsValidator,
) *RegisterVideoMetaUseCase {
	return &RegisterVideoMetaUseCase{
		VideoRepository:      repository,
		VideoRegisteredEvent: event,
		EventDispatcher:      evDispatcher,
		CategoryIDValidator: categoryIDValidator,
	}
}

func (usecase *RegisterVideoMetaUseCase) Execute(ctx context.Context, input dto.RegisterVideoMetaInputDto) (dto.RegisterVideoMetaOutputDto, error) {
	var outputDto dto.RegisterVideoMetaOutputDto
	
	//check categories ids
	validCategoriesIs, _ := usecase.CategoryIDValidator.ValidateCategoriesIDs(ctx, input.CategoriesIDs)
	
	video, err := entity.NewVideo(input.Title, input.Description, input.YearLaunched, input.Duration, validCategoriesIs)
	if err != nil {
		return dto.RegisterVideoMetaOutputDto{}, err
	}

	outputDto.Title = video.Title
	outputDto.Description = video.Description
	outputDto.YearLaunched = video.YearLaunched
	outputDto.Duration = video.Duration
	outputDto.IsPublished = video.IsPublished
	outputDto.CategoriesIDs = validCategoriesIs

	err = usecase.VideoRepository.Insert(ctx, video)
	if err != nil {
		return dto.RegisterVideoMetaOutputDto{}, err
	}

	usecase.VideoRegisteredEvent.SetPayload(outputDto)
	usecase.EventDispatcher.Dispatch(usecase.VideoRegisteredEvent)

	return outputDto, nil
}
