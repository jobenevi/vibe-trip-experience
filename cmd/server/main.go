package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	config "github.com/jobenevi/vibe-trip-experience/configs"
	"github.com/jobenevi/vibe-trip-experience/database"
)

func main() {

	_, err := config.LoadConfig("cmd/server")
	if err != nil {
		log.Printf("Error loading config: %v", err)
	}

	database.ConnectionToDatabase()
	fmt.Println("Connected to database")

	router := gin.Default()
	router.Run()

}
