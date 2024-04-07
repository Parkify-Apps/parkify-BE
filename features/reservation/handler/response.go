package handler

// import "time"

// type GetHistoryResponse struct {
// 	ID            uint      `json:"id"`
// 	CreatedAt     time.Time `json:"created_at"`
// 	ExitedAt      time.Time `json:"exited_at"`
// 	Email         string    `json:"email"`
// 	ParkingSlotID uint      `json:"parkingslot_id"`
// }

type ReservationResponse struct {
	ID uint `json:"reservation_id"`
}
