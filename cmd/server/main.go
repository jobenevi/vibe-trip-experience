package main

import (
	"log"

	config "github.com/jobenevi/vibe-trip-experience/configs"
	"github.com/jobenevi/vibe-trip-experience/internal/database"
	"github.com/jobenevi/vibe-trip-experience/internal/routes"
)

func main() {
	_, err := config.LoadConfig("cmd/server")
	if err != nil {
		log.Printf("Error loading config: %v", err)
	}

	database.ConnectionToDatabase()
	routes.HandlerRequests()
}
