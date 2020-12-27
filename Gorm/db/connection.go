package db

import (
	"fmt"

	mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbName = "golangdb"
	dbHost = "127.0.0.1"
	dbUser = "root"
	dbPass = "golangpass"
	dbPort = "3307"
)

//Connection export connection mysql
func Connection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
