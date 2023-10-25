package web

import (
	"admin-catalogo-go/internal/application/dto"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func(app *Application) RegisterVideoMetaHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto.RegisterVideoMetaInputDto
	
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}
	
	outputDto,err := app.RegisterVideoMetatUseCase.Execute(r.Context(), inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}

	app.writeJson(w, http.StatusCreated, outputDto)
}

func(app *Application) RegisterVideoFilesHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto.RegisterVideoFilesInputDto
	id := chi.URLParam(r,"id")
	//limitar tamanho da requisicao para 32mb
	r.Body = http.MaxBytesReader(w, r.Body, 32<<20+512)
	//arquivo de upload de banner maximo 10MB
	r.ParseMultipartForm(10 << 20)
	file,_,err := r.FormFile("banner")
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}
	defer file.Close()
	
	video,_,err := r.FormFile("video")
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}
	defer video.Close()
	
	inputDto.Video = video
	inputDto.Banner = file
	
	outputDto,err := app.RegisterVideoFileUseCase.Execute(r.Context(),id, inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
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
		return
	}

	app.writeJson(w,http.StatusOK,output)
}

func(app *Application) GetVideoByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r,"id")

	output,err := app.GetVideoByIDUseCase.Execute(r.Context(),id)
	if err != nil {
		app.errorJson(w,err,http.StatusNotFound)
		return
	}

	app.writeJson(w,http.StatusOK,output)
}

func(app *Application) GetVideoByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r,"category_id")

	output,err := app.GetVideoByCategoryUseCase.Execute(r.Context(),id)
	if err != nil {
		app.errorJson(w,err,http.StatusNotFound)
		return
	}

	app.writeJson(w,http.StatusOK,output)
}

func(app *Application) UpadteVideoPublishStateHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r,"id")
	var inputDto dto.UpdateVideoPublishStateInputDto
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
		return
	}

	output,err := app.UpdateVideoToPublishedUseCase.Execute(r.Context(),id,inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusNotFound)
		return
	}

	app.writeJson(w,http.StatusOK,output)
}