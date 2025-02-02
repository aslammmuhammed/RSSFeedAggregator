package rss

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/aslammmuhammed/RSSFeedAggregator/internal/database"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/entity"
	"github.com/google/uuid"
)

func StartScraping(apiCfg *entity.ApiCfg) {
	log.Printf("Staring to fetch feeds every %v seconds in %d go routines", apiCfg.ScrapeInterval, apiCfg.ScrapeConcurrency)
	concurrency := apiCfg.ScrapeConcurrency
	timeBetweenRequest := time.Duration(time.Second * time.Duration(apiCfg.ScrapeInterval))
	ticker := time.NewTicker(timeBetweenRequest)

	// infinte loop executed on every 'timeBetweenRequest'
	// no conditions added , to execute the loop on first time
	for ; ; <-ticker.C {
		feedsToFetch, err := apiCfg.DB.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("error executing GetNextFeedsToFetch query: %v", err)
			continue
		}
		log.Printf("found %d feeds to fetch", len(feedsToFetch))

		wg := &sync.WaitGroup{}
		for _, feed := range feedsToFetch {
			wg.Add(1)
			go scrapeFeed(apiCfg, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(apiCfg *entity.ApiCfg, wg *sync.WaitGroup, feedToFetch database.Feed) {
	defer wg.Done()
	_, err := apiCfg.DB.MarkFeedAsFetched(context.Background(), feedToFetch.ID)
	if err != nil {
		log.Printf("error executing MarkFeedAsFetched query: %v", err)
		return
	}

	rssFeed, err := fetchFeed(feedToFetch.Url, apiCfg.ScrapeTimeout)
	if err != nil {
		log.Printf("couldn't fetch feed %s error: %v", feedToFetch.Name, err)
		return
	}

	for _, post := range rssFeed.Channel.Item {

		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, post.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		_, err := apiCfg.DB.CreatePosts(context.Background(), database.CreatePostsParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title:     post.Title,
			Url:       post.Link,
			Description: sql.NullString{
				String: post.Description,
				Valid:  true,
			},
			PublishedAt: publishedAt,
			FeedID:      feedToFetch.ID,
		})
		if err != nil {
			//suppress error log for
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("couldn't execute CreatePosts query: %v", err)
			continue
		}
	}
	log.Printf("fetched feed %s with %d posts", feedToFetch.Name, len(rssFeed.Channel.Item))
}
