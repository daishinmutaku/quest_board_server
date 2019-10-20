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
	messageModel := models.MessageModel{controller.Db}
	messages := messageModel.FindMessage()
	c.JSON(200, gin.H{
		"Messages": messages,
	})
}
