package data

import (
	"errors"
	"log"
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

func (m *model) GetDataParkingByID(parkingID uint) (parking.Parking, error) {
	var result parking.Parking
	if err := m.connection.Where("id = ?", parkingID).First(&result).Error; err != nil {
		return parking.Parking{}, err
	}
	return result, nil
}

func (m *model) Update(parkingID int, updateFields map[string]interface{}, userID uint) error {
	var query = m.connection.Model(&Parking{}).Where("id = ? AND user_id = ?", parkingID, userID).Updates(updateFields)
	if err := query.Error; err != nil {
		log.Print("error to database :", err.Error())
		return err
	}
	if query.RowsAffected < 1 {
		return errors.New("no data affected")
	}
	log.Print(parkingID, userID)
	return nil
}

func (m *model) GetPicture(parkingID int) (parking.Parking, error) {
	var result parking.Parking
	if err := m.connection.Where("id = ?", parkingID).First(&result).Error; err != nil {
		return parking.Parking{}, err
	}
	return result, nil
}

func (m *model) GetAllParking(parkingID int) ([]parking.Parking, error) {
	var result []parking.Parking
	if err := m.connection.Where("id = ?", parkingID).Error; err != nil {
		return nil, err
	}
	return result, nil
}
