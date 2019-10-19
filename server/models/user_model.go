package models

import (
	"github.com/daishinmutaku/quest_board_server/server/entities"
	"github.com/daishinmutaku/quest_board_server/server/infra"
)

type UserModel struct{}

// User作成
func (model *UserModel) CreateUser() entities.User {
	db := infra.NewSqlHandler()
	user := entities.User{Name: "tester"}
	db.Create(&user)
	return user
}

// User全取得
func (model *UserModel) FindUser() []entities.User {
	db := infra.NewSqlHandler()
	var users []entities.User
	db.Find(&users)
	return users
}
