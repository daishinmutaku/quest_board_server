package controllers

import (
	"github.com/daishinmutaku/quest_board_server/server/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type QuestController struct {
	Db *gorm.DB
}

func (controller *QuestController) Index(c *gin.Context) {
	quests := controller.findQuest()
	c.JSON(200, gin.H{
		"Quests": quests,
	})
}

// User全取得
func (controller *QuestController) findQuest() []models.Quest {
	var quests []models.Quest
	controller.Db.Find(&quests)
	defer controller.Db.Close()
	return quests
}
