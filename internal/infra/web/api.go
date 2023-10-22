package web

import (
	"admin-catalogo-go/internal/application/usecase"
	"log"
	"net/http"
	"time"
)

type Application struct {
	CreateCategoryUseCase         usecase.CreateCategoryUseCase
	GetCategoryUseCase            usecase.GetCategoryUseCase
	DeleteCategoryUseCase         usecase.DeleteCategoryUseCase
	ListCategoryUseCase           usecase.ListCategoryUseCase
	RegisterVideoUseCase          usecase.RegisterVideoUseCase
	ListVideosUseCase             usecase.ListVideoUseCase
	GetVideoByIDUseCase           usecase.GetVideoByIDUseCase
	GetVideoByCategoryUseCase     usecase.GetVideoByCategoryUseCase
	UpdateVideoToPublishedUseCase usecase.UpdateVideoToPublishUseCase
}

func NewApplication(
	createCategoryUseCase usecase.CreateCategoryUseCase,
	getCategoryUseCase usecase.GetCategoryUseCase,
	deleteCategoryUseCase usecase.DeleteCategoryUseCase,
	listCategoryUseCase usecase.ListCategoryUseCase,
	registerVideoUseCase usecase.RegisterVideoUseCase,
	listVideosUseCase usecase.ListVideoUseCase,
	getVideoByIDUseCase usecase.GetVideoByIDUseCase,
	getVideoByCategoryUseCase usecase.GetVideoByCategoryUseCase,
	updateVideoToPublishedUseCase usecase.UpdateVideoToPublishUseCase,
) *Application {
	return &Application{
		CreateCategoryUseCase:     createCategoryUseCase,
		ListCategoryUseCase:       listCategoryUseCase,
		GetCategoryUseCase:        getCategoryUseCase,
		DeleteCategoryUseCase:     deleteCategoryUseCase,
		RegisterVideoUseCase:      registerVideoUseCase,
		ListVideosUseCase:         listVideosUseCase,
		GetVideoByIDUseCase:       getVideoByIDUseCase,
		GetVideoByCategoryUseCase: getVideoByCategoryUseCase,
		UpdateVideoToPublishedUseCase: updateVideoToPublishedUseCase,
	}
}

func (app *Application) Server() error {
	srv := &http.Server{
		Addr:              ":8000",
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	log.Println("Runing server on port 8000...")
	return srv.ListenAndServe()
}
