package main

import (
	"time"

	"github.com/Lanrey-waju/rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type UsersFeed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}

func databaseFeedsToFeeds(feeds []database.Feed) []Feed {
	respfeeds := make([]Feed, len(feeds))
	for i, feed := range feeds {
		respfeeds[i] = databaseFeedToFeed(feed)
	}
	return respfeeds

}

func databaseUsersFeedsToUsersFeeds(usersFeed database.UsersFeed) UsersFeed {
	return UsersFeed{
		ID:        usersFeed.ID,
		CreatedAt: usersFeed.CreatedAt,
		UpdatedAt: usersFeed.UpdatedAt,
		FeedID:    usersFeed.FeedID,
		UserID:    usersFeed.UserID,
	}
}

func databaseUserFeedFollowsToFeedsFollows(feedFollows []database.UsersFeed) []UsersFeed {
	respUserFeeds := make([]UsersFeed, len(feedFollows))
	for i, userFeed := range feedFollows {
		respUserFeeds[i] = databaseUsersFeedsToUsersFeeds(userFeed)
	}
	return respUserFeeds
}
