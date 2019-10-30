package datastore

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SQLHandler struct {
	DB *gorm.DB
}

func NewSqlHandler() SQLHandler {
	dbURL := os.Getenv("MYSQL_URL")
	db, err := gorm.Open("mysql", dbURL+"?parseTime=true")
	if err != nil {
		log.Println(err)
		log.Fatal("データベース開けず！（dbInsert)")
	} else {
		fmt.Println("DB Connect Success!")
	}
	db.SingularTable(true)

	return SQLHandler{DB: db}
}

func (handler SQLHandler) Get(dist interface{}, statement string, args ...interface{}) error {
	err := handler.DB.Where(statement, args...).First(dist).Error

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (handler SQLHandler) Create(src interface{}) error {
	err := handler.DB.Create(src).Error

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
