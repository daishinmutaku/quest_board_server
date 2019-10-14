package controllers

import (
	"github.com/daishinmutaku/quest_board_server/server/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type MessageController struct {
	Db *gorm.DB
}

func (controller *MessageController) Index(c *gin.Context) {
	messages := controller.findMessage()
	c.JSON(200, gin.H{
		"Messages": messages,
	})
}

// User全取得
func (controller *MessageController) findMessage() []models.Message {
	var messages []models.Message
	controller.Db.Find(&messages)
	defer controller.Db.Close()
	return messages
}
