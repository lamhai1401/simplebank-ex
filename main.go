package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Go Registration API updated dockerhub, deploy to heroku",
		})
	})

	r.POST("/register", func(c *gin.Context) {
		fullname := c.PostForm("fullname")
		phone := c.PostForm("phone")
		email := c.PostForm("email")

		c.JSON(200, gin.H{
			"status":   "success",
			"message":  "User details successfully posted",
			"fullname": fullname,
			"phone":    phone,
			"email":    email,
		})
	})
	_ = r.Run()

	factorial(4, func(result int) {
		fmt.Println("result", result)
	})

	// cmd.RunGRPCServer()
}

func factorial(x int, next func(int)) {
	if x == 0 {
		next(1)
	} else {
		factorial(x-1, func(y int) {
			fmt.Println(y)
			next(x * y)
		})
	}
}
