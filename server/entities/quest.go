package entities

import (
	"time"
)

type Quest struct {
	Id                int64
	Name              string
	Capacity          int64
	MemberDescription string
	QuestDescription  string
	Period            time.Time
	Reward            string
	IsFinished        bool

	ProducerID int64
	Producer   User

	Member []User

	TagID int64
	Tag   Tag

	CreatedDate time.Time
	UpdatedDate time.Time
}
