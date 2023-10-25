package dto

import (
	"admin-catalogo-go/internal/domain/entity"
	"mime/multipart"
	"time"
)

type RegisterVideoMetaInputDto struct {
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	YearLaunched  int            `json:"year_launched"`
	CategoriesIDs []string       `json:"categories_ids"`
	//BannerName    string         `json:"banner_name"`
	//VideoName     string         `json:"video_name"`
	Duration      float64        `json:"duration"`
}

type RegisterVideoMetaOutputDto struct {
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	YearLaunched  int      `json:"year_launched"`
	CategoriesIDs []string `json:"categories_ids"`
	Duration      float64  `json:"duration"`
	IsPublished   bool     `json:"is_published"`
}

type RegisterVideoFilesInputDto struct {
	Banner        multipart.File `json:"banner"`
	Video         multipart.File `json:"video"`
}

type RegisterVideoFilesOutputDto struct {
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	YearLaunched  int      `json:"year_launched"`
	CategoriesIDs []string `json:"categories_ids"`
	Video_Url     string   `json:"video_url"`
	Banner_Url    string   `json:"banner_url"`
	Duration      float64  `json:"duration"`
	IsPublished   bool     `json:"is_published"`
}

type ListVideoInputDto struct {
	PerPage int `json:"per_page"`
	Page    int `json:"page"`
	//Sort    string
	//Filter  string
}

type ListVideoOutputDto struct {
	Items       []entity.Video `json:"items"`
	Total       int            `json:"total"`
	CurrentPage int            `json:"current_page"`
	//LastPage int
	PerPage int `json:"per_page"`
}

type GetVideoByIDOutputDto struct {
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	YearLaunched  int       `json:"year_launched"`
	CategoriesIDs []string  `json:"categories_ids"`
	BannerUrl     string    `json:"banner_url"`
	VideoUrl      string    `json:"video_url"`
	Duration      float64   `json:"duration"`
	CreatedAt     time.Time `json:"created_at"`
}

type GetVideoByCategoryOutputDto struct {
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	YearLaunched  int       `json:"year_launched"`
	CategoriesIDs []string  `json:"categories_ids"`
	BannerUrl     string    `json:"banner_url"`
	VideoUrl      string    `json:"video_url"`
	Duration      float64   `json:"duration"`
	CreatedAt     time.Time `json:"created_at"`
}

type UpdateVideoPublishStateInputDto struct {
	IsPublished bool   `json:"is_published"`
}

type UpdateVideoPublishStateOutputDto struct {
	ID          string `json:"id"`
	IsPublished bool   `json:"is_published"`
}
