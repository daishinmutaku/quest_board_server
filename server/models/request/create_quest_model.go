package request

import (
	"github.com/daishinmutaku/quest_board_server/server/entities"
	"time"
)

type CreateQuestModel struct {
	Name, MemberDescription, QuestDescription, Reward string
	Capacity                                          int64
	Period                                            time.Time
	IsFinished                                        bool
	Producer                                          entities.User
	Member                                            []entities.User
	Tag                                               entities.Tag
}
