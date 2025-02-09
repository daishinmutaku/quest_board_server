package repository

import "github.com/daishinmutaku/quest_board_server/server/domain/model"

type UserRepository interface {
	GetUser(int) (*model.User, error)
	CreateUser(string) (*model.User, error)
}
