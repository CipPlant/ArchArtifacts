package main

import (
	"akim/internal/app"
	"akim/internal/config"
	"log"
)

func main() {
	cfg := config.MustLoad()
	if err := app.Run(cfg); err != nil {
		log.Fatal("can not run app, error:", err)
	}
}
