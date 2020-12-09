package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	fmt.Println("Index ...")
	c.JSON(http.StatusOK, gin.H{
		"mes": "index",
	})
}
func m1(c *gin.Context) {
	fmt.Println("M1 in ...")
	start := time.Now()
	c.Next()
	cost := time.Since(start)
	fmt.Printf("cost: %v.\n", cost)
	//c.Abort()
	fmt.Println("M1 out  ...")
}

func main() {
	r := gin.Default()

	r.Use(m1)

	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mes": "shop",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mes": "shop",
		})
	})
	r.Run()
}
