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
	c.JSON(200, gin.H{
		"Quests": quests,
	})
}

func (controller *QuestController) Create(c *gin.Context) {
	questModel := models.QuestModel{controller.Db}
	userModel := models.UserModel{controller.Db}
	tagModel := models.TagModel{controller.Db}

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

	quests := questModel.CreateQuest(name, memberDescription, questDescription, reward, capacity, period, isFinished, producer, member, tag)
	c.JSON(200, gin.H{
		"Quests": quests,
	})
}
