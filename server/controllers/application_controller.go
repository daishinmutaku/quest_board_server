package controllers

import (
	"github.com/daishinmutaku/quest_board_server/server/models"
	"github.com/daishinmutaku/quest_board_server/server/models/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ApplicationController struct {
	Db *gorm.DB
}

func (controller *ApplicationController) Index(c *gin.Context) {
	applicationModel := models.ApplicationModel{controller.Db}
	applications := applicationModel.FindApplication()

	response := response.Response{Key: "Applications", Value: applications}
	c.JSON(200, response.FormatToJson())
}
