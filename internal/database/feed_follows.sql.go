// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feed_follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedsFollow = `-- name: CreateFeedsFollow :one
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
RETURNING id, created_at, updated_at, feed_id, user_id
`

type CreateFeedsFollowParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedID    uuid.UUID
	UserID    uuid.UUID
}

func (q *Queries) CreateFeedsFollow(ctx context.Context, arg CreateFeedsFollowParams) (UsersFeed, error) {
	row := q.db.QueryRowContext(ctx, createFeedsFollow,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.FeedID,
		arg.UserID,
	)
	var i UsersFeed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FeedID,
		&i.UserID,
	)
	return i, err
}

const deleteFeedFollow = `-- name: DeleteFeedFollow :exec
DELETE FROM users_feeds WHERE id = $1 and user_id = $2
`

type DeleteFeedFollowParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) DeleteFeedFollow(ctx context.Context, arg DeleteFeedFollowParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollow, arg.ID, arg.UserID)
	return err
}

const getUserFeedFollows = `-- name: GetUserFeedFollows :many
SELECT id, created_at, updated_at, feed_id, user_id
FROM users_feeds
WHERE user_id = $1
`

func (q *Queries) GetUserFeedFollows(ctx context.Context, userID uuid.UUID) ([]UsersFeed, error) {
	rows, err := q.db.QueryContext(ctx, getUserFeedFollows, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UsersFeed
	for rows.Next() {
		var i UsersFeed
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.FeedID,
			&i.UserID,
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
