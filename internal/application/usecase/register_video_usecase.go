package usecase

import (
	"admin-catalogo-go/internal/application/dto"
	"admin-catalogo-go/internal/application/validation"
	"admin-catalogo-go/internal/domain/entity"
	"admin-catalogo-go/internal/domain/gateway"
	"admin-catalogo-go/internal/infra/cloud"
	"admin-catalogo-go/pkg/events"
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/s3"
)

type RegisterVideoUseCase struct {
	VideoRepository      gateway.VideoRepositoryInterface
	VideoRegisteredEvent events.EventInterface
	EventDispatcher      events.EventDispatcherInterface
	S3Client             *s3.S3
	CategoryIDValidator  *validation.CategoryIDsValidator
}

func NewRegisterVideoUseCase(
	repository gateway.VideoRepositoryInterface,
	event events.EventInterface,
	evDispatcher events.EventDispatcherInterface,
	s3Client *s3.S3,
	categoryIDValidator  *validation.CategoryIDsValidator,
) *RegisterVideoUseCase {
	return &RegisterVideoUseCase{
		VideoRepository:      repository,
		VideoRegisteredEvent: event,
		EventDispatcher:      evDispatcher,
		S3Client:             s3Client,
		CategoryIDValidator: categoryIDValidator,
	}
}

func (usecase *RegisterVideoUseCase) Execute(ctx context.Context, input dto.RegisterVideoInputDto) (dto.RegisterVideoOutputDto, error) {
	var outputDto dto.RegisterVideoOutputDto
	outputDto.Title = "input.Title"
	outputDto.Description = "input.Description"
	outputDto.YearLaunched = 2004 //input.YearLaunched

	//check categories ids
	validCategoriesIs, _ := usecase.CategoryIDValidator.ValidateCategoriesIDs(ctx, input.CategoriesIDs)

	//channel para controle de criacaco de threads, evitar lentidao no sistema
	controlChannel := make(chan struct{}, 300)
	errorUploadChan := make(chan string)
	controlChannel <- struct{}{}
	go cloud.UploadFileToS3(input.VideoName, input.Video, usecase.S3Client, os.Getenv("VIDEO_BUCKET_NAME"), errorUploadChan, controlChannel)
	controlChannel <- struct{}{}
	go cloud.UploadFileToS3(input.BannerName, input.Banner, usecase.S3Client, os.Getenv("VIDEO_BUCKET_NAME"), errorUploadChan, controlChannel)
	go func() {
		for {
			select {
			case filename := <-errorUploadChan:
				controlChannel <- struct{}{}
				go cloud.UploadFileToS3(filename, input.Banner, usecase.S3Client, os.Getenv("VIDEO_BUCKET_NAME"), errorUploadChan, controlChannel)
			}
		}
	}()
	bannerUrl := fmt.Sprintf("https://%s.s3.us-east-1.amazonaws.com/%s", os.Getenv("VIDEO_BUCKET_NAME"), input.BannerName)
	videoUrl := fmt.Sprintf("https://%s.s3.us-east-1.amazonaws.com/%s", os.Getenv("VIDEO_BUCKET_NAME"), input.VideoName)
	outputDto.Banner_Url = bannerUrl
	outputDto.Video_Url = videoUrl

	video, err := entity.NewVideo("input.Title", "input.Description", 2004, 2, bannerUrl, videoUrl, validCategoriesIs)
	if err != nil {
		return dto.RegisterVideoOutputDto{}, err
	}
	outputDto.Duration = video.Duration
	outputDto.IsPublished = video.IsPublished

	err = usecase.VideoRepository.Insert(ctx, video)
	if err != nil {
		return dto.RegisterVideoOutputDto{}, err
	}

	usecase.VideoRegisteredEvent.SetPayload(outputDto)
	usecase.EventDispatcher.Dispatch(usecase.VideoRegisteredEvent)

	return outputDto, nil
}
