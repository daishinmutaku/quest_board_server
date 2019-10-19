package controllers

import (
	"github.com/daishinmutaku/quest_board_server/server/models"
	"github.com/gin-gonic/gin"
)

type ApplicationController struct{}

func (controller *ApplicationController) Index(c *gin.Context) {
	model := models.ApplicationModel{}
	applications := model.FindApplication()
	c.JSON(200, gin.H{
		"Applications": applications,
	})
}
