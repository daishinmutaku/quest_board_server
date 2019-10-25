package models

import (
	"github.com/daishinmutaku/quest_board_server/server/entities"
	"time"
)

// Application
type CreateApplicationRequestModel struct {
	Quest entities.Quest
	User  entities.User
}

// Quest
type CreateQuestRequestModel struct {
	Name, MemberDescription, QuestDescription, Reward string
	Capacity                                          int64
	Period                                            time.Time
	IsFinished                                        bool
	Producer                                          entities.User
	Member                                            []entities.User
	Tag                                               entities.Tag
}

type UpdateQuestRequestModel struct {
	Id                                                string
	Name, MemberDescription, QuestDescription, Reward string
	Capacity                                          int64
	Period                                            time.Time
	IsFinished                                        bool
	Producer                                          entities.User
	Member                                            []entities.User
	Tag                                               entities.Tag
}

// User
type CreateUserRequestModel struct {
	Name string
}
