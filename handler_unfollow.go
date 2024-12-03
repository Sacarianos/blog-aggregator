package main

import (
	"context"
	"fmt"

	"github.com/Sacarianos/blog-aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {

	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: unfollow <feed_url>")
	}

	feedURL := cmd.Args[0]

	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    feedURL,
	})
	if err != nil {
		return fmt.Errorf("failed to unfollow feed: %w", err)
	}

	fmt.Printf("You have unfollowed the feed at '%s'.\n", feedURL)
	return nil
}
