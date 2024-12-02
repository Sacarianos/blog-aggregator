-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;


-- name: ListFeeds :many
SELECT 
    feeds.name AS feed_name,
    feeds.url AS feed_url,
    users.name AS user_name
FROM 
    feeds
JOIN 
    users 
ON 
    feeds.user_id = users.id
ORDER BY 
    feeds.created_at;



-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE url = $1;

-- name: GetFeeds :many
SELECT * FROM feeds;