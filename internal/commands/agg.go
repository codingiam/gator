package commands

import (
	"codingiam/gator/internal/feed"
	"codingiam/gator/internal/state"
	"context"
	"fmt"
)

func handlerAgg(st *state.State, _ command) error {
	feed, err := feed.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}
