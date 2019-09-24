package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", Hello)
	r.Run(":8080")
}

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

// Model
type User struct {
    Id int64
    Name string `sql:"size:255"`
    CreatedAt time.Time
    UpdatedAT time.Time
    DeletedAt time.Time
}

type Tag struct {
    Id int64
    Name string `sql:"size:255"`
    CreatedAt time.Time
    UpdatedAT time.Time
    DeletedAt time.Time
}

type Quest struct {
    Id int64
	Name string 
	Capacity int64
	MemberDescription string 
	QuestDescription string 
	Period time.Time
	Reward string 
	IsFinished bool
	Producer User
	Member []User
	Tag Tag
    CreatedAt time.Time
    UpdatedAT time.Time
    DeletedAt time.Time
}

type Application struct {
	Id int64
	Status int64
	Quest Quest
	Applicant User
    CreatedAt time.Time
    UpdatedAT time.Time
    DeletedAt time.Time
}

type Message struct {
    Id int64
	Body string
	SendTime time.Time
	Quest Quest
	Sender User
    CreatedAt time.Time
    UpdatedAT time.Time
    DeletedAt time.Time
}
