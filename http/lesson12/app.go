package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./html/login.html")

	r.GET("/login", login)

	r.Run()

}

func login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)

}
