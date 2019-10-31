package service

import (
	"github.com/daishinmutaku/quest_board_server/server/domain/model"
	"testing"
)

type TagRepositoryMock struct{}

func (mock *TagRepositoryMock) GetTag(id int) (*model.Tag, error) {
	output := &model.Tag{ID: id, Name: "aaa"}
	return output, nil
}

func (mock *TagRepositoryMock) CreateTag(string) (*model.Tag, error) {
	return nil, nil
}

func (mock *TagRepositoryMock) UpdateTag(id int, name string) (*model.Tag, error) {
	output := &model.Tag{ID: id, Name: name}
	return output, nil
}

func (mock *TagRepositoryMock) DeleteTag(id int) (*model.Tag, error) {
	output := &model.Tag{ID: id, Name: "aaa"}
	return output, nil
}

func TestGetTagByInvalidID(t *testing.T) {
	repository := &UserRepositoryMock{}
	service := NewUserService(repository)

	input := -1
	var expect *model.User
	expect = nil

	output, err := service.GetUser(input)
	if err != nil {
		t.Error("GetTag error")
	}

	t.Log(output)

	if output != expect {
		t.Errorf("expect isn't output")
	}
}

func TestCreateTag(t *testing.T) {
	repository := &UserRepositoryMock{}
	service := NewUserService(repository)

	input := "new tag"
	var expect *model.User
	expect = nil

	output, err := service.CreateUser(input)
	if err != nil {
		t.Error("GetTag error")
	}

	t.Log(output)

	if output != expect {
		t.Errorf("expect isn't output")
	}
}

func TestUpdateTag(t *testing.T) {
	repository := &UserRepositoryMock{}
	service := NewUserService(repository)

	inputId := 1
	inputName := "update Tag"
	var expect *model.User
	expect = nil

	output, err := service.UpdateUser(inputId, inputName)
	if err != nil {
		t.Error("Update tag error")
	}

	t.Log(output)

	if output != expect {
		t.Errorf("expect isn't output")
	}
}

func TestDeleteTag(t *testing.T) {
	repository := &UserRepositoryMock{}
	service := NewUserService(repository)

	input := -1
	var expect error
	expect = nil

	err := service.DeleteUser(input)
	if err != expect {
		t.Error("DeleteTag error")
	}

}
