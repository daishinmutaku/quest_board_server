package controller

import (
	"github.com/daishinmutaku/quest_board_server/server/usecase/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{UserService: service}
}

func (controller *UserController) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := controller.UserService.GetUser(id)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "fail getting user",
		})
	}

	c.JSON(200, user)
}

func (controller *UserController) CreateUser(c *gin.Context) {
	name := c.PostForm("name")

	user, err := controller.UserService.CreateUser(name)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "fail creating user",
		})
	}

	c.JSON(200, user)
}

func (controller *UserController) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	name := c.PostForm("name")

	user, err := controller.UserService.UpdateUser(id, name)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "fail updating user",
		})
	}

	c.JSON(200, user)
}
