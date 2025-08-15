package commands

import (
	"codingiam/gator/internal/state"
	"errors"
	"os"
)

type Commands struct {
	registeredCommands map[string]func(*state.State, command) error
}

func New() Commands {
	return Commands{
		registeredCommands: map[string]func(*state.State, command) error{
			"login":    handlerLogin,
			"register": handlerRegister,
			"reset":    handlerReset,
			"users":    handlerUsers,
			"agg":      handlerAgg,
			"addfeed":  handlerAddfeed,
			"feeds":    handlerFeeds,
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
