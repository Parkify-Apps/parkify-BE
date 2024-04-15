package data

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	OrderID        string
	ReservationID  uint
	PaymentMethod  string
	Price          int
	Status         string
	VirtualAccount string
}
