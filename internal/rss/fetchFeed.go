package rss

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(feedUrl string) (*RSSFeed, error) {

	httpClient := http.Client{
		Timeout: 10 * time.Second,
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

	rssFeed := RSSFeed{}
	err = xml.Unmarshal(respBytes, &rssFeed)
	if err != nil {
		log.Printf("error Unmarshalling feed respBytes : %v", err)
		return nil, err
	}

	return &rssFeed, nil
}
