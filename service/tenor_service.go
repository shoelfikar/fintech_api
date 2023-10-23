package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/shoelfikar/kreditplus/model/domain"
	"github.com/shoelfikar/kreditplus/pkg"
)

type tenorService struct {
	tenorRepo domain.TenorReposiroty
	Validate *validator.Validate
}

type TenorService interface {
	Create(ctx context.Context, tnr *domain.LimitTenor) *domain.LimitTenor
	Update(ctx context.Context, tnr *domain.TenorUpdateRequest) *domain.TenorUpdateRequest
	GetTenorByCustomer(ctx context.Context, tnr *domain.TenorRequest) []*domain.LimitTenor
}

func NewTenorService(repo domain.TenorReposiroty, validator *validator.Validate) TenorService {
	return &tenorService{
		tenorRepo: repo,
		Validate: validator,
	}
}

func (svc *tenorService) Create(ctx context.Context, tnr *domain.LimitTenor) *domain.LimitTenor {
	err := svc.Validate.Struct(tnr)
	pkg.PanicIfError(err)
	return svc.tenorRepo.Save(ctx, tnr)
}

func (svc *tenorService) Update(ctx context.Context, tnr *domain.TenorUpdateRequest) *domain.TenorUpdateRequest {
	err := svc.Validate.Struct(tnr)
	pkg.PanicIfError(err)
	return svc.tenorRepo.Update(ctx, tnr)
}

func (svc *tenorService) GetTenorByCustomer(ctx context.Context, tnr *domain.TenorRequest) []*domain.LimitTenor {
	err := svc.Validate.Struct(tnr)
	pkg.PanicIfError(err)
	return svc.tenorRepo.GetTenorByCustomer(ctx, tnr)
}