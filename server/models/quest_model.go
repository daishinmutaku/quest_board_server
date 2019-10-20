package models

import (
	"github.com/daishinmutaku/quest_board_server/server/entities"
	"github.com/daishinmutaku/quest_board_server/server/infra"
	"github.com/jinzhu/gorm"
	"time"
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

// Quest生成
func (model *QuestModel) CreateQuest(name, memberDescription, questDescription, reward string, capacity int64, period time.Time, isFinished bool, producer entities.User, member []entities.User, tag entities.Tag) entities.Quest {
	db := infra.NewSqlHandler()
	quest := entities.Quest{Name: name, Capacity: capacity, MemberDescription: memberDescription, QuestDescription: questDescription, Period: period, Reward: reward, IsFinished: isFinished, ProducerID: producer.Id, Producer: producer, Member: member, TagID: tag.Id, Tag: tag, CreatedDate: time.Now(), UpdatedDate: time.Now()}
	db.Create(&quest)
	return quest
}
