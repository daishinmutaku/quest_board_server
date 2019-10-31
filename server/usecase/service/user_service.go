package service

import (
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
	if isIncorrectID(id) {
		return nil, nil
	}
	user, err := service.UserRepository.GetUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserService) CreateUser(name string) (*model.User, error) {
	if hasBlank(name) {
		return nil, nil
	}
	user, err := service.UserRepository.CreateUser(name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserService) UpdateUser(id int, name string) (*model.User, error) {
	if hasBlank(name) {
		return nil, nil
	}
	user, err := service.UserRepository.UpdateUser(id, name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserService) DeleteUser(id int) error {
	if isIncorrectID(id) {
		return nil
	}
	err := service.UserRepository.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
