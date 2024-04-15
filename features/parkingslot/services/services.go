package services

import (
	"errors"
	"log"
	"parkify-BE/features/parkingslot"
	"parkify-BE/helper"
	"parkify-BE/middlewares"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type services struct {
	m parkingslot.ParkingSlotModel
	v *validator.Validate
	md    middlewares.JwtInterface
}

func ParkingSlotService(model parkingslot.ParkingSlotModel, md middlewares.JwtInterface) parkingslot.ParkingSlotServices {
	return &services{
		m: model,
		md: md,
		v: validator.New(),
	}
}

func (s *services) Add(token *jwt.Token, newSlot parkingslot.ParkingSlot) error {
	decodeEmail := s.md.DecodeToken(token)
	if decodeEmail == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	decodeRole := s.md.DecodeRole(token)
	if decodeRole == "user" {
		log.Println("role restricted:", "user tidak bisa mengakses fitur ini")
		return errors.New("user tidak bisa mengakses fitur ini")

	} else if decodeRole == "operator" {
		err := s.m.Add(decodeEmail, newSlot)
		if err != nil {
			return errors.New(helper.ServerGeneralError)
		}
	}

	return nil
}

func (s *services) AllParkingSlot(token *jwt.Token) ([]parkingslot.ParkingSlot, error) {
	decodeEmail := s.md.DecodeToken(token)
	if decodeEmail == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return nil, errors.New("data tidak valid")
	}

	result, err := s.m.AllParkingSlot(decodeEmail)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *services) Edit(token *jwt.Token, parkingSlotID string, editSlot parkingslot.ParkingSlot) error {
	decodeEmail := s.md.DecodeToken(token)
	if decodeEmail == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	err := s.m.Edit(decodeEmail, parkingSlotID, editSlot)
	if err != nil {
		return errors.New(helper.ServerGeneralError)
	}

	return nil
}

func (s *services) Delete(token *jwt.Token, parkingSlotID string) error {
	decodeEmail := s.md.DecodeToken(token)
	if decodeEmail == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	decodeRole := s.md.DecodeRole(token)
	if decodeRole == "user" {
		log.Println("role restricted:", "user tidak bisa mengakses fitur ini")
		return errors.New("user tidak bisa mengakses fitur ini")

	} else if decodeRole == "operator" {
		err := s.m.Delete(decodeEmail, parkingSlotID)
		if err != nil {
			return errors.New(helper.ServerGeneralError)
		}
	}

	return nil
}
