package database

import (
	"github.com/daishinmutaku/quest_board_server/server/domain/model"
)

type TagRepository struct {
	SQLHandler SQLHandler
}

func NewTagRepository(sqlHandler SQLHandler) *TagRepository {
	return &TagRepository{SQLHandler: sqlHandler}
}

func (repository *TagRepository) GetTag(id int) (*model.Tag, error) {
	tag := &model.Tag{}
	err := repository.SQLHandler.Get(tag, "id = ?", id)

	if err != nil {
		return nil, nil
	}

	return tag, nil
}

func (repository *TagRepository) CreateTag(name string) (*model.Tag, error) {
	tag := &model.Tag{Name: name}
	err := repository.SQLHandler.Create(tag)

	if err != nil {
		return nil, nil
	}

	return tag, nil
}

func (repository *TagRepository) UpdateTag(id int, name string) (*model.Tag, error) {
	tag := &model.Tag{}
	updatedTag := &model.Tag{Name: name}
	err := repository.SQLHandler.Update(tag, updatedTag, "id = ?", id)

	if err != nil {
		return nil, nil
	}

	return tag, nil
}

func (repository *TagRepository) DeleteTag(id int) error {
	tag := &model.Tag{}
	err := repository.SQLHandler.Delete(tag, "id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
