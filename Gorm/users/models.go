package users

import "gorm.io/gorm"

//Users export
type Users struct {
	gorm.Model
	Name  string `gorm:"size:255; not null" json:"name"`
	Email string `gorm:"size:255; not null; unique;" json:"email"`
}
