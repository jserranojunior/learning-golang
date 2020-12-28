package main

import (
	"./db"
	"./users"
)

func main() {
	db.Migrate()
	//users.Insert()
	// users.Select()
	//users.Update()
	//users.Delete()
	users.DeletePermanently()
}
