package main

import "net/http"

func (cfg *apiConfig) handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
