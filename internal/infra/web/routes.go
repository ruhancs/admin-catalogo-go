package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Heartbeat("/health"))

	mux.Post("/category",app.CreateCategoryHandler)
	mux.Get("/category",app.ListcategoryHandler)
	mux.Get("/category/{id}",app.GetCategoryByID)
	mux.Delete("/category/{id}",app.DeleteCategory)

	mux.Post("/register-video-info",app.RegisterVideoMetaHandler)
	mux.Post("/register-video-files/{id}",app.RegisterVideoFilesHandler)
	mux.Get("/videos",app.ListVideosHandler)
	mux.Get("/video/{id}",app.GetVideoByIDHandler)
	mux.Get("/video/{category_id}",app.GetVideoByCategoryHandler)
	mux.Patch("/video/{id}",app.UpadteVideoPublishStateHandler)

	return mux
}