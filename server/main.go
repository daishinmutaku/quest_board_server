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
	questController := controllers.QuestController{}
	userController := controllers.UserController{}
	applicationController := controllers.ApplicationController{}
	messageController := controllers.MessageController{}

	//quest
	r.GET("/quest/index", questController.Index)
	r.POST("/quest/create", questController.Create)

	//user
	r.GET("/user/index", userController.Index)
	r.POST("/user/create", userController.Create)

	//application
	r.GET("/application/index", applicationController.Index)

	//message
	r.GET("/message/index", messageController.Index)

	r.Run(":8080")
}
