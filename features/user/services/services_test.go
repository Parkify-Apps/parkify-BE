package services_test

import (
	"parkify-BE/features/user"
	"parkify-BE/features/user/services"

	"parkify-BE/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	model := mocks.NewUserModel(t)
	pm := mocks.NewPasswordManager(t)
	md := mocks.NewJwtInterface(t)
	// rPm := helper.NewPasswordManager()
	srv := services.NewService(model, pm, md)

	registerData := user.User{
		Fullname: "rizal", Email: "rizal@gmail.com", Password: "hahahaha", Role: "user"}
	// hashedPassword, _ := rPm.HashPassword(registerData.Password)

	insertData := user.User{
		Fullname: "rizal", Email: "rizal@gmail.com", Password: "hallo", Role: "user"}

	t.Run("Success register", func(t *testing.T) {
		pm.On("HashPassword", registerData.Password).Return("hallo", nil).Once()
		// insertData.Password = hashedPassword
		model.On("AddUser", insertData).Return(nil).Once()

		err := srv.Register(registerData)

		pm.AssertExpectations(t)
		model.AssertExpectations(t)

		assert.Nil(t, err)
	})

	t.Run("Gagal validasi", func(t *testing.T) {
		regisdata := user.User{
			Fullname: "rizal", Email: "", Password: "hallo", Role: "user"}

		err := srv.Register(regisdata)

		assert.Error(t, err)
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
