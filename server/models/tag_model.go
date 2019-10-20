package models

import (
	"github.com/daishinmutaku/quest_board_server/server/entities"
	"github.com/jinzhu/gorm"
)

type TagModel struct {
	Db *gorm.DB
}

// 最初のTag取得
func (model *TagModel) FirstTag() entities.Tag {
	tag := entities.Tag{}
	model.Db.First(&tag)

	return tag
}
