package utils

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type mdtrans struct {
	c coreapi.Client
}

// type PaymentVAResponse struct {
// 	Status        []coreapi.VANumber
// 	TransactionID string
// }

type PaymentFunc interface {
	PaymentVABCA(idBook string, price int) (*coreapi.ChargeResponse, error)
}

func NewMidtrans(sk string) PaymentFunc {
	c := coreapi.Client{}
	c.New(sk, midtrans.Sandbox)
	return &mdtrans{
		c: c,
	}
}

func (m *mdtrans) PaymentVABCA(idBook string, price int) (*coreapi.ChargeResponse, error) {
	midtrans.ServerKey = m.c.ServerKey
	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  idBook,
			GrossAmt: int64(price),
		}, BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.BankBca,
			// VaNumber: "1111111",
			// Bca: &coreapi.BcaBankTransferDetail{
			// 	SubCompanyCode: "0000",
			// },
		},
	}

	chargeResp, err := coreapi.ChargeTransaction(chargeReq)
	if err != nil {
		return &coreapi.ChargeResponse{}, err
	}

	// // Create payment response
	// response := PaymentVAResponse{
	// 	Status:        chargeResp.VaNumbers,
	// 	TransactionID: chargeResp.TransactionID,
	// 	// Add other fields as needed
	// }

	return chargeResp, nil
}
