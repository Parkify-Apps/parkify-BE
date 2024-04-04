package handler

import (
	"log"
	"net/http"
	"parkify-BE/features/parkingslot"
	"parkify-BE/helper"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type controller struct {
	s parkingslot.ParkingSlotServices
}

func NewHandler(service parkingslot.ParkingSlotServices) parkingslot.ParkingSlotController {
	return &controller{
		s: service,
	}
}

func (ct *controller) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AddParkingSlotRequest
		err := c.Bind(&input)
		if err != nil {
			log.Println("error bind data:", err.Error())
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType, helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.UserInputFormatError, nil))
			}
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		token, ok := c.Get("user").(*jwt.Token)
		defer func() {
			if err := recover(); err != nil {
				log.Println("error jwt creation:", err)

			}
		}()
		if !ok {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		var status string = "available"
		var inputProcess parkingslot.ParkingSlot
		inputProcess.ParkingID = input.ParkingID
		inputProcess.VehicleType = input.VehicleType
		inputProcess.Floor = input.Floor
		inputProcess.Slot = input.Slot
		inputProcess.Price = input.Price
		inputProcess.Status = status

		err = ct.s.Add(token, inputProcess)
		if err != nil {
			log.Println("error insert db:", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "success add parking slot", nil))
	}
}

func (ct *controller) AllParkingSlot() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		defer func() {
			if err := recover(); err != nil {
				log.Println("error jwt creation:", err)

			}
		}()
		if !ok {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		result, err := ct.s.AllParkingSlot(token)
		if err != nil {
			var code = http.StatusInternalServerError
			if strings.Contains(err.Error(), "validation") || strings.Contains(err.Error(), "cek kembali") {
				code = http.StatusBadRequest
			}
			return c.JSON(code, helper.ResponseFormat(code, err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "success get all parking slot", result))
	}
}

func (ct *controller) Edit() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input parkingslot.ParkingSlot
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType, helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.UserInputFormatError, nil))
			}
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		token, ok := c.Get("user").(*jwt.Token)
		defer func() {
			if err := recover(); err != nil {
				log.Println("error jwt creation:", err)

			}
		}()
		if !ok {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		parkingslotID := c.Param("id")

		var updateProcess parkingslot.ParkingSlot
		updateProcess.Price = input.Price

		err = ct.s.Edit(token, parkingslotID, updateProcess)
		if err != nil {
			var code = http.StatusInternalServerError
			if strings.Contains(err.Error(), "validation") || strings.Contains(err.Error(), "cek kembali") {
				code = http.StatusBadRequest
			}
			return c.JSON(code, helper.ResponseFormat(code, err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "success edit parking slot", nil))
	}
}

func (ct *controller) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		defer func() {
			if err := recover(); err != nil {
				log.Println("error jwt creation:", err)

			}
		}()
		if !ok {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		parkingSlotID := c.Param("parkingSlotID")

		err := ct.s.Delete(token, parkingSlotID)
		if err != nil {
			var code = http.StatusInternalServerError
			if strings.Contains(err.Error(), "validation") || strings.Contains(err.Error(), "cek kembali") {
				code = http.StatusBadRequest
			}
			return c.JSON(code, helper.ResponseFormat(code, err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "success delete parking slot", nil))
	}
}
