package models

import (
	"github.com/daishinmutaku/quest_board_server/server/entities"
	"github.com/daishinmutaku/quest_board_server/server/models/request"
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
func (model *QuestModel) CreateQuest(questModel request.CreateQuestModel) entities.Quest {
	quest := entities.Quest{
		Name:              questModel.Name,
		Capacity:          questModel.Capacity,
		MemberDescription: questModel.MemberDescription,
		QuestDescription:  questModel.QuestDescription,
		Period:            questModel.Period,
		Reward:            questModel.Reward,
		IsFinished:        questModel.IsFinished,
		ProducerID:        questModel.Producer.Id,
		Producer:          questModel.Producer,
		Member:            questModel.Member,
		TagID:             questModel.Tag.Id,
		Tag:               questModel.Tag,
		CreatedDate:       time.Now(),
		UpdatedDate:       time.Now(),
	}

	model.Db.Create(&quest)
	return quest
}
