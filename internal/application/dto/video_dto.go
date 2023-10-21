package dto

import (
	"mime/multipart"
	"time"
)

type RegisterVideoInputDto struct {
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	YearLaunched  int            `json:"year_launched"`
	CategoriesIDs []string       `json:"categories_ids"`
	Banner        multipart.File `json:"banner"`
	Video         multipart.File `json:"video"`
	BannerName    string         `json:"-"`
	VideoName     string         `json:"-"`
	Duration      float64        `json:"duration"`
}

type RegisterVideoOutputDto struct {
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	YearLaunched  int      `json:"year_launched"`
	CategoriesIDs []string `json:"categories_ids"`
	Video_Url     string   `json:"video_url"`
	Banner_Url    string   `json:"banner_url"`
	Duration      float64  `json:"duration"`
	IsPublished   bool     `json:"is_published"`
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
