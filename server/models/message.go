package models

import "time"

type Message struct {
	Id       int64
	Body     string
	SendTime time.Time
	Quest    Quest
	Sender   User
}
