package parkingslot

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ParkingSlotController interface {
	Add() echo.HandlerFunc
	AllParkingSlot() echo.HandlerFunc
	Edit() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type ParkingSlotModel interface {
	Add(email string, newSlot ParkingSlot) error
	AllParkingSlot(email string) ([]ParkingSlot, error)
	Edit(email string, parkingSlotID string, editSlot ParkingSlot) error
	Delete(email string, parkingSlotID string) error
}

type ParkingSlotServices interface {
	Add(token *jwt.Token, newSlot ParkingSlot) error
	AllParkingSlot(token *jwt.Token) ([]ParkingSlot, error)
	Edit(token *jwt.Token, parkingSlotID string, editSlot ParkingSlot) error
	Delete(token *jwt.Token, parkingSlotID string) error
}

type ParkingSlot struct {
	ID          uint
	Email       string
	ParkingID   uint   `validate:"required,num"`
	VehicleType string `validate:"required"`
	Floor       int    `validate:"required,num"`
	Slot        int    `validate:"required,num"`
	Price       int    `validate:"required,num"`
	Status      string
}
