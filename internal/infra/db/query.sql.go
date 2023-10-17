// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createCategory = `-- name: CreateCategory :exec
INSERT INTO categories (id,name,description,is_active,created_at) VALUES ($1,$2,$3,$4,$5)
`

type CreateCategoryParams struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	IsActive    bool           `json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, createCategory,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.IsActive,
		arg.CreatedAt,
	)
	return err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM categories WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getCategory = `-- name: GetCategory :one
SELECT id, name, description, is_active, created_at FROM categories WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCategory(ctx context.Context, id string) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategory, id)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.IsActive,
		&i.CreatedAt,
	)
	return i, err
}

const listCategories = `-- name: ListCategories :many
SELECT id, name, description, is_active, created_at FROM categories ORDER BY name LIMIT $1 OFFSET $2
`

type ListCategoriesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCategories(ctx context.Context, arg ListCategoriesParams) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, listCategories, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.IsActive,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}