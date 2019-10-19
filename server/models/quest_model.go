package models

import (
	"github.com/daishinmutaku/quest_board_server/server/entities"
	"github.com/daishinmutaku/quest_board_server/server/infra"
	"github.com/jinzhu/gorm"
)

type QuestModel struct {
	Db *gorm.DB
}

// Quest全取得
func (model *QuestModel) FindQuest() []entities.Quest {
	db := infra.NewSqlHandler()
	var quests []entities.Quest
	db.Find(&quests)
	return quests
}
