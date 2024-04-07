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
	GetHistory(email string) ([]Reservation, error)
	GetReservationInfo(email string, reservationID string) (Reservation, error)
}

type ReservationServices interface {
	Create(token *jwt.Token, newData Reservation) (Reservation, error)
	GetHistory(token *jwt.Token) ([]Reservation, error)
	GetReservationInfo(token *jwt.Token, reservationID string) (Reservation, error)
}

type Reservation struct {
	gorm.Model
	ExitedAt      time.Time        `json:"exited_at"`
	Email         string           `json:"email"`
	ParkingSlotID uint             `json:"parkingslot_id"`
	ParkingSlot   data.ParkingSlot `json:"parkingslot"`
}
