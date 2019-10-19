package models

import (
	"github.com/daishinmutaku/quest_board_server/server/entities"
	"github.com/daishinmutaku/quest_board_server/server/infra"
)

type MessageModel struct{}

// User全取得
func (model *MessageModel) FindMessage() []entities.Message {
	db := infra.NewSqlHandler()
	var messages []entities.Message
	db.Find(&messages)
	return messages
}
