package main

import (
	"github.com/TakayaSugiyama/web-service-gin/routes"
)

func main() {
	router := routes.InitRoutes()
	port := routes.SetPortandGinMode()
	router.Run(":" + port)
}
