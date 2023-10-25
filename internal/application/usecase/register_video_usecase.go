package usecase

import (
	"admin-catalogo-go/internal/application/dto"
	"admin-catalogo-go/internal/application/validation"
	"admin-catalogo-go/internal/domain/gateway"
	"admin-catalogo-go/internal/infra/cloud"
	"admin-catalogo-go/pkg/events"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/s3"
)

type RegisterVideoFileUseCase struct {
	VideoRepository      gateway.VideoRepositoryInterface
	VideoRegisteredEvent events.EventInterface
	EventDispatcher      events.EventDispatcherInterface
	S3Client             *s3.S3
	CategoryIDValidator  *validation.CategoryIDsValidator
}

func NewRegisterVideoFileUseCase(
	repository gateway.VideoRepositoryInterface,
	event events.EventInterface,
	evDispatcher events.EventDispatcherInterface,
	s3Client *s3.S3,
	categoryIDValidator *validation.CategoryIDsValidator,
) *RegisterVideoFileUseCase {
	return &RegisterVideoFileUseCase{
		VideoRepository:      repository,
		VideoRegisteredEvent: event,
		EventDispatcher:      evDispatcher,
		S3Client:             s3Client,
		CategoryIDValidator:  categoryIDValidator,
	}
}

func (usecase *RegisterVideoFileUseCase) Execute(ctx context.Context, video_id string, input dto.RegisterVideoFilesInputDto) (dto.RegisterVideoFilesOutputDto, error) {
	var outputDto dto.RegisterVideoFilesOutputDto
	video, err := usecase.VideoRepository.GetVideoByID(ctx, video_id)
	if err != nil {
		return dto.RegisterVideoFilesOutputDto{}, err
	}

	//channel para controle de criacaco de threads, evitar lentidao no sistema
	controlChannel := make(chan struct{}, 100)
	//defer close(controlChannel)
	errorUploadChanBanner := make(chan string)
	//defer close(errorUploadChanBanner)
	errorUploadChanVideo := make(chan string)
	//defer close(errorUploadChanVideo)
	errCountChan := make(chan error, 5)
	//defer close(errCountChan)
	//var wg *sync.WaitGroup
	controlChannel <- struct{}{}
	//wg.Add(1)
	go cloud.UploadFileToS3(
		video.Title, 
		input.Video, 
		usecase.S3Client, 
		os.Getenv("VIDEO_BUCKET_NAME"), 
		errorUploadChanVideo, 
		controlChannel, 
		errCountChan)
	controlChannel <- struct{}{}
	//wg.Add(1)
	go cloud.UploadFileToS3(
		video.ID, 
		input.Banner, 
		usecase.S3Client, 
		os.Getenv("VIDEO_BUCKET_NAME"), 
		errorUploadChanBanner, 
		controlChannel, 
		errCountChan)
	go func() {
		for {
			select {
			case filename := <-errorUploadChanBanner:
				controlChannel <- struct{}{}
				//wg.Add(1)
				go cloud.UploadFileToS3(
					filename, 
					input.Banner, 
					usecase.S3Client, 
					os.Getenv("VIDEO_BUCKET_NAME"), 
					errorUploadChanBanner, 
					controlChannel, 
					errCountChan)
			case filename := <-errorUploadChanVideo:
				controlChannel <- struct{}{}
				//wg.Add(1)
				go cloud.UploadFileToS3(
					filename, 
					input.Video, 
					usecase.S3Client, 
					os.Getenv("VIDEO_BUCKET_NAME"), 
					errorUploadChanVideo, 
					controlChannel, 
					errCountChan)
			}
		}
	}()
	//wg.Wait()
	if len(errCountChan) == 5 {
		return dto.RegisterVideoFilesOutputDto{}, errors.New("error to upload file to bucket")
	}
	bannerUrl := fmt.Sprintf("https://%s.s3.us-east-1.amazonaws.com/%s", os.Getenv("VIDEO_BUCKET_NAME"), video.ID)
	videoUrl := fmt.Sprintf("https://%s.s3.us-east-1.amazonaws.com/%s", os.Getenv("VIDEO_BUCKET_NAME"), video.Title)

	_, err = usecase.VideoRepository.UpdateFiles(ctx, video.ID, videoUrl, bannerUrl)
	if err != nil {
		return dto.RegisterVideoFilesOutputDto{}, err
	}

	outputDto.Title = video.Title
	outputDto.Description = video.Description
	outputDto.YearLaunched = video.YearLaunched
	outputDto.Duration = video.Duration
	outputDto.IsPublished = video.IsPublished
	outputDto.Banner_Url = bannerUrl
	outputDto.Video_Url = videoUrl
	outputDto.CategoriesIDs = video.CategoriesID

	usecase.VideoRegisteredEvent.SetPayload(outputDto)
	usecase.EventDispatcher.Dispatch(usecase.VideoRegisteredEvent)

	return outputDto, nil
}
