package main

import (
	"net/http"

	"github.com/Lanrey-waju/rss-feed-aggregator/internal/auth"
	"github.com/Lanrey-waju/rss-feed-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Couldn't get apikey")
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "user does not exist")
			return
		}

		handler(w, r, user)
	}
}

// func middlewareAuthorize(handler http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		apiKey, err := auth.GetAPIKey(r.Header)
// 		if err != nil {
// 			respondWithError(w, http.StatusUnauthorized, "couldn't get apikey")
// 			return
// 		}
// 		ctx := context.WithValue(r.Context(), "apikey", apiKey)
// 		handler.ServeHTTP(w, r.WithContext(ctx))
// 	}
// }
