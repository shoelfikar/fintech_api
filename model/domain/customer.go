package domain

import "context"

type CustomerRepository interface {
	Save(ctx context.Context, cust *Customer) *Customer
	Update(ctx context.Context, cust *Customer) *Customer
	FindById(ctx context.Context, custId int) *Customer
	FindAll(ctx context.Context, limit int, page int, search string) *CustomerSummary
}

type Customer struct {
	Id           int    `json:"id,omitempty"`
	NIK          string `json:"nik" validate:"required,min=16"`
	FullName     string `json:"full_name,omitempty"`
	LegalName    string `json:"legal_name,omitempty"`
	PlaceOfBirth string `json:"place_of_birth,omitempty"`
	DateOfBrith  string `json:"date_of_birth,omitempty"`
	Salary       int64  `json:"salary,omitempty"`
	Idcard_image string `json:"id_card_image,omitempty"`
	SelfieImage  string `json:"selfie_image,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	CreatedBy    string `json:"created_by,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	UpdatedBy    string `json:"updated_by,omitempty"`
}

type CustomerUpdateRequest struct {
	Id           int    `json:"id,omitempty"`
	NIK          string `json:"nik,omitempty" validate:"min=16"`
	FullName     string `json:"full_name,omitempty" validate:"min=5"`
	LegalName    string `json:"legal_name,omitempty" validate:"min=5"`
	PlaceOfBirth string `json:"place_of_birth,omitempty" validate:"min=5"`
	DateOfBrith  string `json:"date_of_birth,omitempty" validate:"min=5"`
	Salary       int64  `json:"salary,omitempty" validate:"min=100000"`
	Idcard_image string `json:"id_card_image,omitempty" validate:"min=5"`
	SelfieImage  string `json:"selfie_image,omitempty" validate:"min=5"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	UpdatedBy    string `json:"updated_by,omitempty"`
}

type CustomerSummary struct {
	Count     int         `json:"count,omitempty"`
	Customers []*Customer `json:"customers,omitempty"`
	TotalPage int         `json:"total_page,omitempty"`
}

