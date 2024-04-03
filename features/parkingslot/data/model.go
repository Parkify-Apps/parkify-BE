package data

import "gorm.io/gorm"

type ParkingSlot struct {
	gorm.Model
	Email       string
	ParkingID   uint
	VehicleType string
	Floor       int
	Slot        int
	Price       int
	Status      string
}
