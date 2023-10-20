package factory

import (
	"admin-catalogo-go/internal/application/usecase"
	"admin-catalogo-go/internal/event"
	"admin-catalogo-go/internal/infra/repository"
	"admin-catalogo-go/pkg/events"
	"database/sql"

	"github.com/aws/aws-sdk-go/service/s3"
)

func RegisterVideoUseCaseFactory(
	db *sql.DB, 
	eventDispatcher events.EventDispatcherInterface, 
	s3Client *s3.S3,
) *usecase.RegisterVideoUseCase {
	videoRepository := repository.NewVideoRepository(db)
	videoRegisteredEvent := event.NewVideoRegistered()
	videoRegisteredUseCase := usecase.NewRegisterVideoUseCase(videoRepository,videoRegisteredEvent,eventDispatcher,s3Client)
	return videoRegisteredUseCase
}