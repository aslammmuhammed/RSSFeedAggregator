package rss

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/aslammmuhammed/RSSFeedAggregator/internal/entity"
)

func fetchFeed(feedUrl string, timeOut int) (*entity.RSSFeed, error) {

	httpClient := http.Client{
		Timeout: time.Duration(timeOut) * time.Second,
	}

	resp, err := httpClient.Get(feedUrl)
	if err != nil {
		log.Printf("error fecthing feed %s : %v", feedUrl, err)
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error converting feed resp to respBytes: %v", err)
		return nil, err
	}

	rssFeed := entity.RSSFeed{}
	err = xml.Unmarshal(respBytes, &rssFeed)
	if err != nil {
		log.Printf("error Unmarshalling feed respBytes : %v", err)
		return nil, err
	}

	return &rssFeed, nil
}
