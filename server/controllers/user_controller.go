package controllers

import (
	"github.com/daishinmutaku/quest_board_server/server/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UserController struct {
	Db *gorm.DB
}

func (controller *UserController) Index(c *gin.Context) {
	userModel := models.UserModel{controller.Db}
	users := userModel.FindUser()
	c.JSON(200, gin.H{
		"Users": users,
	})
}

func (controller *UserController) Create(c *gin.Context) {
	userModel := models.UserModel{controller.Db}
	name := c.PostForm("name")
	user := userModel.CreateUser(name)
	c.JSON(200, gin.H{
		"User": user,
	})
}
