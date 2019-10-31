package controller

import (
	"github.com/daishinmutaku/quest_board_server/server/usecase/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TagController struct {
	TagService *service.TagService
}

func NewTagController(service *service.TagService) *TagController {
	return &TagController{TagService: service}
}

func (controller *TagController) GetTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	tag, err := controller.TagService.GetTag(id)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "fail getting user",
		})
	}

	c.JSON(200, tag)
}

func (controller *TagController) CreateTag(c *gin.Context) {
	name := c.PostForm("name")

	tag, err := controller.TagService.CreateTag(name)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "fail creating user",
		})
	}

	c.JSON(200, tag)
}

func (controller *TagController) UpdateTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	name := c.PostForm("name")

	tag, err := controller.TagService.UpdateTag(id, name)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "fail updating user",
		})
	}

	c.JSON(200, tag)
}

func (controller *TagController) DeleteTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := controller.TagService.DeleteTag(id)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "fail getting user",
		})
	}

	c.JSON(200, err)
}
