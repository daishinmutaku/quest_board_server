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

	CreatedDate time.Time
	UpdatedDate time.Time
	ProducerID  int64
	TagID       int64
	Tag         Tag
	Producer    User
	Member      []User

	IsFinished bool
}
