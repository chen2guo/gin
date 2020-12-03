package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		name := context.Query("name")
		fmt.Printf("Username is %s.\n", name)

		context.Writer.Write([]byte("hello, " + name + "."))
	})

	engine.POST("/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		name := context.Query("name")
		pass := context.Query("pass")
		fmt.Printf("Username is %s; Pass is %s.\n", name, pass)

		context.Writer.Write([]byte("hello, " + name + "."))
	})

	engine.DELETE("/user/:id", func(context *gin.Context) {
		userID := context.Param("id")
		fmt.Println(userID)
		context.Writer.Write([]byte("Delete Id: " + userID + "."))
	})

	engine.Run()
}
