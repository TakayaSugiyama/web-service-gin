package controller

import (
	"net/http"

	"github.com/TakayaSugiyama/web-service-gin/model"
	"github.com/gin-gonic/gin"
)

func TeamIndex(c *gin.Context) {
	teams, _ := model.GetTeams()
	c.IndentedJSON(http.StatusOK, teams)
}

//func TeamCreate(c *gin.Context) {
//var newAlbum model.Album
//if err := c.BindJSON(&newAlbum); err != nil {
//return
//}
//albums = append(albums, newAlbum)
//c.IndentedJSON(http.StatusCreated, newAlbum)
//}

//func TeamShow(c *gin.Context) {
//id := c.Param("id")
//for _, a := range albums {
//if a.ID == id {
//c.IndentedJSON(http.StatusOK, a)
//return
//}
//}
//c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
//}

//func TeamDestroy(c *gin.Context) {
//id := c.Param("id")

//newAlbums := []model.Album{}
//for _, a := range albums {
//if a.ID != id {
//newAlbums = append(newAlbums, a)
//}
//}
//albums = newAlbums
//c.IndentedJSON(http.StatusOK, gin.H{"message": "album is deleted"})
//}
