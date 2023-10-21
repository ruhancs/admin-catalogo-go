package web

import (
	"admin-catalogo-go/internal/application/dto"
	"encoding/json"
	"fmt"
	"net/http"
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