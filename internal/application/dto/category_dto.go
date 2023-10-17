package dto

import (
	"admin-catalogo-go/internal/domain/entity"
	"time"
)

type CreateCategoryInputDto struct {
	Name        string
	Description string
}

type CreateCategoryOutputDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
}

type ListCategoryInputDto struct {
	PerPage int
	Page    int
	Sort    string
	Filter  string
}

type ListCategoryOutputDto struct {
	Items []entity.Category
	Total int
	CurrentPage int
	//LastPage int
	PerPage int
}
