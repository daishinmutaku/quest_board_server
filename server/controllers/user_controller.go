package controllers

import (
	"github.com/daishinmutaku/quest_board_server/server/models"
	"github.com/daishinmutaku/quest_board_server/server/models/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UserController struct {
	Db *gorm.DB
}

func (controller *UserController) Index(c *gin.Context) {
	userModel := models.UserModel{controller.Db}
	users := userModel.FindUser()

	response := response.Response{Key: "Users", Value: users}
	c.JSON(200, response.FormatToJson())
}

func (controller *UserController) Create(c *gin.Context) {
	userModel := models.UserModel{controller.Db}
	name := c.PostForm("name")
	user := userModel.CreateUser(name)

	response := response.Response{Key: "User", Value: user}
	c.JSON(200, response.FormatToJson())
}
