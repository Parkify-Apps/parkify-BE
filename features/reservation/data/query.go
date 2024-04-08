package data

import (
	"errors"
	"parkify-BE/features/parkingslot/data"
	"parkify-BE/features/reservation"
	"strconv"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) reservation.ReservationModel {
	return &model{
		connection: db,
	}
}

func (rm *model) Create(email string, newData reservation.Reservation) (reservation.Reservation, error) {
	var inputProcess = Reservation{
		ExitedAt:      newData.ExitedAt,
		Email:         email,
		ParkingSlotID: newData.ParkingSlotID,
	}

	qry := rm.connection.Create(&inputProcess)
	if err := qry.Error; err != nil {
		return reservation.Reservation{}, err
	}

	if qry.RowsAffected < 1 {
		return reservation.Reservation{}, errors.New("no data affected")
	}

	var reservationResponse reservation.Reservation
	if err := rm.connection.Where("email = ? AND parking_slot_id = ?", email, newData.ParkingSlotID).First(&reservationResponse).Error; err != nil {
		return reservation.Reservation{}, err
	}

	var qry2 = rm.connection.Model(&data.ParkingSlot{}).Where("id = ?", newData.ParkingSlotID).Update("status", "not available")

	err := qry2.Error
	if err != nil {
		return reservation.Reservation{}, err
	}

	if qry2.RowsAffected < 1 {
		return reservation.Reservation{}, errors.New("no data affected")
	}

	return reservationResponse, nil
}

func (rm *model) GetHistory(email string) ([]reservation.ReservationResponse, error) {
	var response []reservation.ReservationResponse

	err := rm.connection.Raw("SELECT reservations.id as id, reservations.email, reservations.parking_slot_id, parking_slots.vehicle_type, parking_slots.floor, parking_slots.slot, parking_slots.price, parkings.id as parking_id, parkings.image_loc, parkings.location, parkings.city FROM reservations JOIN parking_slots ON reservations.parking_slot_id = parking_slots.id JOIN parkings ON parking_slots.parking_id = parkings.id where reservations.email = ? order by 1 desc", email).Scan(&response).Error
	if err != nil {
		return nil, err
	}

	return response, err
}

func (rm *model) GetReservationInfo(email string, reservationID string) (reservation.ReservationResponse, error) {
	var getReservationInfo reservation.ReservationResponse
	rsvpID, _ := strconv.ParseUint(reservationID, 10, 64)

	err := rm.connection.Raw("SELECT reservations.id as id, reservations.email, reservations.parking_slot_id, parking_slots.vehicle_type, parking_slots.floor, parking_slots.slot, parking_slots.price, parkings.id as parking_id, parkings.image_loc, parkings.location, parkings.city FROM reservations JOIN parking_slots ON reservations.parking_slot_id = parking_slots.id JOIN parkings ON parking_slots.parking_id = parkings.id where reservations.email = ? AND reservations.id = ?", email, rsvpID).Scan(&getReservationInfo).Error

	if err != nil {
		return reservation.ReservationResponse{}, err
	}

	return getReservationInfo, err
}
