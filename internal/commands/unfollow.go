package commands

import (
	"codingiam/gator/internal/database"
	"codingiam/gator/internal/state"
	"context"
	"errors"
	"fmt"
)

func handlerUnfollow(st *state.State, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("unfollow requires url")
	}

	url := cmd.Args[0]

	feed, err := st.Db.FeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	err = st.Db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return err
	}

	fmt.Printf("%s no longer following URL %s\n", user.Name, feed.Url)

	return nil
}
