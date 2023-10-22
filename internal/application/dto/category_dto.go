package dto

import (
	"admin-catalogo-go/internal/domain/entity"
	"time"
)

type CreateCategoryInputDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateCategoryOutputDto struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
}

type ListCategoryInputDto struct {
	PerPage int `json:"per_page"`
	Page    int `json:"page"`
	//Sort    string
	//Filter  string
}

type ListCategoryOutputDto struct {
	Items       []entity.Category `json:"items"`
	Total       int `json:"total"`
	CurrentPage int `json:"current_page"`
	//LastPage int
	PerPage int `json:"per_page"`
}
