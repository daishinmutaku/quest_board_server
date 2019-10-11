package controllers

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

type DbConroller struct{}

func (controller *DbConroller) ConnectGorm() *gorm.DB {
	dbURL := os.Getenv("MYSQL_URL")
	db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		log.Fatal("データベース開けず！（dbInsert)")
	} else {
		fmt.Println("DB Connect Success!")
	}
	return db
}
