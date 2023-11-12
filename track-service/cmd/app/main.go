package main

import (
	"log"
	"track-service/internal/app"
	"track-service/pkg/utils/config"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
