package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Lanrey-waju/rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters")
		return
	}

	usersFeed, err := cfg.DB.CreateFeedsFollow(r.Context(), database.CreateFeedsFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create feed follow")
		return
	}

	respondWithJson(w, http.StatusCreated, databaseUsersFeedsToUsersFeeds(usersFeed))
}

func (cfg *apiConfig) handleDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDString := r.PathValue("feedFollowID")

	feedFollowID, err := uuid.Parse(feedFollowIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "couldn't parse UUID string")
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		ID:     feedFollowID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error deleting resource")
		return
	}

	resp := struct {
		Status string
	}{
		Status: "Deleted",
	}
	respondWithJson(w, http.StatusOK, resp)

}

func (cfg *apiConfig) handleGetUserFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := cfg.DB.GetUserFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "feed follows not found")
		return
	}

	respondWithJson(w, http.StatusOK, databaseUserFeedFollowsToFeedsFollows(feedFollows))

}
