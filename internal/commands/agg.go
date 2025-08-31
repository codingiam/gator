package commands

import (
	"codingiam/gator/internal/feed"
	"codingiam/gator/internal/state"
	"errors"
	"fmt"
	"time"
)

func handlerAgg(st *state.State, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("agg requires interval")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		ffeed, err := feed.ScrapeFeeds(st)
		if err != nil {
			return err
		}
		fmt.Println(ffeed)
	}

	return nil
}
