package usecase

import (
	"admin-catalogo-go/internal/application/dto"
	"admin-catalogo-go/internal/domain/gateway"
	"admin-catalogo-go/pkg/events"
	"context"
)

type UpdateVideoToPublishUseCase struct {
	VideoRepository     gateway.VideoRepositoryInterface
	VideoPublishedEvent events.EventInterface
	EventDispatcher     events.EventDispatcherInterface
}

func NewUpdateVideoToPublishUseCase(
	repository gateway.VideoRepositoryInterface, 
	videoPublisheEvent events.EventInterface, 
	EvDispatcher events.EventDispatcherInterface,
) *UpdateVideoToPublishUseCase {
	return &UpdateVideoToPublishUseCase{
		VideoRepository: repository,
		VideoPublishedEvent: videoPublisheEvent,
		EventDispatcher: EvDispatcher,
	}
}

func(usecase *UpdateVideoToPublishUseCase) Execute(ctx context.Context,id string, input dto.UpdateVideoPublishStateInputDto) (dto.UpdateVideoPublishStateOutputDto,error) {
	err := usecase.VideoRepository.UpdatePublishState(ctx,id, input)
	if err != nil {
		return dto.UpdateVideoPublishStateOutputDto{},err
	}

	output := dto.UpdateVideoPublishStateOutputDto{
		ID: id,
		IsPublished: input.IsPublished,
	}

	usecase.VideoPublishedEvent.SetPayload(output)
	usecase.EventDispatcher.Dispatch(usecase.VideoPublishedEvent)

	return output,nil
}
