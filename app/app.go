package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	engine := gin.Default()
	engine.GET("/hello", func(c *gin.Context) {
		fmt.Println("GET URL is: ", c.FullPath())
		c.Writer.Write([]byte("Hell, Gin.\n"))
	})

	engine.GET("/ping", func(c *gin.Context) {
		fmt.Println("GET URL is: ", c.FullPath())
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := engine.Run("0.0.0.0:8080"); err != nil {
		log.Fatal(err.Error())
	}
	//engine.Run()
}
