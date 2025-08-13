package state

import "codingiam/gator/internal/config"

type State struct {
	Cfg *config.Config
}

func New(cfg *config.Config) State {
	return State{Cfg: cfg}
}
