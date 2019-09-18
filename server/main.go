package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", Hello)
	r.Run(":8080")
}

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
