package models

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
	Producer          User
	Member            []User
	Tag               Tag
	CreatedDate       time.Time
	UpdatedDate       time.Time
}
