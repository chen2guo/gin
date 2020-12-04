package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Respon struct {
	Code    int
	Message string
	Data    interface{}
}

func main() {
	engine := gin.Default()

	engine.GET("/hello", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"Code":    1,
			"Message": "Ok",
			"Data":    context.FullPath(),
		})
	})

	engine.GET("/world", func(context *gin.Context) {
		resp := Respon{Code: 1, Message: "OK", Data: context.FullPath()}
		context.JSON(200, &resp)
	})

	engine.LoadHTMLFiles("./html/index.html")
	//engine.LoadHTMLGlob("./html/*")
	engine.Static("/img", "./img")

	engine.GET("/html", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"FullPath": "Hello  " + context.FullPath(),
		})
	})

	engine.Run()
}
