package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ApplyCORSMiddleware(router *gin.Engine) {
	// router.Use(cors.New(cors.Config{
  //   AllowOrigins:     []string{"https://foo.com"},
  //   AllowMethods:     []string{"PUT", "PATCH"},
  //   AllowHeaders:     []string{"Origin"},
  //   ExposeHeaders:    []string{"Content-Length"},
  //   AllowCredentials: true,
  //   AllowOriginFunc: func(origin string) bool {
  //     return origin == "https://github.com"
  //   },
  //   MaxAge: 12 * time.Hour,
  // }))

	// allow everything
	router.Use(cors.Default())
}