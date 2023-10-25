package factory

import (
	"admin-catalogo-go/internal/application/usecase"
	"admin-catalogo-go/internal/application/validation"
	"admin-catalogo-go/internal/event"
	"admin-catalogo-go/internal/infra/repository"
	"admin-catalogo-go/pkg/events"
	"database/sql"

	"github.com/aws/aws-sdk-go/service/s3"
)

func RegisterVideoFileUseCaseFactory(
	db *sql.DB, 
	eventDispatcher events.EventDispatcherInterface, 
	s3Client *s3.S3,
	categoryIDValidator *validation.CategoryIDsValidator,
) *usecase.RegisterVideoFileUseCase {
	videoRepository := repository.NewVideoRepository(db)
	videoRegisteredEvent := event.NewVideoRegistered()
	videoRegisteredUseCase := usecase.NewRegisterVideoFileUseCase(videoRepository,videoRegisteredEvent,eventDispatcher,s3Client,categoryIDValidator)
	return videoRegisteredUseCase
}

func RegisterVideoMetaUseCaseFactory(
	db *sql.DB, 
	eventDispatcher events.EventDispatcherInterface, 
	s3Client *s3.S3,
	categoryIDValidator *validation.CategoryIDsValidator,
) *usecase.RegisterVideoMetaUseCase {
	videoRepository := repository.NewVideoRepository(db)
	videoRegisteredEvent := event.NewVideoRegistered()
	videoRegisteredUseCase := usecase.NewRegisterVideoMetaUseCase(videoRepository,videoRegisteredEvent,eventDispatcher,s3Client,categoryIDValidator)
	return videoRegisteredUseCase
}

func ListVideosUsecaseFactory(db *sql.DB) *usecase.ListVideoUseCase {
	videoRepository := repository.NewVideoRepository(db)
	listVideoUseCase := usecase.NewListVideoUseCase(videoRepository)
	return listVideoUseCase
}

func GetVideoByIDUsecaseFactory(db *sql.DB) *usecase.GetVideoByIDUseCase {
	videoRepository := repository.NewVideoRepository(db)
	getVideoByIdUseCase := usecase.NewGetVideoByIDUseCase(videoRepository)
	return getVideoByIdUseCase
}

func GetVideoByCategoryUsecaseFactory(db *sql.DB) *usecase.GetVideoByCategoryUseCase {
	videoRepository := repository.NewVideoRepository(db)
	getVideoByCategoryUseCase := usecase.NewGetVideoByCategoryUseCase(videoRepository)
	return getVideoByCategoryUseCase
}

func UpdateVideoPublishedUseCaseFactory(
	db *sql.DB, 
	eventDispatcher events.EventDispatcherInterface, 
) *usecase.UpdateVideoToPublishUseCase {
	videoRepository := repository.NewVideoRepository(db)
	videoPublishedEvent := event.NewVideoPublish()
	updateVideoPublishUseCase := usecase.NewUpdateVideoToPublishUseCase(videoRepository,videoPublishedEvent,eventDispatcher)
	return updateVideoPublishUseCase
}