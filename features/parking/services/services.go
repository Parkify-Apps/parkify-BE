package services

import (
	"errors"
	"log"
	"parkify-BE/features/parking"
	"parkify-BE/features/parkingslot"
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

	newData.UserID = user.ID

	var parkingValidate parking.AddParkingVal
	parkingValidate.ImageLoc = newData.ImageLoc
	parkingValidate.Location = newData.Location
	parkingValidate.City = newData.City
	parkingValidate.User_ID = newData.UserID

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

func (s *service) UpdateParking(parkingID int, token *jwt.Token, newData parking.Parking) error {
	email := middlewares.DecodeToken(token)
	if email == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	u, err := s.m.GetDataByEmail(email)
	if err != nil {
		log.Println("error getting user:", err)
		return err
	}

	user, error := s.m.GetDataParkingByID(uint(parkingID))
	log.Print(user)
	if error != nil {
		log.Println("error getting parking:", error.Error())
		return error
	}

	if u.ID != user.UserID {
		log.Println("error get account:", "user tidak sesuai")
		return errors.New("user tidak sesuai")
	}

	err = s.v.Struct(&newData)
	if err != nil {
		log.Println("error validasi", err.Error())
		return err
	}
	// membuat map untuk menampung kolom yang akan diperbarui bersama dengan nilainya
	updateFields := make(map[string]interface{})

	// tentukan kolom yang ingin diperbarui dan tambahkan ke dalam map
	if newData.Location != "" {
		updateFields["location"] = newData.Location
	}
	if newData.City != "" {
		updateFields["city"] = newData.City
	}
	if newData.ImageLoc != "" {
		updateFields["image_loc"] = newData.ImageLoc
	}

	err = s.m.Update(parkingID, updateFields, user.UserID)
	if err != nil {
		log.Print("error update to model: ", err.Error())
		return errors.New(helper.ServerGeneralError)
	}

	return nil
}

func (s *service) GetPicture(parkingID int) (parking.Parking, error) {
	result, err := s.m.GetPicture(parkingID)
	if err != nil {
		log.Println("error getting user:", err.Error())
		return parking.Parking{}, err
	}

	return result, err
}

func (s *service) GetParking(token *jwt.Token, parkingID uint) (parking.Parking, error) {
	email := middlewares.DecodeToken(token)
	if email == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return parking.Parking{}, errors.New("data tidak valid")
	}

	result, err := s.m.GetDataParkingByID(parkingID)
	if err != nil {
		return parking.Parking{}, err
	}

	return result, nil
}

func (s *service) GetAllParking(userID uint) ([]parking.Parking, error) {
	result, err := s.m.GetAllParking(userID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) GetAllSlotByID(parkingID uint) ([]parkingslot.ParkingSlot, error) {
	result, err := s.m.GetAllSlotByID(parkingID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
