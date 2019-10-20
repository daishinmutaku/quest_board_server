package models

import (
	"github.com/daishinmutaku/quest_board_server/server/entities"
	"github.com/daishinmutaku/quest_board_server/server/infra"
)

type TagModel struct{}

// 最初のTag取得
func (model *TagModel) FirstTag() entities.Tag {
	db := infra.NewSqlHandler()
	tag := entities.Tag{}
	db.First(&tag)

	return tag
}
