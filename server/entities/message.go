package entities

import "time"

type Message struct {
	Id       int64     `db:"id"`
	Body     string    `db:"body"`
	SendTime time.Time `db:"send_time"`
	Quest    Quest     `db:"quest"`
	Sender   User      `db:"sender"`
}
