package main

import (
	"github.com/daishinmutaku/quest_board_server/server/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	//controller
	dbController := controllers.DbController{}
	userController := controllers.UserController{dbController.NewSqlHandler()}

	r.GET("/", Hello)

	r.POST("/user/create", userController.Create)
	r.GET("/user/index", userController.Index)

	r.GET("/quest/create")

	r.Run(":8080")
}

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Daimu!",
	})
}
