package services_test

import (
	"parkify-BE/features/parkingslot"
	"parkify-BE/features/parkingslot/services"
	"parkify-BE/mocks"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	model := mocks.NewParkingSlotModel(t)
	md := mocks.NewJwtInterface(t)
	srv := services.ParkingSlotService(model, md)

	token := &jwt.Token{}

	newData := parkingslot.ParkingSlot{ParkingID: 1, VehicleType: "Car", Floor: 1, Slot: 1, Price: 10000, Status: "available"}

	t.Run("Success Add Parking Slot", func(t *testing.T) {
		md.On("DecodeToken", token).Return("khomsin@gmail.com", nil)
		md.On("DecodeRole", token).Return("operator")
		model.On("Add", "khomsin@gmail.com", newData).Return(nil).Once()

		// Calling the method under test with the generated/mocked token
		err := srv.Add(token, newData)

		// Asserting expectations and results
		model.AssertExpectations(t)
		assert.Nil(t, err)
	})
}
