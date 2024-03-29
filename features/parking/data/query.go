package data

import (
	"parkify-BE/features/parking"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) parking.ParkingModel {
	return &model{
		connection: db,
	}
}

func (m *model) PostParking(newData parking.Parking, email string) error {
	err := m.connection.Model(&newData).Where("email = ?", email).Create(&newData).Error
	if err != nil {
		return err
	}
	return nil
}
