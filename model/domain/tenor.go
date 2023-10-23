package domain

import "context"

type TenorReposiroty interface {
	Save(ctx context.Context, tnr *LimitTenor) *LimitTenor
	Update(ctx context.Context, tnr *TenorUpdateRequest) *TenorUpdateRequest
	GetTenorByCustomer(ctx context.Context, tnr *TenorRequest) []*LimitTenor
	FindById(ctx context.Context, custId int, tenorId int) *LimitTenor
}

type LimitTenor struct {
	Id         int    `json:"id"`
	CustomerId int    `json:"customer_id" validate:"required"`
	Limit      int64  `json:"limit" validate:"required,min=100000"`
	Tenor      int    `json:"tenor" validate:"required,min=1"`
	CreatedAt  string `json:"created_at,omitempty"`
	CreatedBy  string `json:"created_by,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	UpdatedBy  string `json:"updated_by,omitempty"`
}

type TenorUpdateRequest struct {
	Id         int   `json:"id" validate:"required"`
	CustomerId int   `json:"customer_id" validate:"required"`
	Limit      int64 `json:"limit"`
	Tenor      int32 `json:"tenor"`
}

type TenorRequest struct {
	CustomerId int   `json:"customer_id"`
	Tenor      int32 `json:"tenor"`
}
