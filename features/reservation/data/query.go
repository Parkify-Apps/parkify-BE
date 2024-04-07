package data

import (
	"errors"
	"parkify-BE/features/reservation"

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

	// err := rm.connection.Model(&data.ParkingSlot{}).Where("id = ? AND email = ?", newData.ParkingSlotID, email).Update("status", "not available").Error
	// if err != nil {
	// 	return err
	// }

	return reservationResponse, nil
}

func (rm *model) GetHistory(email string) ([]reservation.Reservation, error) {
	var GetHistory []reservation.Reservation

	err := rm.connection.Model(&Reservation{}).Where("email = ?", email).Preload("ParkingSlot").Order("id desc").Find(&GetHistory).Error
	if err != nil {
		return nil, err
	}

	return GetHistory, err
}

func (rm *model) GetReservationInfo(email string, reservationID string) (reservation.Reservation, error) {
	var GetReservationInfo reservation.Reservation

	err := rm.connection.Model(&Reservation{}).Where("id = ? AND email = ?", reservationID, email).Preload("ParkingSlot").First(&GetReservationInfo).Error
	if err != nil {
		return reservation.Reservation{}, err
	}

	return GetReservationInfo, err
}
