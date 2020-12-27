package db

import users "../users"

//Migrate exec migrate
func Migrate() {
	connection := Connection()
	connection.AutoMigrate(users.Users{})
}
