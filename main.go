package main

import (
	"github.com/brettpechiney/challenge/data"
	"log"

	"github.com/brettpechiney/challenge/config"
)

var (
	// The paths where Challenge will look for a .toml configuration file.
	configPaths = [...]string{".", "/apps/challenge"}
)

func main() {
	cfg, err := config.Load(configPaths[:])
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := data.New(cfg)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	db.Close()
}
