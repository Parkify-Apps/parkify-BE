package data

import "gorm.io/gorm"

type Parking struct {
	gorm.Model
	ImageLoc     string
	LocationName string
	Location     string
	User_ID      uint
}
