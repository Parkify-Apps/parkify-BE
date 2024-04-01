package services

import (
	"errors"
	"log"
	"parkify-BE/features/parking"
	"parkify-BE/helper"
	"parkify-BE/middlewares"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	m parking.ParkingModel
	v *validator.Validate
}

func NewService(model parking.ParkingModel) parking.ParkingServices {
	return &service{
		m: model,
		v: validator.New(),
	}
}

func (s *service) PostParking(token *jwt.Token, newData parking.Parking) error {
	email := middlewares.DecodeToken(token)
	if email == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	user, err := s.m.GetDataByEmail(email)
	if err != nil {
		log.Println("error getting user:", err)
		return err
	}
	

	newData.User_ID = user.ID

	var parkingValidate parking.AddParkingVal
	parkingValidate.ImageLoc = newData.ImageLoc
	parkingValidate.Location = newData.Location
	parkingValidate.City = newData.City
	parkingValidate.User_ID = newData.User_ID

	err = s.v.Struct(&parkingValidate)
	if err != nil {
		log.Println("error validasi", err.Error())
		return err
	}

	err = s.m.PostParking(newData, email)
	if err != nil {
		log.Print("error update to model: ", err.Error())
		return errors.New(helper.ServerGeneralError)
	}
	return nil
}
