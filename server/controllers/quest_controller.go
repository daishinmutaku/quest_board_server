package controllers

import (
	"github.com/daishinmutaku/quest_board_server/server/models"
	"github.com/gin-gonic/gin"
	"time"
)

type QuestController struct{}

func (controller *QuestController) Index(c *gin.Context) {
	model := models.QuestModel{}
	quests := model.FindQuest()
	c.JSON(200, gin.H{
		"Quests": quests,
	})
}

func (controller *QuestController) Create(c *gin.Context) {
	model := models.QuestModel{}
	userModel := models.UserModel{}
	tagModel := models.TagModel{}

	name := c.PostForm("name")
	memberDescription := c.PostForm("memberDescription")
	questDescription := c.PostForm("questDescription")
	reward := c.PostForm("reward")
	capacity := int64(2)
	period := time.Now()
	isFinished := false
	producer := userModel.FirstUser()
	member := userModel.FindUser()
	tag := tagModel.FirstTag()

	quests := model.CreateQuest(name, memberDescription, questDescription, reward, capacity, period, isFinished, producer, member, tag)
	c.JSON(200, gin.H{
		"Quests": quests,
	})
}
