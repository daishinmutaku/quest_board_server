package service

import (
	"strings"

	"github.com/daishinmutaku/quest_board_server/server/domain/model"
	"github.com/daishinmutaku/quest_board_server/server/usecase/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{UserRepository: repository}
}

func (service *UserService) GetUser(id int) (*model.User, error) {
	if id <= 0 {
		return nil, nil
	}
	user, err := service.UserRepository.GetUser(id)
	if err != nil {
		return nil, nil
	}

	return user, nil
}

func (service *UserService) CreateUser(name string) (*model.User, error) {
	if strings.Index(name, " ") > -1 {
		return nil, nil
	}
	user, err := service.UserRepository.CreateUser(name)
	if err != nil {
		return nil, nil
	}

	return user, nil
}
