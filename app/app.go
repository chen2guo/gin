package main

import "github.com/gin-gonic/gin"

func main() {

	engine := gin.Default()
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	engine.Run()
}
