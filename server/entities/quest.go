package entities

import (
	"time"
)

type Quest struct {
	Id                int64  `db:"id"`
	Name              string `db:"name"`
	Capacity          int64  `db:"capacity"`
	MemberDescription string `db:"memberDescription"`
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
