package main

import "net/http"

func handlerReady(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status string `json:"status"`
	}
	respondWithJson(w, http.StatusOK, response{
		Status: "ok",
	})
}
