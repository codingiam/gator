package commands

import (
	"codingiam/gator/internal/database"
	"codingiam/gator/internal/state"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

func handlerAddfeed(st *state.State, cmd command) error {
	if len(cmd.Args) != 2 {
		return errors.New("addfeed requires name and url")
	}

	username := st.Cfg.CurrentUserName
	name := cmd.Args[0]
	url := cmd.Args[1]

	user, err := st.Db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}

	_, err = st.Db.CreateFeed(context.Background(), database.CreateFeedParams{
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

	return nil
}
