package data

import (
	"parkify-BE/features/parking/data"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string
	Email    string `gorm:"type:varchar(30);"`
	Role     bool
	Password string
	Parkings []data.Parking `gorm:"foreignKey:User_ID;references:ID"`
}
