package controller

import (
	"net/http"

	"github.com/TakayaSugiyama/web-service-gin/model"
	"github.com/gin-gonic/gin"
)

var albums = []model.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrance", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "John Coltrance", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbums(c *gin.Context) {
	var newAlbum model.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	newAlbums := []model.Album{}
	for _, a := range albums {
		if a.ID != id {
			newAlbums = append(newAlbums, a)
		}
	}
	albums = newAlbums
	c.IndentedJSON(http.StatusOK, gin.H{"message": "album is deleted"})
}
