package main

import (
	"log"

	"github.com/program-world-labs/DDDGo/config"
	"github.com/program-world-labs/DDDGo/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
