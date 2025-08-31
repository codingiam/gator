package commands

import (
	"codingiam/gator/internal/database"
	"codingiam/gator/internal/state"
	"context"
	"errors"
	"os"
)

type Commands struct {
	registeredCommands map[string]func(*state.State, command) error
}

func New() Commands {
	return Commands{
		registeredCommands: map[string]func(*state.State, command) error{
			"login":     handlerLogin,
			"register":  handlerRegister,
			"reset":     handlerReset,
			"users":     handlerUsers,
			"agg":       handlerAgg,
			"addfeed":   middlewareLoggedIn(handlerAddfeed),
			"feeds":     handlerFeeds,
			"follow":    middlewareLoggedIn(handlerFollow),
			"following": middlewareLoggedIn(handlerFollowing),
			"unfollow":  middlewareLoggedIn(handlerUnfollow),
		},
	}
}

func (c Commands) Execute(st *state.State) error {
	args := os.Args

	if len(args) < 2 {
		return errors.New("no cmd given")
	}

	cmdName := args[1]
	cmdArgs := args[2:]

	return c.run(st, command{Name: cmdName, Args: cmdArgs})
}

func (c Commands) run(s *state.State, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}

func middlewareLoggedIn(handler func(s *state.State, cmd command, user database.User) error) func(*state.State, command) error {
	return func(s *state.State, cmd command) error {
		user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
