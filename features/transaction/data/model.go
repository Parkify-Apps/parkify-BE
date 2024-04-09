package data

import (
	
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ReservationID uint
	PaymentMethod string
	Price         int
	Status        string
	VirtualAccount string
}
