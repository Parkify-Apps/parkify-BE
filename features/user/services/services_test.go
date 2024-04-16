package services_test

import (
	"parkify-BE/features/user"
	"parkify-BE/features/user/services"
	"parkify-BE/helper"
	"parkify-BE/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	// "gorm.io/datatypes"
	// "golang.org/x/crypto/bcrypt"
)

func TestRegister(t *testing.T) {
	model := mocks.NewUserModel(t)
	pm := mocks.NewPasswordManager(t)
	md := mocks.NewJwtInterface(t)
	rPm := helper.NewPasswordManager()
	srv := services.NewService(model, pm, md)

	createdAtStr := "2024-04-07 19:32:56.650"
	layout := "2006-01-02 15:04:05.000"
	createdAt, _ := time.Parse(layout, createdAtStr)

	updatedAtStr := "2024-04-07 19:38:56.650"
	updatedAt, _ := time.Parse(layout, updatedAtStr)

	currentTime := time.Now()
	deletedAt := gorm.DeletedAt{
		Time:  currentTime,
		Valid: false,
	}

	// var registerData user.User
	// registerData.ID = 1
	// registerData.CreatedAt = createdAt
	// registerData.UpdatedAt = updatedAt
	// registerData.DeletedAt = deletedAt
	// registerData.Fullname = "rizal"
	// registerData.Email = "rizal@gmail.com"
	// registerData.Password = "hahahaha"
	// registerData.Role = "user"
	registerData := user.User{Model: gorm.Model{
		ID: 1, CreatedAt: createdAt, UpdatedAt: updatedAt, DeletedAt: deletedAt},
		Fullname: "rizal", Email: "rizal@gmail.com", Password: "hahahaha", Role: "user"}
	hashedPassword, _ := rPm.HashPassword(registerData.Password)

	// var insertData user.User
	// insertData.ID = 1
	// insertData.CreatedAt = createdAt
	// insertData.UpdatedAt = updatedAt
	// insertData.DeletedAt = deletedAt
	// insertData.Fullname = "rizal"
	// insertData.Email = "rizal@gmail.com"
	// insertData.Password = "hahahaha"
	// insertData.Role = "user"
	insertData := user.User{Model: gorm.Model{
		ID: 1, CreatedAt: createdAt, UpdatedAt: updatedAt, DeletedAt: deletedAt},
		Fullname: "rizal", Email: "rizal@gmail.com", Password: "hahahaha", Role: "user"}

	t.Run("Success register", func(t *testing.T) {
		pm.On("HashPassword", registerData.Password).Return(hashedPassword, nil).Once()
		insertData.Password = hashedPassword
		model.On("InsertUser", insertData).Return(nil).Once()

		err := srv.Register(registerData)

		pm.AssertExpectations(t)
		model.AssertExpectations(t)

		assert.Nil(t, err)
	})

	// t.Run("Error hash password", func(t *testing.T) {
	// 	pm.On("HashPassword", registerData.Password).Return("", bcrypt.ErrHashTooShort).Once()

	// 	err := srv.Register(registerData)

	// 	pm.AssertExpectations(t)

	// 	assert.Error(t, err)
	// })

	// t.Run("Error from model", func(t *testing.T) {
	// 	pm.On("HashPassword", registerData.Password).Return(hashedPassword, nil).Once()
	// 	insertData.Password = hashedPassword
	// 	model.On("InsertUser", insertData).Return(gorm.ErrInvalidData).Once()

	// 	err := srv.Register(registerData)

	// 	pm.AssertExpectations(t)
	// 	model.AssertExpectations(t)

	// 	assert.Error(t, err)
	// 	assert.EqualError(t, err, helper.ServerGeneralError)
	// })

	// t.Run("Error validation", func(t *testing.T) {
	// 	registerData.Password = "hahahaha"
	// 	err := srv.Register(registerData)

	// 	assert.Error(t, err)
	// })
}
