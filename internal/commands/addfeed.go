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

func handlerAddfeed(st *state.State, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return errors.New("addfeed requires name and url")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := st.Db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
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
