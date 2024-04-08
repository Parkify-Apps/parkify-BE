package handler

type PaymentRequest struct {
	ReservationID uint
	PaymentMethod string
}

type CallbackRequest struct {
	OrderID           string `json:"order_id"`
	GrossAmount       string `json:"gross_amount"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
	TransactionStatus string `json:"transaction_status"`
}
