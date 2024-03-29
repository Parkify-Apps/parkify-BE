package parking

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ParkingController interface {
	PostParking() echo.HandlerFunc
}

type ParkingServices interface {
	PostParking(token *jwt.Token, newData Parking) error
}

type ParkingModel interface {
	PostParking(newData Parking, email string) error
}

type Parking struct {
	ImageLoc     string
	LocationName string
	Location     string
	User_ID      uint
}

type AddParkingVal struct {
	ImageLoc     string `validate:"required" form:"imageloc"`
	LocationName string `validate:"required" form:"locname"`
	Location     string `validate:"required" form:"location"`
	User_ID      uint
}
