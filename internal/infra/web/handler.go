package web

import (
	"admin-catalogo-go/internal/application/dto"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func(app *Application) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var inputDto dto.CreateCategoryInputDto
	err := json.NewDecoder(r.Body).Decode(&inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
	}
	
	output,err := app.CreateCategoryUseCase.Execute(r.Context(),inputDto)
	if err != nil {
		app.errorJson(w,err,http.StatusBadRequest)
	}

	app.writeJson(w,http.StatusCreated,output)
}

func (app *Application) ListcategoryHandler(w http.ResponseWriter, r *http.Request) {
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

	inputListCategoryDto := dto.ListCategoryInputDto{
		PerPage: limit,
		Page: page,
	}
	output,err := app.ListCategoryUseCase.Execute(r.Context(),inputListCategoryDto)
	if err != nil {
		app.errorJson(w,err, http.StatusInternalServerError)
	}

	app.writeJson(w,http.StatusOK,output)
}

func (app *Application) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	output,err := app.GetCategoryUseCase.Execute(r.Context(), id)
	if err != nil {
		app.errorJson(w,err, http.StatusNotFound)
	}

	app.writeJson(w,http.StatusOK,output)
}

func (app *Application) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := app.DeleteCategoryUseCase.Execute(r.Context(),id)
	if err != nil {
		app.errorJson(w,err, http.StatusNotFound)
	}



	app.writeJson(w,http.StatusNoContent,"{status:deleted}")
}