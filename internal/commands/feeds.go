package commands

import (
	"codingiam/gator/internal/database"
	"codingiam/gator/internal/state"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func handlerFeeds(st *state.State, cmd command) error {
	feeds, err := st.Db.ListFeeds(context.Background())
	if err != nil {
		return err
	}

	users, err := st.Db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	usersByIDs := make(map[uuid.UUID]database.User)
	for _, user := range users {
		usersByIDs[user.ID] = user
	}

	for _, feed := range feeds {
		fmt.Println(feed.Name, "-", feed.Url, "-", usersByIDs[feed.UserID].Name)
	}

	return nil
}
