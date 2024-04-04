package data

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	Exited_at     time.Time
	ParkingSlotID uint
	Price         int
}
