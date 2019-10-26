package models

import (
	"github.com/daishinmutaku/quest_board_server/server/entities"
	"github.com/jinzhu/gorm"
)

type ApplicationModel struct {
	Db *gorm.DB
}

// Applicatioon全取得
func (model *ApplicationModel) FindApplication() []entities.Application {
	var applications []entities.Application
	model.Db.Find(&applications)
	return applications
}

// Applicatioon生成
func (model *ApplicationModel) CreateApplication(requestModel CreateApplicationRequestModel) entities.Application {
	application := entities.Application{
		Status:      false,
		Quest:       requestModel.Quest,
		QuestId:     requestModel.Quest.Id,
		Applicant:   requestModel.User,
		ApplicantId: requestModel.User.Id,
	}

	model.Db.Create(&application)
	return application
}
