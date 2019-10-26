package controllers

import (
	"github.com/daishinmutaku/quest_board_server/server/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

type ApplicationController struct {
	Db *gorm.DB
}

func (controller *ApplicationController) Index(c *gin.Context) {
	applicationModel := models.ApplicationModel{controller.Db}
	applications := applicationModel.FindApplication()

	response := models.IndexApplicationResponseModel{Applications: applications}
	c.JSON(200, response)
}

func (controller *ApplicationController) Create(c *gin.Context) {
	applicationModel := models.ApplicationModel{controller.Db}
	questModel := models.QuestModel{controller.Db}
	userModel := models.UserModel{controller.Db}

	questId, _ := strconv.ParseInt(c.PostForm("questId"), 10, 64)
	userId, _ := strconv.ParseInt(c.PostForm("userId"), 10, 64)

	quest := questModel.FirstQuestWhereID(questId)
	user := userModel.FirstUserWhereID(userId)

	requestModel := models.CreateApplicationRequestModel{
		Quest: quest,
		User:  user,
	}

	application := applicationModel.CreateApplication(requestModel)

	response := models.CreateApplicationtResponseModel{Application: application}
	c.JSON(200, response)
}
