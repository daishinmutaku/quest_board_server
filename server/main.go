package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", Hello)
	r.POST("/create/user", CreateUser)
	r.GET("/index/user", GetUsers)
	r.Run(":8080")
}

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

func GetUsers(c *gin.Context) {
	users := getUsers()
	c.JSON(200, gin.H{
		"Users": users,
	})
}

func CreateUser(c *gin.Context) {
	user := createUser()
	c.JSON(200, gin.H{
		"User": user,
	})
}

func connectGorm() *gorm.DB {
	db, err := gorm.Open("mysql", "quest-board:quest-board@/quest-board")
	if err != nil {
		panic("データベース開けず！（dbInsert)")
	}
	return db
}

// User作成
func createUser() User {
	db := connectGorm()
	user := User{Name: "tester"}
	db.Create(&user)
	defer db.Close()
	return user
}

// User全取得
func getUsers() []User {
	db := connectGorm()
	var users []User
	db.Find(&users)
	defer db.Close()
	return users
}

// Model
type User struct {
	Id   int64
	Name string
}

type Tag struct {
	Id   int64
	Name string
}

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

type Application struct {
	Id        int64
	Status    int64
	Quest     Quest
	Applicant User
}

type Message struct {
	Id       int64
	Body     string
	SendTime time.Time
	Quest    Quest
	Sender   User
}
