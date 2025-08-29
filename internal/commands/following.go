package commands

import (
	"codingiam/gator/internal/state"
	"context"
	"fmt"
)

func handlerFollowing(st *state.State, cmd command) error {
	name := st.Cfg.CurrentUserName
	follows, err := st.Db.FeedFollowsForUser(context.Background(), name)
	if err != nil {
		return err
	}

	for _, follow := range follows {
		fmt.Println(follow.FeedName)
	}

	return nil
}
