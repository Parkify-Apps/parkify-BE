package data

import (
	"parkify-BE/features/parkingslot/data"

	"gorm.io/gorm"
)

type Parking struct {
	gorm.Model
	ImageLoc     string
	Location     string
	City         string
	UserID       uint
	ParkingSlots []data.ParkingSlot `gorm:"foreignKey:ParkingID;references:ID"`
}
