package models

import (
	"github.com/daishinmutaku/quest_board_server/server/entities"
	"github.com/jinzhu/gorm"
)

type ApplicationModel struct {
	Db *gorm.DB
}

// User全取得
func (model *ApplicationModel) FindApplication() []entities.Application {
	var applications []entities.Application
	model.Db.Find(&applications)
	return applications
}
