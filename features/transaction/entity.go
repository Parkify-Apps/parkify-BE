package transaction

import (
	"parkify-BE/features/parking"
	"parkify-BE/features/parkingslot"
	"parkify-BE/features/reservation"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TransactionController interface {
	Transaction() echo.HandlerFunc
	PaymentCallback() echo.HandlerFunc
	Get() echo.HandlerFunc
}

type TransactionServices interface {
	Transaction(payment PaymentRequest, token *jwt.Token) (any, error)
	PaymentCallback(payment CallbackRequest) error
	Get(id int, token *jwt.Token) (any, error)
}

type TransactionModel interface {
	GetReservation(id uint) (reservation.Reservation, error)
	GetParkingSlot(id uint) (parkingslot.ParkingSlot, error)
	GetParking(id uint) (parking.Parking, error)
	CreateTransaction(newData Transaction, reservationID uint) (Transaction, error)
	UpdateSuccess(newData Transaction, orderID uint) error
	UpdateAvailable(newData parkingslot.ParkingSlot, slotID uint) error
	Get(id int) (Transaction, error)
	GetIDByOrderID(orderID string) (Transaction, error)
}

type Transaction struct {
	gorm.Model
	OrderID string
	ReservationID  uint
	PaymentMethod  string
	Price          int
	Status         string
	VirtualAccount string
}

type PaymentRequest struct {
	ReservationID uint   `json:"reservation_id"`
	PaymentMethod string `json:"payment_method"`
}

type CallbackRequest struct {
	OrderID           string `json:"order_id"`
	GrossAmount       string `json:"gross_amount"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
	TransactionStatus string `json:"transaction_status"`
}

type CallbackResponse struct {
	TransactionStatus string `json:"transaction_status"`
	StatusMessage     string `json:"status_message"`
	FraudStatus       string `json:"fraud_status"`
}

// type PaymentResponse struct {
// 	VirtualAccount []coreapi.VANumber
// 	TransactionID  string
// }
