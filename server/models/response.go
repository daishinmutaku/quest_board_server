package models

import (
	"github.com/daishinmutaku/quest_board_server/server/entities"
	"github.com/gin-gonic/gin"
)

// Application
type IndexApplicationResponseModel struct {
	Applications []entities.Application `json:"Applications"`
}

//func (response IndexApplicationResponseModel) FormatToJson() interface{} {
//	return gin.H{
//		"Applications": response.Applications,
//	}
//}

// Message
type IndexMessageResponseModel struct {
	Messages []entities.Message `json:"Messages"`
}

//func (response IndexMessageResponseModel) FormatToJson() interface{} {
//	return gin.H{
//		"Messages": response.Messages,
//	}
//}

// Quest
type IndexQuestResponseModel struct {
	Quests []entities.Quest `json:"Quests"`
}

func (response IndexQuestResponseModel) FormatToJson() interface{} {
	return gin.H{
		"Quests": response.Quests,
	}
}

type CreateQuestResponseModel struct {
	Quest entities.Quest `json:"Quest"`
}

//func (response CreateQuestResponseModel) FormatToJson() interface{} {
//	return gin.H{
//		"Quest": response.Quest,
//	}
//}

// User
type IndexUserResponseModel struct {
	Users []entities.User `json:"Users"`
}

//func (response IndexUserResponseModel) FormatToJson() interface{} {
//	return gin.H{
//		"Users": response.Users,
//	}
//}

type CreateUserResponseModel struct {
	User entities.User `json:"User"`
}

//func (response CreateUserResponseModel) FormatToJson() interface{} {
//	return gin.H{
//		"User": response.User,
//	}
//}
