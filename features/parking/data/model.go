package data

import (
	// "parkify-BE/features/parkingSlot/data"

	"gorm.io/gorm"
)

type Parking struct {
	gorm.Model
	ImageLoc    string
	Location    string
	City        string
	UserID      uint
	// ParkingSlot []data.ParkingSlot `gorm:"foreignKey:ParkingID;references:ID"`
}
