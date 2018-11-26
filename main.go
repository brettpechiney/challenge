package main

import (
	"log"

	"github.com/brettpechiney/challenge/challenge"
	"github.com/brettpechiney/challenge/cockroach"
	"github.com/brettpechiney/challenge/config"
	"github.com/brettpechiney/challenge/http"
)

var (
	configPaths = [...]string{".", "/apps/challenge"}
)

func main() {
	cfg, err := config.Load(configPaths[:])
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	dao, err := cockroach.NewDAO(cfg.DataSource())
	defer dao.Close()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	userRepo := challenge.NewUserRepo(dao)
	attestationRepo := challenge.NewAttestationRepo(dao)
	server := http.NewServer(userRepo, attestationRepo)
	server.Start()
}
