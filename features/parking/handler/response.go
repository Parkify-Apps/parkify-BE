package handler

import (
	"parkify-BE/features/parkingslot"
)

type ParkingResponse struct {
	Location     string                    `json:"location"`
	City         string                    `json:"city"`
	ImageLoc     string                    `json:"image_loc"`
	ParkingSlots []parkingslot.ParkingSlot `json:"parking_slots"`
}
