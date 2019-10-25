package controllers

import (
	"github.com/daishinmutaku/quest_board_server/server/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

type QuestController struct {
	Db *gorm.DB
}

func (controller *QuestController) Index(c *gin.Context) {
	questModel := models.QuestModel{controller.Db}
	quests := questModel.FindQuest()

	response := models.IndexQuestResponseModel{quests}
	c.JSON(200, response)
}

func (controller *QuestController) Create(c *gin.Context) {
	questModel := models.QuestModel{controller.Db}
	userModel := models.UserModel{controller.Db}
	tagModel := models.TagModel{controller.Db}

	requestModel := models.CreateQuestRequestModel{
		Name:              c.PostForm("name"),
		MemberDescription: c.PostForm("memberDescription"),
		QuestDescription:  c.PostForm("questDescription"),
		Reward:            c.PostForm("reward"),
		Capacity:          int64(2),
		Period:            time.Now(),
		IsFinished:        false,
		Producer:          userModel.FirstUser(),
		Member:            userModel.FindUser(),
		Tag:               tagModel.FirstTag(),
	}

	quest := questModel.CreateQuest(requestModel)

	response := models.CreateQuestResponseModel{quest}
	c.JSON(200, response)
}

func (controller *QuestController) Update(c *gin.Context) {
	questModel := models.QuestModel{controller.Db}
	userModel := models.UserModel{controller.Db}
	tagModel := models.TagModel{controller.Db}

	requestModel := models.UpdateQuestRequestModel{
		Id:                c.Param("id"),
		Name:              c.PostForm("name"),
		MemberDescription: c.PostForm("memberDescription"),
		QuestDescription:  c.PostForm("questDescription"),
		Reward:            c.PostForm("reward"),
		Capacity:          int64(2),
		Period:            time.Now(),
		IsFinished:        false,
		Producer:          userModel.FirstUser(),
		Member:            userModel.FindUser(),
		Tag:               tagModel.FirstTag(),
	}

	quest := questModel.UpdateQuest(requestModel)

	response := models.UpdateQuestResponseModel{quest}
	c.JSON(200, response)
}
