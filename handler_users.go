package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Lanrey-waju/rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters")
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	respondWithJson(w, http.StatusCreated, databaseUserToUser(user))

}

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJson(w, http.StatusOK, databaseUserToUser(user))
}

// func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request) {
// 	apiKey := r.Context().Value("apikey").(string)
// 	user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
// 	if err != nil {
// 		respondWithError(w, http.StatusNotFound, "couldn't get user")
// 		return
// 	}

// 	respondWithJson(w, http.StatusOK, databseUserToUser(user))
// }
