package service

import (
	"github.com/daishinmutaku/quest_board_server/server/domain/model"
	"testing"
)

type UserRepositoryMock struct{}

func (mock *UserRepositoryMock) GetUser(id int) (*model.User, error) {
	output := &model.User{ID: id, Name: "aaa"}
	return output, nil
}

func (mock *UserRepositoryMock) CreateUser(string) (*model.User, error) {
	//output := &model.User{ID: 1, Name: "aaa"}
	//return output, nil
	return nil, nil
}

func TestGetUserByInvalidID(t *testing.T) {
	repository := &UserRepositoryMock{}
	service := NewUserService(repository)

	input := -1
	var expect *model.User
	expect = nil

	output, err := service.GetUser(input)
	if err != nil {
		t.Error("GetUser error")
	}

	t.Log(output)

	if output != expect {
		t.Errorf("expect isn't output")
	}
}

func TestCreateUser(t *testing.T) {
	repository := &UserRepositoryMock{}
	service := NewUserService(repository)

	input := "new user"
	var expect *model.User
	expect = nil

	output, err := service.CreateUser(input)
	if err != nil {
		t.Error("GetUser error")
	}

	t.Log(output)

	if output != expect {
		t.Errorf("expect isn't output")
	}
}
