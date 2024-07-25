package api

import "github.com/gin-gonic/gin"

func Ping(router *gin.RouterGroup) {
	handler := func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	}

	router.GET("/ping", handler)
}
