package data

import (
	"log"
	"parkify-BE/features/parkingslot"
	"parkify-BE/features/reservation"
	"parkify-BE/features/transaction"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) transaction.TransactionModel {
	return &model{
		connection: db,
	}
}

func (m *model) GetReservation(id uint) (reservation.Reservation, error) {
	var result reservation.Reservation
	if err := m.connection.Where("id = ?", id).Last(&result).Error; err != nil {
		return reservation.Reservation{}, err
	}
	return result, nil
}

func (m *model) GetParkingSlot(id uint) (parkingslot.ParkingSlot, error) {
	var result parkingslot.ParkingSlot
	if err := m.connection.Where("id = ?", id).Find(&result).Error; err != nil {
		return parkingslot.ParkingSlot{}, err
	}
	return result, nil
}

func (m *model) CreateTransaction(newData transaction.Transaction) error {
	// var result transaction.Transaction
	log.Print(newData.ReservationID)
	log.Print(newData.Price)
	log.Print(newData.Status)
	if err := m.connection.Create(&newData).Error; err != nil {
		return err
	}
	return nil
}

func (m *model) UpdateSuccess(newData transaction.Transaction, orderID uint) error {
	var qry = m.connection.Where("id = ?", orderID).Updates(&newData)
	if err := qry.Error; err != nil {
		return err
	}
	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
