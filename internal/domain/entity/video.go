package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Video struct {
	ID           string    `json:"id" valid:"required"`
	Title        string    `json:"title" valid:"required"`
	Description  string    `json:"description"`
	YearLaunched int       `json:"year_launched" valid:"required"`
	Duration     float64   `json:"duration"`
	IsPublished  bool      `json:"is_published"`
	BannerUrl    string    `json:"banner_url"`
	VideoUrl     string    `json:"video_url"`
	CategoriesID []string  `json:"categories_id"`
	CreatedAt    time.Time `json:"created_at"`
}

func NewVideo(
	title,
	description string,
	yearLaunched int,
	duration float64,
	bannerUrl string,
	videorUrl string,
	categoriesID []string,
) (*Video, error) {
	video := &Video{
		ID:           uuid.NewV4().String(),
		Title:        title,
		Description:  description,
		YearLaunched: yearLaunched,
		Duration:     duration,
		BannerUrl:    bannerUrl,
		VideoUrl:     videorUrl,
		IsPublished:  false,
		CategoriesID: categoriesID,
		CreatedAt:    time.Now(),
	}
	err := video.Validate()
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (v *Video) ChangeTitle(newTitle string) error {
	title := v.Title
	v.Title = newTitle
	err := v.Validate()
	if err != nil {
		v.Title = title
		return err
	}
	return nil
}

func (v *Video) ChangeDescription(description string) {
	v.Description = description
}

func (v *Video) MarkPublished() {
	v.IsPublished = true
}

func (v *Video) MarkAsNotPublished() {
	v.IsPublished = false
}

func (v *Video) AddCategoryID(categoryID string) {
	v.CategoriesID = append(v.CategoriesID, categoryID)
}

func (v *Video) Validate() error {
	_, err := govalidator.ValidateStruct(v)
	if err != nil {
		return err
	}

	return nil
}
