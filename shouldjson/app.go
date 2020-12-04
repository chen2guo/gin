package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Perspn struct {
	Name string `from :"name"`
	Sex  string `from :"sex"`
	Age  int    `from :"age"`
}

func main() {
	engine := gin.Default()

	engine.POST("/addstudent", func(context *gin.Context) {
		var person Perspn
		if err := context.BindJSON(&person); err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println(person.Name, person.Age, person.Sex)
		context.Writer.Write([]byte("Add " + person.Name + " ok."))

	})
	engine.Run()
}
