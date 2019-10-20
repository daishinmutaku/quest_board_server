package models

import (
	"github.com/daishinmutaku/quest_board_server/server/entities"
	"github.com/daishinmutaku/quest_board_server/server/infra"
)

type UserModel struct{}

// User全取得
func (model *UserModel) FindUser() []entities.User {
	db := infra.NewSqlHandler()
	var users []entities.User
	db.Find(&users)
	return users
}

// 最初のUser取得
func (model *UserModel) FirstUser() entities.User {
	db := infra.NewSqlHandler()
	user := entities.User{}
	db.First(&user)

	return user
}

// User作成
func (model *UserModel) CreateUser(name string) entities.User {
	db := infra.NewSqlHandler()
	user := entities.User{Name: name}
	db.Create(&user)
	return user
}
