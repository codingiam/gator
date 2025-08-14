package commands

import (
	"codingiam/gator/internal/database"
	"codingiam/gator/internal/state"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(st *state.State, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("register requires name")
	}

	user := cmd.Args[0]

	_, err := st.Db.GetUser(context.Background(), user)
	if err == nil {
		return errors.New("user already exists")
	}

	_, err = st.Db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      user,
	})
	if err != nil {
		return err
	}

	err = st.Cfg.SetUser(user)
	if err != nil {
		return err
	}

	return nil
}
