package domain

import "context"

type TransactionRepository interface {
	Save(ctx context.Context, trx *Transaction) *Transaction
}

type Transaction struct {
	Id             int    `json:"id,omitempty"`
	ContractNumber string `json:"contract_number" validate:"required"`
	OTR            int64  `json:"otr" validate:"required"`
	CustomerId     int    `json:"customer_id" validate:"required"`
	AdminFee       int64  `json:"admin_fee" validate:"required"`
	InterestFee    int    `json:"interest_fee"`
	Installments   int    `json:"installments" validate:"required"`
	AssetName      string `json:"asset_name" validate:"required,min=5"`
}
