package handler

type AddParkingSlotRequest struct {
	ParkingID   uint   `json:"parking_id"`
	VehicleType string `json:"vehicle_type"`
	Floor       int    `json:"floor"`
	Slot        int    `json:"slot"`
	Price       int    `json:"price"`
	Status      string `json:"status"`
}
