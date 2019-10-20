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

	response := models.IndexUserResponseModel{Users: users}
	c.JSON(200, response)
}

func (controller *UserController) Create(c *gin.Context) {
	userModel := models.UserModel{controller.Db}
	requestModel := models.CreateUserRequestModel{
		Name: c.PostForm("name"),
	}
	user := userModel.CreateUser(requestModel)

	response := models.CreateUserResponseModel{User: user}
	c.JSON(200, response)
}
