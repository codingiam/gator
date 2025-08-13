package commands

import (
	"codingiam/gator/internal/state"
	"errors"
)

func handlerLogin(st *state.State, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("login requires username")
	}
	return st.Cfg.SetUser(cmd.Args[0])
}
