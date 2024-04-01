package data

import (
	"parkify-BE/features/parking"
	"parkify-BE/features/user"

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

func (m *model) GetDataByEmail(email string) (user.User, error) {
	var result user.User
	if err := m.connection.Where("email = ?", email).First(&result).Error; err != nil {
		return user.User{}, err
	}
	return result, nil
}
