package main

import (
	"artist-service/internal/app"
	"artist-service/pkg/utils/config"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
