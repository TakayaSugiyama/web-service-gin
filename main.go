package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrance", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "John Coltrance", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	newAlbums := []album{}
	for _, a := range albums {
		if a.ID != id {
			newAlbums = append(newAlbums, a)
		}
	}
	albums = newAlbums
	c.IndentedJSON(http.StatusOK, gin.H{"message": "album is deleted"})
}

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
