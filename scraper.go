package main

import (
	"context"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Lanrey-waju/rss-feed-aggregator/internal/database"
)

// Rss was generated 2024-08-13 22:25:32 by https://xml-to-go.github.io/ in Ukraine.
type RSSFeed struct {
	Channel struct {
		Title         string    `xml:"title"`
		Link          string    `xml:"link"`
		Description   string    `xml:"description"`
		Generator     string    `xml:"generator"`
		Language      string    `xml:"language"`
		LastBuildDate string    `xml:"lastBuildDate"`
		Item          []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
}

// Rss was generated 2024-08-13 22:38:26 by https://xml-to-go.github.io/ in Ukraine.
// type Rss struct {
// 	Channel struct {
// 		Text          string `xml:",chardata"`
// 		Title         string `xml:"title"`
// 		Link          string `xml:"link"`
// 		Description   string `xml:"description"`
// 		Generator     string `xml:"generator"`
// 		Language      string `xml:"language"`
// 		LastBuildDate string `xml:"lastBuildDate"`
// 		Item          []struct {
// 			Title       string `xml:"title"`
// 			Link        string `xml:"link"`
// 			PubDate     string `xml:"pubDate"`
// 			Guid        string `xml:"guid"`
// 			Description string `xml:"description"`
// 		} `xml:"item"`
// 	} `xml:"channel"`
// }

func fetchFeed(url string) (*RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed RSSFeed
	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return nil, err
	}

	return &rssFeed, nil
}

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("couldn't get next feeds to fetch:", err)
			return
		}
		log.Printf("Found %v feeds  to fetch", len(feeds))

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("couldn't mark feed %s as fetched: %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(feed.Url)
	if err != nil {
		log.Printf("couldn't collect feed %s: %v", feed.Name, err)
		return
	}

	for _, item := range feedData.Channel.Item {
		log.Println("Found post:", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))

}
