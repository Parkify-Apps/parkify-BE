package parking

import (
	// "parkify-BE/features/parkingslot"
	"parkify-BE/features/parkingslot"
	"parkify-BE/features/user"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ParkingController interface {
	PostParking() echo.HandlerFunc
	UpdateParking() echo.HandlerFunc
	GetParking() echo.HandlerFunc
	GetAllParking() echo.HandlerFunc
}

type ParkingServices interface {
	PostParking(token *jwt.Token, newData Parking) error
	UpdateParking(parkingID int, token *jwt.Token, newData Parking) error
	GetPicture(parkingID int) (Parking, error)
	GetParking(token *jwt.Token, parkingID uint) (Parking, error)
	GetAllParking(userID uint) ([]Parking, error)
	GetAllSlotByID(parkingID uint) ([]parkingslot.ParkingSlot, error)
}

type ParkingModel interface {
	PostParking(newData Parking, email string) error
	GetDataByEmail(email string) (user.User, error)
	GetDataParkingByID(parkingID uint) (Parking, error)
	Update(parkingID int, updateFields map[string]interface{}, userID uint) error
	GetPicture(parkingID int) (Parking, error)
	GetAllParking(userID uint) ([]Parking, error)
	GetAllSlotByID(parkingID uint) ([]parkingslot.ParkingSlot, error)
}

type Parking struct {
	gorm.Model
	ImageLoc string `json:"imageloc" form:"imageloc"`
	Location string `json:"location" form:"location"`
	City     string `json:"city" form:"city"`
	UserID   uint   `json:"user_id" form:"user_id"`
}
type GetParking struct {
	gorm.Model
	ImageLoc string `json:"imageloc" form:"imageloc"`
	Location string `json:"location" form:"location"`
	City     string `json:"city" form:"city"`
	UserID   uint   `json:"user_id" form:"user_id"`
	// ParkingSlots []parkingslot.ParkingSlot `json:"parking_slots"`
}

type AddParkingVal struct {
	ImageLoc string `validate:"required" form:"imageloc"`
	Location string `validate:"required" form:"location"`
	City     string `validate:"required" form:"city"`
	User_ID  uint
}
