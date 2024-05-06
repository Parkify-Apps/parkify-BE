package services_test

import (
	"parkify-BE/features/user"
	"parkify-BE/features/user/services"
	"parkify-BE/helper"

	"parkify-BE/mocks"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

	t.Run("Error hash password", func(t *testing.T) {
		pm.On("HashPassword", registerData.Password).Return("", bcrypt.ErrHashTooShort).Once()

		err := srv.Register(registerData)

		pm.AssertExpectations(t)

		assert.Error(t, err)
	})

	t.Run("Error from model", func(t *testing.T) {
		pm.On("HashPassword", registerData.Password).Return("hallo", nil).Once()
		// insertData.Password = hashedPassword
		model.On("AddUser", insertData).Return(gorm.ErrInvalidData).Once()

		err := srv.Register(registerData)

		pm.AssertExpectations(t)
		model.AssertExpectations(t)

		assert.Error(t, err)
		assert.EqualError(t, err, helper.ServerGeneralError)
	})

	// t.Run("Error validation", func(t *testing.T) {
	// 	registerData.Password = "hahahaha"
	// 	err := srv.Register(registerData)

	// 	assert.Error(t, err)
	// })
}

func TestLogin(t *testing.T) {
	model := mocks.NewUserModel(t)
	pm := mocks.NewPasswordManager(t)
	md := mocks.NewJwtInterface(t)
	// rPm := helper.NewPasswordManager()
	srv := services.NewService(model, pm, md)

	insertData := user.User{
		Email: "rizal@gmail.com", Password: "hallo"}
	GenerateJWT := user.User{
		Email: "rizal@gmail.com", Password: "hallo", Role: ""}
	Login := user.User{
		Email: "rizal@gmail.com", Password: "hahahaha"}

	t.Run("Success login", func(t *testing.T) {
		// insertData.Password = hashedPassword
		model.On("Login", insertData.Email).Return(user.User{Email: "rizal@gmail.com", Password: "hallo"}, nil).Once()
		pm.On("ComparePassword", insertData.Password, GenerateJWT.Password).Return(nil).Once()
		md.On("GenerateJWT", GenerateJWT.Email, GenerateJWT.Role).Return("rizal@gmail.com, user", nil).Once()

		user, str, err := srv.Login(insertData)

		pm.AssertExpectations(t)
		md.AssertExpectations(t)
		model.AssertExpectations(t)

		assert.Nil(t, err, user, str)
	})

	t.Run("Gagal validasi", func(t *testing.T) {
		validateData := user.User{
			Email: "rizal", Password: ""}

		user, str, err := srv.Login(validateData)

		assert.Error(t, err, user, str)
	})

	t.Run("Error login model", func(t *testing.T) {

		model.On("Login", insertData.Email).Return(user.User{Email: "", Password: ""}, gorm.ErrInvalidData).Once()

		user, str, err := srv.Login(insertData)

		model.AssertExpectations(t)

		assert.Error(t, err, user, str)
		assert.EqualError(t, err, helper.UserCredentialError)
	})

	t.Run("Error compare password", func(t *testing.T) {
		model.On("Login", insertData.Email).Return(Login, nil).Once()
		pm.On("ComparePassword", insertData.Password, Login.Password).Return(gorm.ErrInvalidData).Once()

		user, str, err := srv.Login(insertData)

		pm.AssertExpectations(t)
		model.AssertExpectations(t)

		assert.Error(t, err, user, str)
		assert.EqualError(t, err, helper.UserCredentialError)
	})

	// t.Run("Error generate jwt", func(t *testing.T) {
	// 	md.On("GenerateJWT", GenerateJWT.Email, GenerateJWT.Role).Return("rizal@gmail.com, user", nil).Once()

	// 	user, str, err := srv.Login(insertData)

	// 	md.AssertExpectations(t)

	// 	assert.Error(t, err, user, str)
	// 	assert.EqualError(t, err, helper.UserCredentialError)
	// })
}

func TestProfile(t *testing.T) {
	model := mocks.NewUserModel(t)
	pm := mocks.NewPasswordManager(t)
	md := mocks.NewJwtInterface(t)
	// rPm := helper.NewPasswordManager()
	srv := services.NewService(model, pm, md)

	token := &jwt.Token{}
	email := "rizal@gmail.com"
	profile := user.User{Email: "rizal@gmail.com", Password: "hallo"}

	t.Run("Success get profile", func(t *testing.T) {
		md.On("DecodeToken", token).Return("rizal@gmail.com").Once()
		model.On("Profile", email).Return(profile, nil).Once()

		user, err := srv.Profile(token)

		md.AssertExpectations(t)
		model.AssertExpectations(t)

		assert.Nil(t, err, "seharusnya nilainya nil")
		assert.Equal(t, "rizal@gmail.com", user.Email, "seharusnya nilainya adalah rizal@gmail.com")

	})
}
