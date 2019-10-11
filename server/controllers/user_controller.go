package controllers

import (
	"github.com/daishinmutaku/quest_board_server/server/models"
)

type UserConroller struct{}

// User作成
func (controller *UserConroller) CreateUser() models.User {
	dbController := DbConroller{}
	db := dbController.ConnectGorm()
	user := models.User{Name: "tester"}
	db.Create(&user)
	defer db.Close()
	return user
}

// User全取得
func (controller *UserConroller) GetUsers() []models.User {
	dbController := DbConroller{}
	db := dbController.ConnectGorm()
	var users []models.User
	db.Find(&users)
	defer db.Close()
	return users
}
