package users

import (
	"fmt"

	"../config"
	"gorm.io/gorm"
)

var conn = config.Connection()

//Users export
type Users struct {
	gorm.Model
	Name  string `gorm:"size:255; not null" json:"name"`
	Email string `gorm:"size:255; not null; unique;" json:"email"`
}

//Insert return rows affected
func Insert() int64 {
	user := Users{Name: "Jorge", Email: "jorgeserranojunior"}

	result := conn.Create(&user)
	return result.RowsAffected
}

//Select return user
func Select() {
	var user Users
	conn.First(&user, 1)
	fmt.Println(user)
}

//Update of user
func Update() {
	var user Users
	conn.First(&user, 1)
	user.Name = "Juninho"
	user.Email = "jserranojunior@gmail.com"
	conn.Save(&user)
}

//Delete row selected
func Delete() {
	var user Users
	conn.First(&user, 1)
	conn.Delete(&user)
}

//DeletePermanently force delete
func DeletePermanently() {
	var user Users
	conn.First(&user, 1)
	conn.Unscoped().Delete(&user)
}
