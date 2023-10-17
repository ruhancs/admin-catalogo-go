package usecase

import (
	"admin-catalogo-go/internal/domain/gateway"
	"admin-catalogo-go/pkg/events"
	"context"
)

type DeleteCategoryUseCase struct {
	CategoryRepository gateway.CategoryRepositoryInterface
	CategoryDeletedEvent events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewDeleteCategoryUseCase(
	repository gateway.CategoryRepositoryInterface, 
	event events.EventInterface, 
	evDispatcher events.EventDispatcherInterface) *DeleteCategoryUseCase {
		return &DeleteCategoryUseCase{
			CategoryRepository: repository,
			CategoryDeletedEvent: event,
			EventDispatcher: evDispatcher,
		}
}

func (usecase *DeleteCategoryUseCase) Execute(ctx context.Context, id string) error {
	category,err := usecase.CategoryRepository.FindByID(ctx,id)
	if err != nil {
		return err
	}
	
	err = usecase.CategoryRepository.Delete(ctx,id)
	if err != nil {
		return err
	}

	usecase.CategoryDeletedEvent.SetPayload(category)
	usecase.EventDispatcher.Dispatch(usecase.CategoryDeletedEvent)

	return nil
}