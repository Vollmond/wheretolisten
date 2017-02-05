package main

import (
	"encoding/json"
	"net/http"

	"github.com/vollmond/wheretolisten/streamingclient"

	"gopkg.in/gin-gonic/gin.v1"
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
	_, links := json.Marshal(album.Links())
	c.JSON(http.StatusOK, gin.H{"links": json.mas})
}
