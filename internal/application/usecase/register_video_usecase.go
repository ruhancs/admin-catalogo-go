package usecase

import (
	"admin-catalogo-go/internal/application/dto"
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
}

func NewRegisterVideoUseCase(
	repository gateway.VideoRepositoryInterface,
	event events.EventInterface,
	evDispatcher events.EventDispatcherInterface,
	s3Client *s3.S3,
) *RegisterVideoUseCase {
	return &RegisterVideoUseCase{
		VideoRepository:      repository,
		VideoRegisteredEvent: event,
		EventDispatcher:      evDispatcher,
		S3Client:             s3Client,
	}
}

func (usecase *RegisterVideoUseCase) Execute(ctx context.Context,input dto.RegisterVideoInputDto) (dto.RegisterVideoOutputDto, error) {
	var outputDto dto.RegisterVideoOutputDto
	outputDto.Title = input.Title
	outputDto.Description = input.Description
	outputDto.YearLaunched = input.YearLaunched
	//upload do video, enviar para s3

	//upload do banner, enviar para s3
	errorUploadChan := make(chan string)
	go cloud.UploadFileToS3(input.VideoName, input.Video, usecase.S3Client, os.Getenv("VIDEO_BUCKET_NAME"), errorUploadChan)
	go cloud.UploadFileToS3(input.BannerName, input.Banner, usecase.S3Client, os.Getenv("VIDEO_BUCKET_NAME"), errorUploadChan)
	go func() {
		for {
			select{
			case filename := <- errorUploadChan:
				go cloud.UploadFileToS3(filename, input.Banner, usecase.S3Client, os.Getenv("VIDEO_BUCKET_NAME"), errorUploadChan)
			}
		}
	}()
	
	outputDto.Banner_Url = fmt.Sprintf("https://%s.s3.us-east-1.amazonaws.com/%s",os.Getenv("VIDEO_BUCKET_NAME"),input.BannerName)

	video, err := entity.NewVideo(input.Title, input.Description, input.YearLaunched, 1)
	if err != nil {
		return dto.RegisterVideoOutputDto{}, err
	}
	outputDto.Duration = input.Duration
	outputDto.IsPublished = video.IsPublished

	err = usecase.VideoRepository.Insert(ctx,video)
	if err != nil {
		return dto.RegisterVideoOutputDto{}, err
	}

	usecase.VideoRegisteredEvent.SetPayload(outputDto)
	usecase.EventDispatcher.Dispatch(usecase.VideoRegisteredEvent)

	return outputDto, nil
}

