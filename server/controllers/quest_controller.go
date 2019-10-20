package controllers

import (
	"github.com/daishinmutaku/quest_board_server/server/models"
	"github.com/daishinmutaku/quest_board_server/server/models/request"
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

	request := request.CreateQuestModel{}

	request.Name = c.PostForm("name")
	request.MemberDescription = c.PostForm("memberDescription")
	request.QuestDescription = c.PostForm("questDescription")
	request.Reward = c.PostForm("reward")
	request.Capacity = int64(2)
	request.Period = time.Now()
	request.IsFinished = false
	request.Producer = userModel.FirstUser()
	request.Member = userModel.FindUser()
	request.Tag = tagModel.FirstTag()

	quests := questModel.CreateQuest(request)
	c.JSON(200, gin.H{
		"Quests": quests,
	})
}
