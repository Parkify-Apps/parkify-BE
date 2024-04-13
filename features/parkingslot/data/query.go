package data

import (
	"errors"
	"log"
	"parkify-BE/features/parkingslot"

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
	var existingSlot ParkingSlot
	if err := psm.connection.
		Model(&ParkingSlot{}).
		Where("floor = ? AND slot = ?", newSlot.Floor, newSlot.Slot).
		First(&existingSlot).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err // Error other than "record not found"
		}
	} else {
		return errors.New("data already exists")
	}

	// Create new entry
	var inputProcess = ParkingSlot{
		Email:       email,
		ParkingID:   newSlot.ParkingID,
		VehicleType: newSlot.VehicleType,
		Floor:       newSlot.Floor,
		Slot:        newSlot.Slot,
		Price:       newSlot.Price,
		Status:      newSlot.Status,
	}

	qry := psm.connection.Create(&inputProcess)
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

	err := psm.connection.Raw("SELECT * FROM parking_slots WHERE parking_slots.email = ? AND parking_slots.deleted_at IS NULL", email).Find(&AllSlot).Error
	if err != nil {
		return nil, err
	}

	return AllSlot, err
}

func (psm *model) Edit(email string, parkingSlotID string, editSlot parkingslot.ParkingSlot) error {

	qry := psm.connection.Model(&ParkingSlot{}).Where("email = ? AND id = ?", email, parkingSlotID).Updates(&editSlot)
	if err := qry.Error; err != nil {
		return err
	}

	if qry.RowsAffected < 1 {
		return errors.New("no data affected")
	}

	return nil
}

func (psm *model) Delete(email string, parkingSlotID string) error {

	qry := psm.connection.Model(&ParkingSlot{}).Where("email = ? AND id = ?", email, parkingSlotID).Delete(&parkingSlotID)

	if err := qry.Error; err != nil {
		log.Print("error to database :", err.Error())
		return err
	}

	if qry.RowsAffected < 1 {
		return errors.New("no data affected")
	}

	return nil
}
