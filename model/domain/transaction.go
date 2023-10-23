package domain

import "context"

type TransactionRepository interface {
	Save(ctx context.Context, trx *Transaction) *Transaction
	GetTransactionByCustomer(ctx context.Context, custId int, tenor int) []*Transaction
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
	TrxStatus      string `json:"trx_status" validate:"oneof=finish active"`
	CreatedAt      string `json:"created_at,omitempty"`
	CreatedBy      string `json:"created_by,omitempty"`
}
