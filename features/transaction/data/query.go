package data

import (
	
	"parkify-BE/features/parking"
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

func (m *model) GetParking(id uint) (parking.Parking, error) {
	var result parking.Parking
	if err := m.connection.Where("id = ?", id).Find(&result).Error; err != nil {
		return parking.Parking{}, err
	}
	return result, nil
}

func (m *model) CreateTransaction(newData transaction.Transaction, reservationID uint) (transaction.Transaction, error) {
	if err := m.connection.Create(&newData).Error; err != nil {
		return transaction.Transaction{}, err
	}

	var response transaction.Transaction
	if err := m.connection.Where("reservation_id = ?", reservationID).Find(&response).Error; err != nil {
		return transaction.Transaction{}, err
	}

	return response, nil
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

func (m *model) UpdateAvailable(newData parkingslot.ParkingSlot, slotID uint) error {
	var qry = m.connection.Where("id = ?", slotID).Updates(&newData)
	if err := qry.Error; err != nil {
		return err
	}
	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (m *model) Get(id int) (transaction.Transaction, error) {
	var result transaction.Transaction
	if err := m.connection.Where("id = ?", id).First(&result).Error; err != nil {
		return transaction.Transaction{}, err
	}
	return result, nil
}