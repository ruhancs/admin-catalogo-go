package repository

import (
	"admin-catalogo-go/internal/application/dto"
	"admin-catalogo-go/internal/domain/entity"
	"admin-catalogo-go/internal/infra/db"
	"context"
	"database/sql"
)

type VideoRepository struct {
	DB *sql.DB
	Queries *db.Queries
}

func NewVideoRepository(database *sql.DB) *VideoRepository {
	return &VideoRepository{
		DB: database,
		Queries: db.New(database),
	}
}

func (repository *VideoRepository) Insert(ctx context.Context,video *entity.Video) error{
	err := repository.Queries.RegisterVideo(ctx,db.RegisterVideoParams{
		ID: video.ID,
		Title: video.Title,
		Description: sql.NullString{String: video.Description, Valid: true},
		Duration: sql.NullInt64{Int64: int64(video.Duration), Valid: true},
		IsPublished: video.IsPublished,
		BannerUrl: sql.NullString{String: video.BannerUrl, Valid: true},
		VideoUrl: sql.NullString{String: video.VideoUrl, Valid: true},
		CategoriesID: video.CategoriesID,
		CreatedAt: video.CreatedAt,
	})

	if err != nil {
		return err
	}

	return nil
}

func (repository *VideoRepository) ListVideos(ctx context.Context, input dto.ListVideoInputDto) ([]entity.Video,error) {
	offset := (input.Page - 1) * input.PerPage
	videosModel,err := repository.Queries.ListVideos(ctx,db.ListVideosParams{
		Limit: int32(input.PerPage),
		Offset: int32(offset),
	})
	if err != nil {
		return nil,nil
	}

	var videosEntity []entity.Video
	for _,videoModel := range videosModel {
		video := entity.Video{
			ID: videoModel.ID,
			Title: videoModel.Title,
			Description: videoModel.Description.String,
			YearLaunched: int(videoModel.YearLaunched),
			Duration: float64(videoModel.Duration.Int64),
			IsPublished: videoModel.IsPublished,
			BannerUrl: videoModel.BannerUrl.String,
			VideoUrl: videoModel.VideoUrl.String,
			CategoriesID: videoModel.CategoriesID,
			CreatedAt: videoModel.CreatedAt,
		}
		videosEntity = append(videosEntity, video)
	}

	return videosEntity,nil
}

func (repository *VideoRepository) GetVideoByID(ctx context.Context, id string) (entity.Video,error){
	videoModel,err := repository.Queries.GetVideoById(ctx,id)
	if err != nil {
		return entity.Video{},err
	}

	video := entity.Video{
		ID: videoModel.ID,
		Title: videoModel.Title,
		Description: videoModel.Description.String,
		YearLaunched: int(videoModel.YearLaunched),
		Duration: float64(videoModel.Duration.Int64),
		IsPublished: videoModel.IsPublished,
		BannerUrl: videoModel.BannerUrl.String,
		VideoUrl: videoModel.VideoUrl.String,
		CategoriesID: videoModel.CategoriesID,
		CreatedAt: videoModel.CreatedAt,
	}

	return video,nil
}

func (repository *VideoRepository) GetVideoByCategoryID(ctx context.Context, categoryID string) ([]entity.Video,error) {
	categoriesID := []string{categoryID}
	videosModel,err := repository.Queries.GetVideoByCategoryId(ctx, categoriesID)
	if err != nil {
		return nil,err
	}

	var videos []entity.Video
	for _,video := range videosModel {
		videoEntity := entity.Video{
			ID: video.ID,
			Title: video.Title,
			Description: video.Description.String,
			YearLaunched: int(video.YearLaunched),
			Duration: float64(video.Duration.Int64),
			IsPublished: video.IsPublished,
			BannerUrl: video.BannerUrl.String,
			VideoUrl: video.VideoUrl.String,
			CategoriesID: video.CategoriesID,
			CreatedAt: video.CreatedAt,
		}
		videos = append(videos, videoEntity)
	}

	return videos,nil
}

func (repository *VideoRepository) UpdateFiles(ctx context.Context,id string, videoUrl,bannerUrl string) (entity.Video,error){
	videoModel,err := repository.Queries.UpdateVideoFiles(ctx, db.UpdateVideoFilesParams{
		ID: id,
		VideoUrl: sql.NullString{String: videoUrl,Valid: true},
		BannerUrl: sql.NullString{String: bannerUrl,Valid: true},
	})
	if err != nil {
		return entity.Video{},err
	}

	videoEntity := entity.Video{
		ID: videoModel.ID,
		Title: videoModel.Title,
		Description: videoModel.Description.String,
		YearLaunched: int(videoModel.YearLaunched),
		Duration: float64(videoModel.Duration.Int64),
		VideoUrl: videoModel.VideoUrl.String,
		BannerUrl: videoModel.BannerUrl.String,
		CategoriesID: videoModel.CategoriesID,
		IsPublished: videoModel.IsPublished,
		CreatedAt: videoModel.CreatedAt,
	}

	return videoEntity,nil
}

func (repository *VideoRepository) UpdatePublishState(ctx context.Context,id string, input dto.UpdateVideoPublishStateInputDto) error{
	_,err := repository.Queries.UpdateVideoIsPublished(ctx,db.UpdateVideoIsPublishedParams{
		ID: id,
		IsPublished: input.IsPublished,
	})
	if err != nil {
		return err
	}
	return nil
}