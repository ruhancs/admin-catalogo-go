package web

import (
	"admin-catalogo-go/internal/application/dto"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func(app *Application) RegisterVideoHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto.RegisterVideoInputDto
	//arquivo de upload de banner maximo 10MB
	r.ParseMultipartForm(10 << 20)
	file,handler,err := r.FormFile("banner")
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
	}
	defer file.Close()
	inputDto.BannerName = handler.Filename 
	
	video,videoHandler,err := r.FormFile("video")
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
	}
	defer video.Close()
	inputDto.VideoName = videoHandler.Filename
	
	err = json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		fmt.Println(err)
		app.errorJson(w,err,http.StatusBadRequest)
	}
	
	inputDto.Video = video
	inputDto.Banner = file
	
	outputDto,err := app.RegisterVideoUseCase.Execute(r.Context(),inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
	}

	app.writeJson(w, http.StatusCreated, outputDto)
}

func(app *Application) ListVideosHandler(w http.ResponseWriter, r *http.Request) {
	var limit int
	var page int
	queryLimit := r.URL.Query().Get("limit")
	queryPage := r.URL.Query().Get("page")
	limit,err := strconv.Atoi(queryLimit)
	if err != nil {
		limit = 10
	}
	page,err = strconv.Atoi(queryPage)
	if err != nil {
		page = 1
	}

	inputListVideoDto := dto.ListVideoInputDto{
		PerPage: limit,
		Page: page,
	}

	output,err := app.ListVideosUseCase.Execute(r.Context(),inputListVideoDto)
	if err != nil {
		app.errorJson(w,err,http.StatusInternalServerError)
	}

	app.writeJson(w,http.StatusOK,output)
}

func(app *Application) GetVideoByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r,"id")

	output,err := app.GetVideoByIDUseCase.Execute(r.Context(),id)
	if err != nil {
		app.errorJson(w,err,http.StatusNotFound)
	}

	app.writeJson(w,http.StatusOK,output)
}

func(app *Application) GetVideoByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r,"category_id")

	output,err := app.GetVideoByCategoryUseCase.Execute(r.Context(),id)
	if err != nil {
		app.errorJson(w,err,http.StatusNotFound)
	}

	app.writeJson(w,http.StatusOK,output)
}

func(app *Application) UpadteVideoPublishStateHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r,"id")
	var inputDto dto.UpdateVideoPublishStateInputDto
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
	}

	output,err := app.UpdateVideoToPublishedUseCase.Execute(r.Context(),id,inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusNotFound)
	}

	app.writeJson(w,http.StatusOK,output)
}