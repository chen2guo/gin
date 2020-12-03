package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.Handle("GET", "/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		name := context.DefaultQuery("name", "hello")

		context.Writer.Write([]byte("hello ," + name))
	})

	engine.Handle("POST", "/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		name := context.Query("name")
		pass := context.Query("pass")

		fmt.Printf("username: %s; password: %s.\n", name, pass)

		context.Writer.Write([]byte(name + " login OK."))
	})

	engine.Handle("DELETE", "/user/:id", func(context *gin.Context) {
		userID := context.Param("id")
		fmt.Println(userID)
		context.Writer.Write([]byte("Delete Id: " + userID + "."))
	})

	engine.Run()
}
