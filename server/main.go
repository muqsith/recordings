package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"recordings/handlers"
	"recordings/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/muqsith/oconf"
)

func RegisterHttpPub(router *gin.Engine) {
	// the files are loaded into the binary during compile time
	// the static files can be skipped based on the build env
	router.LoadHTMLGlob("../http-pub/build/index.html")
	router.GET("/", func (c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.Static("/static", "../http-pub/build/static")
}

func main() {

	var configFilePath string
	flag.StringVar(&configFilePath, "config", "", "Config file path")

	flag.Parse()

	if configFilePath == "" {
		fmt.Printf("Invalid config file path %s", configFilePath)
		os.Exit(1)
	}

	fmt.Println("CONFIG: ", configFilePath)

	config := oconf.GetFlatConfig(configFilePath)
	fmt.Println("env: ", config["env"])
	router := gin.Default()
	
	// register frontend
	RegisterHttpPub(router)

	// apply middlewares
	middlewares.ApplyCORSMiddleware(router)
	middlewares.ApplyMyCustomMiddleware(router)

	// register API handlers
	handlers.Register(router)

	router.Run("localhost:8080")
}