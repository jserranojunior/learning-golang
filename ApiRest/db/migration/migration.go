package migration

import (
	db ".."
	"../models"
)

//AutoMigration create tables
func AutoMigration() {
	db := db.Connect()
	defer db.Close()
	db.AutoMigrate(models.News{})
}
