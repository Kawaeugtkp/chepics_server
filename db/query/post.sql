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