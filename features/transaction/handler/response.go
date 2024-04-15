package handler

import "github.com/midtrans/midtrans-go/coreapi"

type PaymentResponse struct {
	OrderID        string
	TransActionID  uint               `json:"transaction_id"`
	VirtualAccount []coreapi.VANumber `json:"virtual_account"`
	// TransactionID  string `json:"transaction_id`
	ParkingID     uint   `json:"parking_id"`
	ParkingSlotID uint   `json:"parkingslot_id"`
	Location      string `json:"location"`
	City          string `json:"city"`
	VehicleType   string `json:"vehicle_type"`
	Floor         int    `json:"floor"`
	Slot          int    `json:"slot"`
	Price         int    `json:"price"`
	StatusMessage string `json:"status_message"`
}

type FinishPaymentResponse struct {
	OrderID        string
	VirtualAccount string `json:"virtual_account"`
	PaymentMethod  string `json:"payment_method"`
	ParkingID      uint   `json:"parking_id"`
	ParkingSlotID  uint   `json:"parkingslot_id"`
	ReservationID  uint   `json:"reservation_id"`
	Location       string `json:"location"`
	City           string `json:"city"`
	VehicleType    string `json:"vehicle_type"`
	Floor          int    `json:"floor"`
	Slot           int    `json:"slot"`
	Price          int    `json:"price"`
	ImageLoc       string `json:"image_loc"`
	// StatusMessage  string `json:"status_message"`
}
