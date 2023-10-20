package repository

import (
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
		Banner: sql.NullString{String: video.Banner, Valid: true},
		CreatedAt: video.CreatedAt,
	})

	if err != nil {
		return err
	}

	return nil
}