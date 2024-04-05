package data

import (
	parking "parkify-BE/features/parking/data"
	parkingslot "parkify-BE/features/parkingslot/data"
	reservation "parkify-BE/features/reservation/data"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname    string
	Email       string `gorm:"type:varchar(30);unique"`
	Role        string
	Password    string
	Parking     parking.Parking         `gorm:"foreignKey:UserID;references:ID"`
	ParkingSlot parkingslot.ParkingSlot `gorm:"foreignKey:Email;references:Email"`
	Reservation reservation.Reservation `gorm:"foreignKey:Email;references:Email"`
}
