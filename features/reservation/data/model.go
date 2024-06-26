package data

import (
	"parkify-BE/features/parkingslot/data"
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	ExitedAt      time.Time
	Email         string
	ParkingSlotID uint
	ParkingSlot   data.ParkingSlot `gorm:"foreignKey:ParkingSlotID;references:ID"`
}

type ReservationResponse struct {
	ID            uint
	Email         string
	ParkingSlotID uint
	VehicleType   string
	Floor         int
	Slot          int
	Price         int
	ParkingID     uint
	ImageLoc      string
	Location      string
	City          string
}
