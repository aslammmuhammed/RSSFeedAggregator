package utilities

import (
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/database"
	"github.com/aslammmuhammed/RSSFeedAggregator/internal/entity"
)

func databasePostToPost(dbPost database.Post) entity.Post {
	return entity.Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Url:         dbPost.Url,
		Description: dbPost.Description,
		PublishedAt: dbPost.PublishedAt,
		FeedID:      dbPost.FeedID,
	}
}

func DatabasePostsToPosts(dbPost []database.Post) []entity.Post {
	posts := make([]entity.Post, len(dbPost))
	for i, post := range dbPost {
		posts[i] = databasePostToPost(post)
	}
	return posts
}
