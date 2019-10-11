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

	r.GET("/", Hello)

	r.POST("/user/create", CreateUser)
	r.GET("/user/index", GetUsers)

	r.GET("/quest/create")

	r.Run(":8080")
}

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Daimu!",
	})
}

func GetUsers(c *gin.Context) {
	userController := controllers.UserConroller{}
	users := userController.GetUsers()
	c.JSON(200, gin.H{
		"Users": users,
	})
}

func CreateUser(c *gin.Context) {
	userController := controllers.UserConroller{}
	user := userController.CreateUser()
	c.JSON(200, gin.H{
		"User": user,
	})
}
