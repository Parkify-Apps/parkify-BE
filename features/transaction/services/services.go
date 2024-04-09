package services

import (
	"errors"
	"log"
	"parkify-BE/features/reservation"
	"parkify-BE/features/transaction"
	"parkify-BE/features/transaction/handler"
	"parkify-BE/middlewares"
	"parkify-BE/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	m       transaction.TransactionModel
	v       *validator.Validate
	mdtrans utils.PaymentFunc
}

func NewServices(model transaction.TransactionModel, m utils.PaymentFunc) transaction.TransactionServices {
	return &service{
		m:       model,
		v:       validator.New(),
		mdtrans: m,
	}
}

func (s *service) Transaction(payment transaction.PaymentRequest, token *jwt.Token) (any, error) {
	email := middlewares.DecodeToken(token)
	if email == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return reservation.Reservation{}, errors.New("data tidak valid")
	}

	res, err := s.m.GetReservation(payment.ReservationID)
	if err != nil {
		log.Println("error getting reservation:", err)
		return reservation.Reservation{}, err
	}

	result, err := s.m.GetParkingSlot(res.ParkingSlotID)
	if err != nil {
		log.Println("error getting parking slot:", err)
		return reservation.Reservation{}, err
	}

	resultP, err := s.m.GetParking(result.ParkingID)
	if err != nil {
		log.Println("error getting parking slot:", err)
		return reservation.Reservation{}, err
	}

	num := uint(res.ID)
	str := strconv.FormatUint(uint64(num), 10)

	var response handler.PaymentResponse

	decodeRole := middlewares.DecodeRole(token)
	if decodeRole == "operator" {
		log.Println("role restricted:", "operator tidak bisa mengakses fitur ini")
		return reservation.Reservation{}, errors.New("operator tidak bisa mengakses fitur ini")

	} else if decodeRole == "user" {
		if res.Email != email {
			return nil, errors.New("anda tidak diizinkan mengakses profil pengguna lainn")
		}

		resp, err := s.mdtrans.PaymentVABCA(str, result.Price)
		if err != nil {
			log.Println("error payment:", err)
			return reservation.Reservation{}, err
		}

		var newData transaction.Transaction
		newData.ReservationID = res.ID
		newData.PaymentMethod = "VA BCA"
		newData.Price = result.Price
		// newData.Status = "success"
		newData.VirtualAccount = resp.VaNumbers[0].VANumber

		resultCT, err := s.m.CreateTransaction(newData, res.ID)
		if err != nil {
			log.Println("error create transaction:", err)
			return reservation.Reservation{}, err
		}

		response.TransActionID = resultCT.ID
		response.VirtualAccount = resp.VaNumbers
		// response.TransactionID = resp.TransactionID
		response.City = resultP.City
		response.Location = resultP.Location
		response.Floor = result.Floor
		response.VehicleType = result.VehicleType
		response.Price = result.Price
		response.Slot = result.Slot
		response.ParkingID = result.ParkingID
		response.ParkingSlotID = res.ParkingSlotID
		response.StatusMessage = resp.StatusMessage

	}
	return response, nil
}

func (s *service) PaymentCallback(payment transaction.CallbackRequest) error {
	num, err := strconv.ParseUint(payment.OrderID, 10, 64)
	if err != nil {
		return err
	}

	// _, err = s.m.GetReservation(uint(num))
	// if err != nil {
	// 	log.Println("error getting reservation:", err)
	// 	return err
	// }

	var update transaction.Transaction
	update.Status = "success"
	update.ID = uint(num)
	err = s.m.UpdateSuccess(update, uint(num))
	if err != nil {
		log.Println("error update success:", err)
		return err
	}

	return nil
}

func (s *service) Get(id int, token *jwt.Token) (any, error) {
	var response handler.FinishPaymentResponse

	decodeRole := middlewares.DecodeRole(token)
	if decodeRole == "operator" {
		log.Println("role restricted:", "operator tidak bisa mengakses fitur ini")
		return nil, errors.New("operator tidak bisa mengakses fitur ini")
	} else if decodeRole == "user" {

		email := middlewares.DecodeToken(token)
		if email == "" {
			log.Println("error decode token:", "token tidak ditemukan")
			return reservation.Reservation{}, errors.New("data tidak valid")
		}

		res, err := s.m.GetReservation(uint(id))
		if err != nil {
			log.Println("error getting reservation:", err)
			return reservation.Reservation{}, err
		}

		result, err := s.m.GetParkingSlot(res.ParkingSlotID)
		if err != nil {
			log.Println("error getting parking slot:", err)
			return reservation.Reservation{}, err
		}

		resultP, err := s.m.GetParking(result.ParkingID)
		if err != nil {
			log.Println("error getting parking slot:", err)
			return reservation.Reservation{}, err
		}

		resGet, err := s.m.Get(id)
		if err != nil {
			log.Println("error update success:", err)
			return transaction.Transaction{}, err
		}

		if res.Email != email {
			return nil, errors.New("anda tidak diizinkan mengakses profil pengguna lainn")
		}

		response.VirtualAccount = resGet.VirtualAccount
		response.PaymentMethod = resGet.PaymentMethod
		response.City = resultP.City
		response.Location = resultP.Location
		response.VehicleType = result.VehicleType
		response.Floor = result.Floor
		response.Slot = result.Slot
		response.Price = result.Price
		response.ParkingID = result.ParkingID
		response.ParkingSlotID = res.ParkingSlotID
		response.ReservationID = res.ID

	}
	return response, nil
}
