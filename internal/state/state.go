package state

import (
	"codingiam/gator/internal/config"
	"codingiam/gator/internal/database"
)

type State struct {
	Db  *database.Queries
	Cfg *config.Config
}

func New(db *database.Queries, cfg *config.Config) State {
	return State{Db: db, Cfg: cfg}
}
