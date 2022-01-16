package handlers

import (
	"net/http"
	"recordings/dal"
	"recordings/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	return router
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	albums, _ := dal.GetAll()
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum model.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
			return
	}

	// Add the new album to the slice.
	dal.AddAlbum(newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.ParseInt(id, 10, 64)

	// an album whose ID value matches the parameter.
	album, err := dal.AlbumByID(int64(i))
	if err == nil {
		c.IndentedJSON(http.StatusOK, album)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}