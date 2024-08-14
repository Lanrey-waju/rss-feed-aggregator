package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Lanrey-waju/rss-feed-aggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}
	dbURL := os.Getenv("DBURL")

	if dbURL == "" {
		log.Fatal("DBURL environment variable not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error connecting to databse: %v", err)
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	go startScraping(apiCfg.DB, 10, 60*time.Second)
	// start a new Servemux
	mux := http.NewServeMux()

	const filePathRoot = "localhost"

	mux.HandleFunc("POST /v1/users", apiCfg.handlerUsersCreate)
	mux.HandleFunc("GET /v1/users", apiCfg.middlewareAuth(apiCfg.handlerUsersGet))
	// mux.Handle("GET /v1/users", middlewareAuthorize(apiCfg.handlerUsersGet))
	mux.HandleFunc("POST /v1/feeds", apiCfg.middlewareAuth(apiCfg.handleCreateFeeds))
	mux.HandleFunc("GET /v1/feeds", apiCfg.handleGetFeeds)
	mux.HandleFunc("POST /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.handleFeedFollow))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handleDeleteFeedFollow))
	mux.HandleFunc("GET /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.handleGetUserFeedFollows))

	mux.HandleFunc("/v1/healthz", apiCfg.handlerReady)
	mux.HandleFunc("/v1/err", apiCfg.handlerError)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files on %s:%s\n", filePathRoot, port)
	log.Fatal(srv.ListenAndServe())

}
