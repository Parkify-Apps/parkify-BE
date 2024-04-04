package data

import (
	"errors"
	"parkify-BE/features/parkingslot"
	"strconv"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) parkingslot.ParkingSlotModel {
	return &model{
		connection: db,
	}
}

func (psm *model) Add(email string, newSlot parkingslot.ParkingSlot) error {
	var inputProcess = ParkingSlot{
		Email:       email,
		ParkingID:   newSlot.ParkingID,
		VehicleType: newSlot.VehicleType,
		Floor:       newSlot.Floor,
		Slot:        newSlot.Slot,
		Price:       newSlot.Price,
		Status:      newSlot.Status,
	}

	qry := psm.connection.Model(&inputProcess).Where("email = ? AND parking_id = ?", email, newSlot.ParkingID).Create(&inputProcess)
	if err := qry.Error; err != nil {
		return err
	}

	if qry.RowsAffected < 1 {
		return errors.New("no data affected")
	}
	return nil
}

func (psm *model) AllParkingSlot(email string) ([]parkingslot.ParkingSlot, error) {
	var AllSlot []parkingslot.ParkingSlot

	err := psm.connection.Where("email = ?", email).Find(&AllSlot).Error
	if err != nil {
		return nil, err
	}

	return AllSlot, err
}

func (psm *model) Edit(email string, parkingSlotID string, editSlot parkingslot.ParkingSlot) error {
	slotID, _ := strconv.ParseUint(parkingSlotID, 10, 64)

	qry := psm.connection.Model(&ParkingSlot{}).Where("email = ? AND id = ?", email, slotID).Updates(&editSlot)
	if err := qry.Error; err != nil {
		return err
	}

	if qry.RowsAffected < 1 {
		return errors.New("no data affected")
	}

	return nil
}

func (psm *model) Delete(email string, parkingSlotID string) error {

	qry := psm.connection.Where("email = ? AND id = ?", email, parkingSlotID).Delete(&parkingslot.ParkingSlot{})

	if err := qry.Error; err != nil {
		return err
	}

	if qry.RowsAffected < 1 {
		return errors.New("no data affected")
	}

	return nil
}
