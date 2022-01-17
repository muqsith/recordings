package main

import (
	"net/http"
	"recordings/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterHttpPub(router *gin.Engine) {
	// the files are loaded into the binary during compile time
	// the static files can be skipped based on the build env
	router.LoadHTMLGlob("http-pub/build/index.html")
	router.GET("/index", func (c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.Static("/static", "http-pub/build/static")
}

func main() {
	router := gin.Default()
	RegisterHttpPub(router)
	handlers.Register(router)
	router.Run("localhost:8080")
}