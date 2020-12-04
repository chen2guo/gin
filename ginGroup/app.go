package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()

	routerGroup := engine.Group("/user")

	routerGroup.POST("/register", func(context *gin.Context) {

		context.Writer.WriteString(context.FullPath())
	})

	routerGroup.POST("/login", func(context *gin.Context) {

		context.Writer.WriteString(context.FullPath())
	})

	routerGroup.DELETE("/:id", func(context *gin.Context) {

		context.Writer.WriteString("Delete " + context.Param("id") + "  Ok.")
	})

	engine.Run()
}
