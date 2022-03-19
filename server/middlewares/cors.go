package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ApplyCORSMiddleware(config map[string]interface{}, router *gin.Engine) {
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

  appEnv := config["env"]
  if appEnv != "production" {
    router.Use(cors.Default())
  }
}