package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func hello(context *gin.Context) {
	fmt.Println(context.FullPath())

	name := context.Query("name")
	fmt.Printf("Username is %s.\n", name)

	context.Writer.Write([]byte("hello, " + name + "."))
}

func login(context *gin.Context) {
	fmt.Println(context.FullPath())

	name := context.Query("name")
	pass := context.Query("pass")
	fmt.Printf("Username is %s; Pass is %s.\n", name, pass)

	context.Writer.Write([]byte("hello, " + name + "."))
}

func user(context *gin.Context) {
	userID := context.Param("id")
	fmt.Println(userID)
	context.Writer.Write([]byte("Delete Id: " + userID + "."))
}

func main() {
	engine := gin.Default()
	engine.GET("/hello", hello)
	engine.POST("/login", login)
	engine.DELETE("/user/:id", user)
	engine.Run()
}
