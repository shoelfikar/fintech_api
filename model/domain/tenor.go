package domain

import "context"

type TenorReposiroty interface {
	Save(ctx context.Context, tnr *LimitTenor) LimitTenor
	Update(ctx context.Context, tnr *LimitTenor) LimitTenor
}

type LimitTenor struct {
	Id          int    `json:"id"`
	Limit       int64  `json:"limit"`
	AdminFee    int64  `json:"admin_fee"`
	InterestFee int64  `json:"interest_fee"`
	Tenor       int32  `json:"tenor"`
	Status      int    `json:"status"`
	MinSalary   int64  `json:"min_salary"`
	CreatedAt   string `json:"created_at,omitempty"`
	CreatedBy   string `json:"created_by,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	UpdatedBy   string `json:"updated_by,omitempty"`
}

type TenorUpdateRequest struct {
	Id         int   `json:"id"`
	CustomerId int   `json:"customer_id"`
	Limit      int64 `json:"limit"`
	Tenor      int32 `json:"tenor"`
}
