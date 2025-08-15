package commands

import (
	"codingiam/gator/internal/state"
	"context"
	"fmt"
)

func handlerUsers(st *state.State, _ command) error {
	users, err := st.Db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		var suffix string
		if user.Name == st.Cfg.CurrentUserName {
			suffix = "(current)"
		}
		fmt.Println("*", user.Name, suffix)
	}

	return nil
}
