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

func (handler SQLHandler) Update(dist interface{}, src interface{}, statement string, args ...interface{}) error {
	getErr := handler.DB.Where(statement, args...).First(dist).Error
	updateErr := handler.DB.Model(dist).Update(src).Error

	if getErr != nil {
		log.Println(getErr.Error())
		return getErr
	}
	if updateErr != nil {
		log.Println(updateErr.Error())
		return updateErr
	}

	return nil
}

func (handler SQLHandler) Delete(dist interface{}, statement string, args ...interface{}) error {
	//getErr := handler.DB.Where(statement, args...).First(dist).Error
	deleteErr := handler.DB.Where(statement, args...).Delete(dist).Error

	//if getErr != nil {
	//	log.Println(getErr.Error())
	//	return getErr
	//}
	if deleteErr != nil {
		log.Println(deleteErr.Error())
		return deleteErr
	}

	return nil
}
