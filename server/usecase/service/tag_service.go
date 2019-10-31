package service

import (
	"github.com/daishinmutaku/quest_board_server/server/domain/model"
	"github.com/daishinmutaku/quest_board_server/server/usecase/repository"
)

type TagService struct {
	TagRepository repository.TagRepository
}

func NewTagService(repository repository.TagRepository) *TagService {
	return &TagService{TagRepository: repository}
}

func (service *TagService) GetTag(id int) (*model.Tag, error) {
	if isIncorrectID(id) {
		return nil, nil
	}
	tag, err := service.TagRepository.GetTag(id)
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (service *TagService) CreateTag(name string) (*model.Tag, error) {
	if hasBlank(name) {
		return nil, nil
	}
	tag, err := service.TagRepository.CreateTag(name)
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (service *TagService) UpdateTag(id int, name string) (*model.Tag, error) {
	if hasBlank(name) {
		return nil, nil
	}
	tag, err := service.TagRepository.UpdateTag(id, name)
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (service *TagService) DeleteTag(id int) error {
	if isIncorrectID(id) {
		return nil
	}
	err := service.TagRepository.DeleteTag(id)
	if err != nil {
		return err
	}

	return nil
}
