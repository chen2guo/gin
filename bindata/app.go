package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Student struct {
	Name    string `form:"name"`
	Classes string `form:"classes"`
}

func main() {

	engine := gin.Default()

	engine.GET("/hello", func(c *gin.Context) {
		fmt.Println(c.FullPath())

		var student Student

		if err := c.ShouldBindQuery(&student); err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Println(student.Name)
		fmt.Println(student.Classes)
		c.Writer.Write([]byte("Hello, " + student.Name + "."))

	})
	engine.Run()
}
