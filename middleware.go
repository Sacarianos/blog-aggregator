package main

import (
	"context"
	"fmt"

	"github.com/Sacarianos/blog-aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		// Ensure the user is logged in
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("user must be logged in: %w", err)
		}

		// Call the wrapped handler with the user
		return handler(s, cmd, user)
	}
}
