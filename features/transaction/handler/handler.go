package handler

import (
	"log"
	"net/http"

	"parkify-BE/features/transaction"
	"parkify-BE/helper"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type controller struct {
	s transaction.TransactionServices
}

func NewHandler(service transaction.TransactionServices) transaction.TransactionController {
	return &controller{
		s: service,
	}
}

func (ct *controller) Transaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		var input transaction.PaymentRequest
		err := c.Bind(&input)
		if err != nil {
			log.Print("error bad request input: ", err.Error())
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirmkan tidak sesuai", nil))
		}

		response, err := ct.s.Transaction(input, token)
		if err != nil {
			log.Println("error bind data:", err.Error())
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
		}

		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "berhasil mendapatkan kode pembayaran", response))
	}
}

func (ct *controller) PaymentCallback() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input transaction.CallbackRequest
		err := c.Bind(&input)
		if err != nil {
			log.Print("error bad request input: ", err.Error())
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirmkan tidak sesuai", nil))
		}

		response := ct.s.PaymentCallback(input)
		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "berhasil mengirim notif", response))
	}
}
