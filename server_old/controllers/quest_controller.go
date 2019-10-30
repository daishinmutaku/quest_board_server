package controllers

import (
	"github.com/daishinmutaku/quest_board_server/server/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type QuestController struct {
	Db *gorm.DB
}

func (controller *QuestController) Index(c *gin.Context) {
	questModel := models.QuestModel{controller.Db}
	quests := questModel.FindQuest()

	response := models.IndexQuestResponseModel{quests}
	c.JSON(200, response)
}

func (controller *QuestController) Create(c *gin.Context) {
	questModel := models.QuestModel{controller.Db}
	userModel := models.UserModel{controller.Db}
	tagModel := models.TagModel{controller.Db}

	capacity, _ := strconv.ParseInt(c.PostForm("capacity"), 10, 64)
	producerID, _ := strconv.ParseInt(c.PostForm("producer_id"), 10, 64)
	tagID, _ := strconv.ParseInt(c.PostForm("tag_id"), 10, 64)
	isFinished, _ := strconv.ParseBool(c.PostForm("is_finished"))

	producer := userModel.FirstUserWhereID(producerID)
	tag := tagModel.FirstTagWhereId(tagID)

	requestModel := models.CreateQuestRequestModel{
		Name:              c.PostForm("name"),
		MemberDescription: c.PostForm("memberDescription"),
		QuestDescription:  c.PostForm("questDescription"),
		Reward:            c.PostForm("reward"),
		Capacity:          capacity,
		Period:            time.Now(),
		IsFinished:        isFinished,
		Producer:          producer,
		Tag:               tag,
	}

	quest := questModel.CreateQuest(requestModel)

	response := models.CreateQuestResponseModel{quest}
	c.JSON(200, response)
}

func (controller *QuestController) Update(c *gin.Context) {
	questModel := models.QuestModel{controller.Db}
	userModel := models.UserModel{controller.Db}
	tagModel := models.TagModel{controller.Db}

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	capacity, _ := strconv.ParseInt(c.PostForm("capacity"), 10, 64)
	producerID, _ := strconv.ParseInt(c.PostForm("producerID"), 10, 64)
	tagID, _ := strconv.ParseInt(c.PostForm("tagID"), 10, 64)
	isFinished, _ := strconv.ParseBool(c.PostForm("is_finished"))
	//memberIDs, _ := strconv.ParseInt(c.PostForm("memberIDs"), 10, 64)

	producer := userModel.FirstUserWhereID(producerID)
	tag := tagModel.FirstTagWhereId(tagID)

	requestModel := models.UpdateQuestRequestModel{
		ID:                id,
		Name:              c.PostForm("name"),
		MemberDescription: c.PostForm("memberDescription"),
		QuestDescription:  c.PostForm("questDescription"),
		Reward:            c.PostForm("reward"),
		Capacity:          capacity,
		Period:            time.Now(),
		IsFinished:        isFinished,
		Producer:          producer,
		Tag:               tag,
	}

	quest := questModel.UpdateQuest(requestModel)

	response := models.UpdateQuestResponseModel{quest}
	c.JSON(200, response)
}
