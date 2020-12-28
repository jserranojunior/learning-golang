package db

import (
	"../config"
	users "../users"
)

//Migrate exec migrate
func Migrate() {
	connection := config.Connection()
	connection.AutoMigrate(users.Users{})
}
