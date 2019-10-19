package models

import (
	"github.com/daishinmutaku/quest_board_server/server/entities"
	"github.com/daishinmutaku/quest_board_server/server/infra"
)

type ApplicationModel struct{}

// User全取得
func (model *ApplicationModel) FindApplication() []entities.Application {
	db := infra.NewSqlHandler()
	var applications []entities.Application
	db.Find(&applications)
	return applications
}
