package controllers

import (
	"github.com/daishinmutaku/quest_board_server/server/models"
	"github.com/gin-gonic/gin"
)

type QuestController struct{}

func (controller *QuestController) Index(c *gin.Context) {
	model := models.QuestModel{}
	quests := model.FindQuest()
	c.JSON(200, gin.H{
		"Quests": quests,
	})
}
