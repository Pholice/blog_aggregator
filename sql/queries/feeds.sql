-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeed :many
SELECT * FROM feeds;

-- name: GetFeedID :one
SELECT * FROM feeds 
WHERE id = ($1);

-- name: MarkFeedFetched :one 
UPDATE feeds
SET last_fetched_at = NOW() AND updated_at = NOW()
WHERE id = ($1)
RETURNING *;

-- name: GetNextFeedToFetch :many 
SELECT * FROM feeds
ORDER BY 
CASE WHEN last_fetched_at IS NULL THEN 0
ELSE 1 END, 
last_fetched_at
LIMIT ($1);