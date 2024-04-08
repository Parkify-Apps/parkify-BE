package handler

import "github.com/midtrans/midtrans-go/coreapi"

type PaymentResponse struct {
	VirtualAccount []coreapi.VANumber
	TransactionID  string
}
