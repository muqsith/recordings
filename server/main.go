package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"recordings/handlers"
	"recordings/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/muqsith/oconf"
)

func getCWD() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

func RegisterHttpPub(router *gin.Engine) {
	cwd := getCWD()
	pathSep := string(os.PathSeparator)
	router.LoadHTMLGlob(cwd + pathSep + "build" + pathSep + "index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.Static("/static", cwd + pathSep + "build" + pathSep + "static")
}

func main() {

	var configFilePath string
	flag.StringVar(&configFilePath, "config", "", "Config file path")

	flag.Parse()

	if configFilePath == "" {
		fmt.Printf("Invalid config file path %s\n", configFilePath)
		os.Exit(1)
	}

	fmt.Println("CONFIG: ", configFilePath)

	config := oconf.GetFlatConfig(configFilePath)
	appEnv := config["env"]
	fmt.Println("env: ", appEnv)

	ginMode := gin.DebugMode
	if appEnv == "production" {
		ginMode = gin.ReleaseMode
	}
	gin.SetMode(ginMode)
	router := gin.Default()

	// register frontend only if not running in development mode
	if appEnv == "production" {
		RegisterHttpPub(router)
	}

	// apply middlewares
	middlewares.ApplyCORSMiddleware(config, router)
	middlewares.ApplyMyCustomMiddleware(router)

	// register API handlers
	handlers.Register(router)

	router.Run("localhost:8080")
}
