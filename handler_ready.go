package main

import "net/http"

func (cfg *apiConfig) handlerReady(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status string `json:"status"`
	}
	respondWithJson(w, http.StatusOK, response{
		Status: "ok",
	})
}
