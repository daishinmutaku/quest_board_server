package controllers

import (
	"github.com/daishinmutaku/quest_board_server/server/models"

	"github.com/gin-gonic/gin"
)

type MessageController struct{}

func (controller *MessageController) Index(c *gin.Context) {
	model := models.MessageModel{}
	messages := model.FindMessage()
	c.JSON(200, gin.H{
		"Messages": messages,
	})
}
