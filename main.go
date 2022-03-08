package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func setPortandGinMode() string {
	if value, ok := os.LookupEnv("PORT"); ok {
		gin.SetMode(gin.ReleaseMode)
		return value
	}
	return "8080"
}

func main() {
	port := setPortandGinMode()
	initRoutes()
	router.Run(":" + port)
}
