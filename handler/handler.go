package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Playboy", Artist: "David Adeleke", Price: 40.99},
	{ID: "2", Title: "Italy", Artist: "Buju", Price: 17.99},
	{ID: "3", Title: "Peru", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

//assign a parameter

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")

	// loop through a slice containing structs of album with the ID value
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
