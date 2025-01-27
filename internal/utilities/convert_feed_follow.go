package utilities

import (
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/database"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/entity"
)

func DatabaseFeedFollowToFeedFollow(feedFollow database.FeedsFollow) entity.FeedFollow {
	return entity.FeedFollow{
		ID:        feedFollow.ID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		UserID:    feedFollow.UserID,
		FeedID:    feedFollow.FeedID,
	}
}

func DatabaseFeedFollowsToFeedFollows(feedFollows []database.FeedsFollow) []entity.FeedFollow {
	feedFollowsResponse := make([]entity.FeedFollow, len(feedFollows))
	for i, feedFollow := range feedFollows {
		feedFollowsResponse[i] = entity.FeedFollow(feedFollow)
	}
	return feedFollowsResponse
}
