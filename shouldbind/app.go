package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Register struct {
	Username string
	Phone    string
	Password string
}

func main() {
	engine := gin.Default()

	engine.POST("/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var register Register
		if err := context.ShouldBind(&register); err != nil {
			log.Fatal(err.Error())

			return
		}

		fmt.Println(register.Username, register.Password, register.Phone)
		context.Writer.Write([]byte(register.Username + "  login OK."))

	})

	engine.Run()

}
