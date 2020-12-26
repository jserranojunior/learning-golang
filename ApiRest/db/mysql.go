package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	//mysql connect
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	dbName = "golangdb"
	dbHost = "localhost"
	dbUser = "root"
	dbPass = "golangpass"
	dbPort = "3306"
)

//Connect return connection
func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		fmt.Println("Error connect database", err)
	}
	return db
}
