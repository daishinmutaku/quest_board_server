package infra

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

func NewSqlHandler() *gorm.DB {
	dbURL := os.Getenv("MYSQL_URL")
	db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		log.Fatal("データベース開けず！（dbInsert)")
	} else {
		fmt.Println("DB Connect Success!")
	}
	return db
}
