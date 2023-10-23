package service

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/shoelfikar/kreditplus/model/domain"
	"github.com/shoelfikar/kreditplus/pkg"
)

type customerService struct {
	CustomerRepo domain.CustomerRepository
	Validate *validator.Validate
}

type CustomerService interface {
	Create(ctx context.Context, cust *domain.Customer) *domain.Customer
	Update(ctx context.Context, req []byte, custId int) *domain.Customer
	FindById(ctx context.Context, custId int) *domain.Customer
	FindAll(ctx context.Context, limit string, page string, search string) *domain.CustomerSummary
}

func NewCustomerService(customer domain.CustomerRepository, validator *validator.Validate) CustomerService {
	return &customerService{
		CustomerRepo: customer,
		Validate: validator,
	}
}

func (svc *customerService) Create(ctx context.Context, cust *domain.Customer) *domain.Customer {
	err := svc.Validate.Struct(cust)
	pkg.PanicIfError(err)

	return svc.CustomerRepo.Save(ctx, cust)
}

func (svc *customerService) Update(ctx context.Context, req []byte, custId int) *domain.Customer {
	cust := svc.CustomerRepo.FindById(ctx, custId)

	json.Unmarshal(req, &cust)

	return svc.CustomerRepo.Update(ctx, cust)
}

func (svc *customerService) FindById(ctx context.Context, custId int) *domain.Customer {
	return svc.CustomerRepo.FindById(ctx, custId)
}

func (svc *customerService) FindAll(ctx context.Context, limit string, page string, search string) *domain.CustomerSummary {

	limitpage, _ := strconv.Atoi(limit)
	currentPage, _ := strconv.Atoi(page)
	return svc.CustomerRepo.FindAll(ctx, limitpage, currentPage, search)
}