package controllers

import (
	"github.com/daishinmutaku/quest_board_server/server/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UserController struct {
	Db *gorm.DB
}

func (controller *UserController) Index(c *gin.Context) {
	users := controller.findUser()
	c.JSON(200, gin.H{
		"Users": users,
	})
}

func (controller *UserController) Create(c *gin.Context) {
	user := controller.createUser()
	c.JSON(200, gin.H{
		"User": user,
	})
}

// User作成
func (controller *UserController) createUser() models.User {
	user := models.User{Name: "tester"}
	controller.Db.Create(&user)
	defer controller.Db.Close()
	return user
}

// User全取得
func (controller *UserController) findUser() []models.User {
	var users []models.User
	controller.Db.Find(&users)
	defer controller.Db.Close()
	return users
}
