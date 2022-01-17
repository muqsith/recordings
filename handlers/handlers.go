package handlers

import (
	"net/http"
	"recordings/dal"
	"recordings/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine)  {
	// api
	router.GET("/api/albums", getAlbums)
	router.GET("/api/albums/:id", getAlbumByID)
	router.POST("/api/albums", postAlbums)
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	albums, _ := dal.GetAll()
	c.IndentedJSON(http.StatusOK, albums)
}


/*
	Post request:
	curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
*/

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