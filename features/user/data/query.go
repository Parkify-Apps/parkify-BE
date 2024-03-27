package data

import (
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
