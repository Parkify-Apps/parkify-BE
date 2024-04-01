package data

import "gorm.io/gorm"

type Parking struct {
	gorm.Model
	ImageLoc string
	Location string
	City     string
	UserID   uint
}
