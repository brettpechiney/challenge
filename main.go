package main

import (
	"log"

	"github.com/brettpechiney/challenge/config"
	"github.com/brettpechiney/challenge/database"
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

	dao, err := database.NewDAO(cfg)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	dao.Close()
}
