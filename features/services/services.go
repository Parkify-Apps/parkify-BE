package services

import (
	"errors"
	"log"
	user "parkify-BE/features"
	"parkify-BE/helper"
	"parkify-BE/middlewares"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	// "github.com/labstack/gommon/email"
)

type service struct {
	model user.UserModel
	pm    helper.PasswordManager
	v     *validator.Validate
}

func NewService(m user.UserModel) user.UserService {
	return &service{
		model: m,
		pm:    helper.NewPasswordManager(),
		v:     validator.New(),
	}
}

func (s *service) Register(newData user.User) error {
	var registerValidate user.Register

	// Mencari UserID terakhir dari database
	lastUserID, error := s.model.GetLastUserID()
	if error != nil {
		return errors.New(helper.ServiceGeneralError)
	}

	// Menentukan UserID untuk pengguna baru
	newUserID := lastUserID + 1
	newData.UserID = newUserID

	registerValidate.UserID = newData.UserID
	registerValidate.Email = newData.Email
	registerValidate.Nama = newData.Nama
	registerValidate.Password = newData.Password
	err := s.v.Struct(&registerValidate)
	if err != nil {
		log.Println("error validasi", err.Error())
		return err
	}

	newPassword, err := s.pm.HashPassword(newData.Password)
	if err != nil {
		return errors.New(helper.ServiceGeneralError)
	}
	newData.Password = newPassword

	err = s.model.AddUser(newData)
	if err != nil {
		return errors.New(helper.ServerGeneralError)
	}
	return nil
}

func (s *service) Login(loginData user.User) (user.User, string, error) {
	var loginValidate user.Login
	loginValidate.Email = loginData.Email
	loginValidate.Password = loginData.Password
	err := s.v.Struct(&loginValidate)
	if err != nil {
		log.Println("error validasi", err.Error())
		return user.User{}, "", err
	}

	dbData, err := s.model.Login(loginValidate.Email)
	if err != nil {
		log.Println("error login model", err.Error())
		return user.User{}, "", errors.New(helper.UserCredentialError) //
	}

	err = s.pm.ComparePassword(loginValidate.Password, dbData.Password)
	if err != nil {
		log.Println("error compare", err.Error())
		return user.User{}, "", errors.New(helper.UserCredentialError)
	}

	token, err := middlewares.GenerateJWT(dbData.Email)
	if err != nil {
		log.Println("error generate", err.Error())
		return user.User{}, "", errors.New(helper.ServiceGeneralError)
	}

	return dbData, token, nil
}

func (s *service) Profile(token *jwt.Token, userID uint) (user.User, error) {
	decodeEmail := middlewares.DecodeToken(token)
	if decodeEmail == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return user.User{}, errors.New("data tidak valid")
	}

	result, err := s.model.GetUserByID(userID)
	if err != nil {
		return user.User{}, err
	}

	if result.Email != decodeEmail {
		return user.User{}, errors.New("anda tidak diizinkan mengakses profil pengguna lainn")
	}

	return result, nil
}
