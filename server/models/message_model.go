package models

import (
	"github.com/daishinmutaku/quest_board_server/server/entities"
	"github.com/jinzhu/gorm"
)

type MessageModel struct {
	Db *gorm.DB
}

// User全取得
func (model *MessageModel) FindMessage() []entities.Message {
	var messages []entities.Message
	model.Db.Find(&messages)
	return messages
}
