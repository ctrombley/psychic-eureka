package service

import "github.com/gin-gonic/gin"

func initRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/locations", func(c *gin.Context) {

	})
}
