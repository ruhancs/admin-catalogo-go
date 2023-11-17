package web

import (
	"admin-catalogo-go/internal/application/usecase"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Application struct {
	ObserverErrors                prometheus.Counter
	CreateCategoryUseCase         usecase.CreateCategoryUseCase
	GetCategoryUseCase            usecase.GetCategoryUseCase
	DeleteCategoryUseCase         usecase.DeleteCategoryUseCase
	ListCategoryUseCase           usecase.ListCategoryUseCase
	RegisterVideoFileUseCase      usecase.RegisterVideoFileUseCase
	RegisterVideoMetatUseCase     usecase.RegisterVideoMetaUseCase
	ListVideosUseCase             usecase.ListVideoUseCase
	GetVideoByIDUseCase           usecase.GetVideoByIDUseCase
	GetVideoByCategoryUseCase     usecase.GetVideoByCategoryUseCase
	UpdateVideoToPublishedUseCase usecase.UpdateVideoToPublishUseCase
}

func NewApplication(
	observerErrors prometheus.Counter,
	createCategoryUseCase usecase.CreateCategoryUseCase,
	getCategoryUseCase usecase.GetCategoryUseCase,
	deleteCategoryUseCase usecase.DeleteCategoryUseCase,
	listCategoryUseCase usecase.ListCategoryUseCase,
	registerVideoUseCase usecase.RegisterVideoFileUseCase,
	registerVideoMetatUseCase usecase.RegisterVideoMetaUseCase,
	listVideosUseCase usecase.ListVideoUseCase,
	getVideoByIDUseCase usecase.GetVideoByIDUseCase,
	getVideoByCategoryUseCase usecase.GetVideoByCategoryUseCase,
	updateVideoToPublishedUseCase usecase.UpdateVideoToPublishUseCase,
) *Application {
	return &Application{
		ObserverErrors:                observerErrors,
		CreateCategoryUseCase:         createCategoryUseCase,
		ListCategoryUseCase:           listCategoryUseCase,
		GetCategoryUseCase:            getCategoryUseCase,
		DeleteCategoryUseCase:         deleteCategoryUseCase,
		RegisterVideoFileUseCase:      registerVideoUseCase,
		RegisterVideoMetatUseCase:     registerVideoMetatUseCase,
		ListVideosUseCase:             listVideosUseCase,
		GetVideoByIDUseCase:           getVideoByIDUseCase,
		GetVideoByCategoryUseCase:     getVideoByCategoryUseCase,
		UpdateVideoToPublishedUseCase: updateVideoToPublishedUseCase,
	}
}

func (app *Application) Server() error {
	srv := &http.Server{
		Addr:              ":8000",
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	log.Println("Runing server on port 8000...")
	return srv.ListenAndServe()
}
