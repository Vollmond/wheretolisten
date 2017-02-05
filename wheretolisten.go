package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vollmond/wheretolisten/streamingclient"
)

func main() {
	router := gin.Default()
	router.GET("/", proceedRequest)
	router.Run()
}

func proceedRequest(c *gin.Context) {
	albumName := c.Query("album")
	artistName := c.Query("artist")

	if artistName == "" || albumName == "" {
		c.String(http.StatusNotFound, "search params not provided")
		return
	}

	album := streamingclient.Album{Name: albumName, ArtistName: artistName}
	c.JSON(http.StatusOK, gin.H{"links": album.Links()})
}
