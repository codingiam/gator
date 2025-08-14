package main

import (
	"codingiam/gator/internal/commands"
	"codingiam/gator/internal/config"
	"codingiam/gator/internal/database"
	"codingiam/gator/internal/state"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("db error: %v", err)
	}

	dbQueries := database.New(db)

	programState := state.New(dbQueries, &cfg)

	cmds := commands.New()

	err = cmds.Execute(&programState)
	if err != nil {
		log.Fatal(err)
	}
}
