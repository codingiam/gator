package main

import (
	"codingiam/gator/internal/commands"
	"codingiam/gator/internal/config"
	"codingiam/gator/internal/state"
	"log"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := state.New(&cfg)

	cmds := commands.New()

	err = cmds.Execute(&programState)
	if err != nil {
		log.Fatal(err)
	}
}
