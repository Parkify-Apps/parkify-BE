package services

import (
	"errors"
	"log"
	"parkify-BE/features/reservation"
	"parkify-BE/helper"
	"parkify-BE/middlewares"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type services struct {
	m reservation.ReservationModel
	v *validator.Validate
}

func ReservationService(model reservation.ReservationModel) reservation.ReservationServices {
	return &services{
		m: model,
		v: validator.New(),
	}
}

func (s *services) Create(token *jwt.Token, newData reservation.Reservation) error {
	decodeEmail := middlewares.DecodeToken(token)
	if decodeEmail == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	decodeRole := middlewares.DecodeRole(token)
	if decodeRole == "operator" {
		log.Println("role restricted:", "operator tidak bisa mengakses fitur ini")
		return errors.New("operator tidak bisa mengakses fitur ini")

	} else if decodeRole == "user" {
		err := s.m.Create(decodeEmail, newData)
		if err != nil {
			return errors.New(helper.ServerGeneralError)
		}
	}

	return nil
}

func (s *services) GetHistory(token *jwt.Token) ([]reservation.Reservation, error) {
	decodeEmail := middlewares.DecodeToken(token)
	if decodeEmail == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return nil, errors.New("data tidak valid")
	}

	result, err := s.m.GetHistory(decodeEmail)
	if err != nil {
		return nil, errors.New(helper.ServerGeneralError)
	}

	return result, nil
}

func (s *services) GetReservationInfo(token *jwt.Token, reservationID string) (reservation.Reservation, error) {
	decodeEmail := middlewares.DecodeToken(token)
	if decodeEmail == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return reservation.Reservation{}, errors.New("data tidak valid")
	}

	result, err := s.m.GetReservationInfo(decodeEmail, reservationID)
	if err != nil {
		return reservation.Reservation{}, errors.New(helper.ServerGeneralError)
	}

	return result, nil
}
