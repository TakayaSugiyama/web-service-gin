package main

import (
	"log"

	"github.com/TakayaSugiyama/web-service-gin/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := routes.InitRoutes()
	port := routes.SetPortandGinMode()
	router.Run(":" + port)
}
