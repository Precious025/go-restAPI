package main

import (
	"github.com/Precious025/go-rest/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handler.GetAlbums)
	router.POST("/albums", handler.PostAlbums)
	router.POST("/albums/:id", handler.GetAlbumById)

	router.Run("localhost:8080")
}
