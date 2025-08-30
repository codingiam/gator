package commands

import (
	"codingiam/gator/internal/database"
	"codingiam/gator/internal/state"
	"context"
	"fmt"
)

func handlerFollowing(st *state.State, cmd command, user database.User) error {
	follows, err := st.Db.FeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return err
	}

	for _, follow := range follows {
		fmt.Println(follow.FeedName)
	}

	return nil
}
