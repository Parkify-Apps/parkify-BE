package data

import (
	"parkify-BE/features/parking/data"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string
	Email    string `gorm:"type:varchar(30);"`
	Role     string
	Password string
	Parking  data.Parking `gorm:"foreignKey:UserID;references:ID"`
}
