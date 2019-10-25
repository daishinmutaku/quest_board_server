package main

import (
	"github.com/daishinmutaku/quest_board_server/server/controllers"
	"github.com/daishinmutaku/quest_board_server/server/infra"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	db := infra.NewSqlHandler()

	//controller
	questController := controllers.QuestController{db}
	userController := controllers.UserController{db}
	applicationController := controllers.ApplicationController{db}
	messageController := controllers.MessageController{db}

	//quest
	r.GET("/quest/index", questController.Index)
	r.POST("/quest/create", questController.Create)
	r.PUT("/quest/update/:id", questController.Update)

	//user
	r.GET("/user/index", userController.Index)
	r.POST("/user/create", userController.Create)

	//application
	r.GET("/application/index", applicationController.Index)

	//message
	r.GET("/message/index", messageController.Index)

	r.Run(":8080")
}
