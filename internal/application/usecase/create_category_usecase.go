package usecase

import (
	"admin-catalogo-go/internal/application/dto"
	"admin-catalogo-go/internal/domain/gateway"
	"admin-catalogo-go/internal/domain/entity"
	"admin-catalogo-go/pkg/events"
	"context"
)

type CreateCategoryUseCase struct {
	CategoryRepository   gateway.CategoryRepositoryInterface
	CategoryCreatedEvent events.EventInterface
	EventDispatcher      events.EventDispatcherInterface
}

func NewCreateCategoryUseCase(
	repository gateway.CategoryRepositoryInterface, 
	createCategoryEvent events.EventInterface, 
	evDispatcher events.EventDispatcherInterface,
) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{
		CategoryRepository: repository,
		CategoryCreatedEvent: createCategoryEvent,
		EventDispatcher: evDispatcher,
	}
}

func (usecase *CreateCategoryUseCase) Execute(ctx context.Context, input dto.CreateCategoryInputDto) (dto.CreateCategoryOutputDto, error) {
	category, err := entity.NewCategory(input.Name, input.Description)
	if err != nil {
		return dto.CreateCategoryOutputDto{}, err
	}

	err = usecase.CategoryRepository.Insert(ctx, category)
	if err != nil {
		return dto.CreateCategoryOutputDto{}, err
	}

	outputDto := dto.CreateCategoryOutputDto{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive,
		CreatedAt:   category.CreatedAt,
	}

	usecase.CategoryCreatedEvent.SetPayload(outputDto)
	usecase.EventDispatcher.Dispatch(usecase.CategoryCreatedEvent)

	return outputDto, nil
}
