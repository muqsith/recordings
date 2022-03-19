package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MyCustomMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(">>>>>> Remote address: ", c.Request.RemoteAddr)
	}
}


func ApplyMyCustomMiddleware(router *gin.Engine) {
	router.Use(MyCustomMiddleware())
}