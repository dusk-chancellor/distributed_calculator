package main

import (
	"log"

	"github.com/dusk-chancellor/distributed_calculator/sso/internal/config"
	"github.com/dusk-chancellor/distributed_calculator/sso/pkg/zaplog"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	logger, err := zaplog.New(cfg.Server.Env)
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
		return
	}
	defer logger.Sync()

	logger.Info("Configuring SSO server...")

	
}
