package commands

import (
	"codingiam/gator/internal/database"
	"codingiam/gator/internal/state"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(st *state.State, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("follow requires url")
	}

	url := cmd.Args[0]

	feed, err := st.Db.FeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	_, err = st.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%s following URL %s\n", user.Name, feed.Url)

	return nil
}
