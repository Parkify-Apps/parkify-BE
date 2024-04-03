package data

import (
	"errors"
	"log"
	user "parkify-BE/features/user"

	"gorm.io/gorm"
)

type UserModel struct {
	connection *gorm.DB
}

func New(db *gorm.DB) user.UserModel {
	return &UserModel{
		connection: db,
	}
}

func (um *UserModel) AddUser(newData user.User) error {
	err := um.connection.Create(&newData).Error
	if err != nil {
		return err
	}
	return nil
}

func (um *UserModel) Login(email string) (user.User, error) {
	var result user.User
	if err := um.connection.Where("email = ? ", email).First(&result).Error; err != nil {
		return user.User{}, err
	}
	return result, nil
}

func (um *UserModel) Profile(email string) (user.User, error) {
	var result user.User
	if err := um.connection.Where("email = ?", email).First(&result).Error; err != nil {
		return user.User{}, err
	}
	return result, nil
}

func (um *UserModel) UpdateProfile(email string, data user.User) error {
	var query = um.connection.Model(&User{}).Where("email = ?", email).Updates(&data)
	if err := query.Error; err != nil {
		log.Print("error to database :", err.Error())
		return err
	}
	if query.RowsAffected < 1 {
		return errors.New("no data affected")
	}
	return nil
}

func (um *UserModel) Delete(email string) error {
	if err := um.connection.Model(&User{}).Where("email = ?", email).Delete(email).Error; err != nil {
		log.Print("error to database :", err.Error())
		return err
	}
	return nil
}
