package utilities

import (
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/database"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/entity"
)

func DatabaseFeedToFeed(feed database.Feed) entity.Feed {
	return entity.Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}

func DatabaseFeedsToFeeds(feeds []database.Feed) []entity.Feed {
	feedsResponse := make([]entity.Feed, len(feeds))
	for i, feed := range feeds {
		feedsResponse[i] = DatabaseFeedToFeed(feed)
	}
	return feedsResponse
}
