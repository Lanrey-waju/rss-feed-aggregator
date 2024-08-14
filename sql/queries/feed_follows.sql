-- name: CreateFeedsFollow :one
INSERT INTO users_feeds
    (
    id,
    created_at,
    updated_at,
    feed_id,
    user_id
    )
VALUES
    (
        $1, $2, $3, $4, $5
    )
RETURNING *;

-- name: DeleteFeedFollow :exec
DELETE FROM users_feeds WHERE id = $1 and user_id = $2;

-- name: GetUserFeedFollows :many
SELECT *
FROM users_feeds
WHERE user_id = $1;