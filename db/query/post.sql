-- name: CreatePost :one
INSERT INTO posts (
    owner_id,
    type,
    is_root_opinion,
    topic,
    description,
    caption,
    topic_id,
    set_id,
    category,
    base_opinion_id,
    post_image_url,
    link
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
) RETURNING *;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1  LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY timestamp
LIMIT $1
OFFSET $2;

-- name: UpdatePost :one
UPDATE posts
SET votes = $2
WHERE id = $1
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;