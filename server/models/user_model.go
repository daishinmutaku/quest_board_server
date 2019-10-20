package models

import (
	"github.com/daishinmutaku/quest_board_server/server/entities"
	"github.com/jinzhu/gorm"
)

type UserModel struct {
	Db *gorm.DB
}

// User全取得
func (model *UserModel) FindUser() []entities.User {
	var users []entities.User
	model.Db.Find(&users)
	return users
}

// 最初のUser取得
func (model *UserModel) FirstUser() entities.User {
	user := entities.User{}
	model.Db.First(&user)

	return user
}

// User作成
func (model *UserModel) CreateUser(requestModel CreateUserRequestModel) entities.User {
	name := requestModel.Name
	user := entities.User{Name: name}
	model.Db.Create(&user)
	return user
}
