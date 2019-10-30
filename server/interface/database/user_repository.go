package database

import (
	"github.com/daishinmutaku/quest_board_server/server/domain/model"
)

type UserRepository struct {
	SQLHandler SQLHandler
}

func NewUserRepository(sqlHandler SQLHandler) *UserRepository {
	return &UserRepository{SQLHandler: sqlHandler}
}

func (repository *UserRepository) GetUser(id int) (*model.User, error) {
	user := &model.User{}
	err := repository.SQLHandler.Get(user, "id = ?", id)

	if err != nil {
		return nil, nil
	}

	return user, nil
}

func (repository *UserRepository) CreateUser(name string) (*model.User, error) {
	user := &model.User{Name: name}
	err := repository.SQLHandler.Create(user)

	if err != nil {
		return nil, nil
	}

	return user, nil
}
