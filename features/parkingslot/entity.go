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
	AllParkingSlot() ([]ParkingSlot, error)
	Edit(email string, parkingSlotID string, editSlot ParkingSlot) error
	Delete(email string, parkingSlotID string) error
}

type ParkingSlotServices interface {
	Add(token *jwt.Token, newSlot ParkingSlot) error
	AllParkingSlot() ([]ParkingSlot, error)
	Edit(token *jwt.Token, parkingSlotID string, editSlot ParkingSlot) error
	Delete(token *jwt.Token, parkingSlotID string) error
}

type ParkingSlot struct {
	Email       string
	ParkingID   uint
	VehicleType string
	Floor       int
	Slot        int
	Price       int
	Status      string
}
