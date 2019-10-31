package repository

import "github.com/daishinmutaku/quest_board_server/server/domain/model"

type TagRepository interface {
	GetTag(int) (*model.Tag, error)
	CreateTag(string) (*model.Tag, error)
	UpdateTag(int, string) (*model.Tag, error)
	DeleteTag(int) (*model.Tag, error)
}
