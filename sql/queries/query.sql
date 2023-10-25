-- name: CreateCategory :exec
INSERT INTO categories (id,name,description,is_active,created_at) VALUES ($1,$2,$3,$4,$5);

-- name: ListCategories :many
SELECT * FROM categories ORDER BY name LIMIT $1 OFFSET $2;

-- name: GetCategory :one
SELECT * FROM categories WHERE id = $1 LIMIT 1;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1;

-- name: RegisterVideo :exec
INSERT INTO videos (id,title,description,duration,year_launched,is_published,banner_url,video_url,categories_id,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);

-- name: ListVideos :many
SELECT * FROM videos ORDER BY title LIMIT $1 OFFSET $2;

-- name: GetVideoById :one
SELECT * FROM videos WHERE id = $1 LIMIT 1;

-- name: GetVideoByCategoryId :many
SELECT * FROM videos WHERE categories_id @> ARRAY[$1];

-- name: UpdateVideoFiles :one
UPDATE videos SET video_url = $2, banner_url = $3 WHERE id = $1 RETURNING *;

-- name: UpdateVideoIsPublished :one
UPDATE videos SET is_published = $2 WHERE id = $1 RETURNING *;