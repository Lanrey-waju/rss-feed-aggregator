package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Lanrey-waju/rss-feed-aggregator/internal/database"
)

func (cfg *apiConfig) handleGetPostsByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	// Get the "N" query parameter from URL
	numOfPostsString := r.URL.Query().Get("n")
	numOfPostsInt, err := strconv.Atoi(numOfPostsString)
	log.Println(numOfPostsInt)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get specified number of posts to get")
		return
	}
	// Get "N" number rows from posts table
	userPosts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(numOfPostsInt),
	})
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusNotFound, "couldn't retrieve posts")
		return
	}

	respondWithJson(w, http.StatusOK, databasePostsToPosts(userPosts))

}
