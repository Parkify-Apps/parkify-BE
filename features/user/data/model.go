package data

import (
	"parkify-BE/features/parkingslot"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname    string
	Email       string `gorm:"type:varchar(30);unique"`
	Role        string
	Password    string
	ParkingSlot []parkingslot.ParkingSlot `gorm:"foreignKey:Email;references:Email"`
}
