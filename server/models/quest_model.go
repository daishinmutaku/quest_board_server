package models

import (
	"github.com/daishinmutaku/quest_board_server/server/entities"
	"github.com/jinzhu/gorm"
	"time"
)

type QuestModel struct {
	Db *gorm.DB
}

// Quest全取得
func (model *QuestModel) FindQuest() []entities.Quest {
	var quests []entities.Quest
	model.Db.Find(&quests)
	return quests
}

// Quest生成
func (model *QuestModel) CreateQuest(requestModel CreateQuestRequestModel) entities.Quest {
	quest := entities.Quest{
		Name:              requestModel.Name,
		Capacity:          requestModel.Capacity,
		MemberDescription: requestModel.MemberDescription,
		QuestDescription:  requestModel.QuestDescription,
		Period:            requestModel.Period,
		Reward:            requestModel.Reward,
		IsFinished:        requestModel.IsFinished,
		ProducerID:        requestModel.Producer.Id,
		Producer:          requestModel.Producer,
		Member:            requestModel.Member,
		TagID:             requestModel.Tag.Id,
		Tag:               requestModel.Tag,
		CreatedDate:       time.Now(),
		UpdatedDate:       time.Now(),
	}

	model.Db.Create(&quest)
	return quest
}
