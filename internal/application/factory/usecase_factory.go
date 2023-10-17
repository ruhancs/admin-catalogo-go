package factory

import (
	"admin-catalogo-go/internal/application/usecase"
	"admin-catalogo-go/internal/event"
	"admin-catalogo-go/internal/infra/repository"
	"admin-catalogo-go/pkg/events"
	"database/sql"
)

func CreateCategoryUseCaseFactory(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateCategoryUseCase{
	categoryRepository := repository.NewCategoryRepository(db)
	categoryCreatedEvent := event.NewCategoryCreated()
	createCategoryUseCase := usecase.NewCreateCategoryUseCase(categoryRepository,categoryCreatedEvent,eventDispatcher)
	return createCategoryUseCase
}

func ListCategoryUsecaseFactory(db *sql.DB) *usecase.ListCategoryUseCase{
	categoryRepository := repository.NewCategoryRepository(db)
	listCategoryUseCase := usecase.NewListCategoryUseCase(categoryRepository)
	return listCategoryUseCase
}

func GetCategoryByIDUsecaseFactory(db *sql.DB) *usecase.GetCategoryUseCase{
	categoryRepository := repository.NewCategoryRepository(db)
	getCategoryUseCase := usecase.NewGetCategoryUseCase(categoryRepository)
	return getCategoryUseCase
}

func DeleteCategoryUsecaseFactory(db *sql.DB,eventDispatcher events.EventDispatcherInterface) *usecase.DeleteCategoryUseCase{
	categoryRepository := repository.NewCategoryRepository(db)
	categoryDeletedEvent := event.NewCategoryDeleted()
	deleteCategoryUseCase := usecase.NewDeleteCategoryUseCase(categoryRepository,categoryDeletedEvent,eventDispatcher)
	return deleteCategoryUseCase
}