package dto

import (
	"mime/multipart"
)

type RegisterVideoInputDto struct {
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	YearLaunched int            `json:"year_launched"`
	Banner       multipart.File `json:"banner"`
	Video        multipart.File `json:"video"`
	BannerName   string         `json:"-"`
	VideoName    string         `json:"-"`
	Duration     float64        `json:"duration"`
}

type RegisterVideoOutputDto struct {
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	YearLaunched int     `json:"year_launched"`
	Video_Url    string  `json:"video_url"`
	Banner_Url   string  `json:"banner_url"`
	Duration     float64 `json:"duration"`
	IsPublished  bool    `json:"is_published"`
}
