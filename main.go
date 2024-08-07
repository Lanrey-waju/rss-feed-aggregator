package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")

	// start a new Servemux
	mux := http.NewServeMux()

	const filePathRoot = "localhost"

	mux.HandleFunc("/v1/healthz", handlerReady)
	mux.HandleFunc("/v1/err", handlerError)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files on %s:%s\n", filePathRoot, port)
	log.Fatal(srv.ListenAndServe())

}
