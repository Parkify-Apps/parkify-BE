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

	var parkingValidate parking.AddParkingVal
	parkingValidate.ImageLoc = newData.ImageLoc
	parkingValidate.Location = newData.Location
	parkingValidate.LocationName = newData.LocationName

	err := s.v.Struct(&parkingValidate)
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
