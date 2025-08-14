package commands

import (
	"codingiam/gator/internal/state"
	"context"
	"errors"
)

func handlerLogin(st *state.State, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("login requires username")
	}

	user := cmd.Args[0]

	_, err := st.Db.GetUser(context.Background(), user)
	if err != nil {
		return errors.New("user does not exists")
	}

	return st.Cfg.SetUser(user)
}
