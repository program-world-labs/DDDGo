package main

import (
	"log"

	"gitlab.com/demojira/template.git/config"
	"gitlab.com/demojira/template.git/internal/app"
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
