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
func (model *QuestModel) CreateQuest(name, memberDescription, questDescription, reward string, capacity int64, period time.Time, isFinished bool, producer entities.User, member []entities.User, tag entities.Tag) entities.Quest {
	quest := entities.Quest{Name: name, Capacity: capacity, MemberDescription: memberDescription, QuestDescription: questDescription, Period: period, Reward: reward, IsFinished: isFinished, ProducerID: producer.Id, Producer: producer, Member: member, TagID: tag.Id, Tag: tag, CreatedDate: time.Now(), UpdatedDate: time.Now()}
	model.Db.Create(&quest)
	return quest
}
