-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING *;


-- name: GetFeeds :many
SELECT
    feeds.name,
    feeds.url,
    users.name AS user_name
FROM feeds
LEFT JOIN users
    ON feeds.user_id = users.id;

-- name: GetFeed :one
SELECT
    id,
    created_at,
    updated_at,
    name,
    url,
    user_id
FROM feeds
WHERE url = $1 LIMIT 1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET
    last_fetched_at = $1,
    updated_at = $2
WHERE feeds.id = $3;

-- name: GetNextFeedToFetch :one
SELECT
    id,
    last_fetched_at,
    url
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST LIMIT 1;
