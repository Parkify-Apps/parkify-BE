package reservation

import (
	"parkify-BE/features/parkingslot/data"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ReservationController interface {
	Create() echo.HandlerFunc
	GetHistory() echo.HandlerFunc
	GetReservationInfo() echo.HandlerFunc
}

type ReservationModel interface {
	Create(email string, newData Reservation) (Reservation, error)
	GetHistory(email string) ([]ReservationResponse, error)
	GetReservationInfo(email string, reservationID string) (ReservationResponse, error)
}

type ReservationServices interface {
	Create(token *jwt.Token, newData Reservation) (Reservation, error)
	GetHistory(token *jwt.Token) ([]ReservationResponse, error)
	GetReservationInfo(token *jwt.Token, reservationID string) (ReservationResponse, error)
}

type Reservation struct {
	gorm.Model
	ExitedAt      time.Time        `json:"exited_at"`
	Email         string           `json:"email"`
	ParkingSlotID uint             `json:"parkingslot_id"`
	ParkingSlot   data.ParkingSlot `json:"parkingslot"`
}

type ReservationResponse struct {
	ID            uint   `json:"reservation_id"`
	Email         string `json:"email"`
	ParkingSlotID uint   `json:"parking_slot_id"`
	VehicleType   string `json:"vehicle_type"`
	Floor         int    `json:"floor"`
	Slot          int    `json:"slot"`
	Price         int    `json:"price"`
	ParkingID     uint   `json:"parking_id"`
	ImageLoc      string `json:"image_loc"`
	Location      string `json:"location"`
	City          string `json:"city"`
}
