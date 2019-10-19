package controllers

import (
	"github.com/daishinmutaku/quest_board_server/server/models"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (controller *UserController) Index(c *gin.Context) {
	model := models.UserModel{}
	users := model.FindUser()
	c.JSON(200, gin.H{
		"Users": users,
	})
}

func (controller *UserController) Create(c *gin.Context) {
	model := models.UserModel{}
	user := model.CreateUser()
	c.JSON(200, gin.H{
		"User": user,
	})
}
