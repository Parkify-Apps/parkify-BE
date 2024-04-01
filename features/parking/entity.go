package parking

import (
	"parkify-BE/features/user"

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
	GetDataByEmail(email string) (user.User, error)
}

type Parking struct {
	ImageLoc string
	Location string
	City     string
	User_ID  uint
}

type AddParkingVal struct {
	ImageLoc string `validate:"required" form:"imageloc"`
	Location string `validate:"required" form:"location"`
	City     string `validate:"required" form:"city"`
	User_ID  uint
}
