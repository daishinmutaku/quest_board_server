package models

import (
	"github.com/daishinmutaku/quest_board_server/server/entities"
)

// Application
type IndexApplicationResponseModel struct {
	Applications []entities.Application `json:"Applications"`
}

// Message
type IndexMessageResponseModel struct {
	Messages []entities.Message `json:"Messages"`
}

// Quest
type IndexQuestResponseModel struct {
	Quests []entities.Quest `json:"Quests"`
}

type CreateQuestResponseModel struct {
	Quest entities.Quest `json:"Quest"`
}

type UpdateQuestResponseModel struct {
	Quest entities.Quest `json:"Quest"`
}

// User
type IndexUserResponseModel struct {
	Users []entities.User `json:"Users"`
}

type CreateUserResponseModel struct {
	User entities.User `json:"User"`
}
