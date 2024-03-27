package data

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname string
	Email    string `gorm:"type:varchar(30);"`
	Role     bool
	Password string
}
