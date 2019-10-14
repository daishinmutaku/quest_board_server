package controllers

import (
	"github.com/daishinmutaku/quest_board_server/server/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ApplicationController struct {
	Db *gorm.DB
}

func (controller *ApplicationController) Index(c *gin.Context) {
	applications := controller.findApplication()
	c.JSON(200, gin.H{
		"Applications": applications,
	})
}

// User全取得
func (controller *ApplicationController) findApplication() []models.Application {
	var applications []models.Application
	controller.Db.Find(&applications)
	defer controller.Db.Close()
	return applications
}
