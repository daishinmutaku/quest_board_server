package infra

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

func NewSqlHandler() *gorm.DB {
	dbURL := os.Getenv("MYSQL_URL")
	db, err := gorm.Open("mysql", dbURL+"?parseTime=true")
	if err != nil {
		log.Fatal("データベース開けず！（dbInsert)")
	} else {
		fmt.Println("DB Connect Success!")
	}
	db.SingularTable(true)

	return db
}
