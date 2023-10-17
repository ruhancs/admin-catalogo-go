-- name: CreateCategory :exec
INSERT INTO categories (id,name,description,is_active,created_at) VALUES ($1,$2,$3,$4,$5);

-- name: ListCategories :many
SELECT * FROM categories ORDER BY name LIMIT $1 OFFSET $2;

-- name: GetCategory :one
SELECT * FROM categories WHERE id = $1 LIMIT 1;

-- name: DeleteAccount :exec
DELETE FROM categories WHERE id = $1;