-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, feed_id, description, published_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostByUser :many 
SELECT 
    p.id, p.title, p.url, p.description, p.published_at, p.created_at, p.updated_at 
FROM 
    posts p 
JOIN 
    feeds f ON p.feed_id = f.id 
WHERE 
    f.user_id = ($1)
ORDER BY 
    p.published_at DESC 
LIMIT ($2);