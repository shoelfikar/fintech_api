package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/shoelfikar/kreditplus/model/domain"
	"github.com/shoelfikar/kreditplus/pkg"
)

type transactionService struct{
	transacationRepo domain.TransactionRepository
	Validate *validator.Validate

}

type TransactionService interface {
	Create(ctx context.Context,transaction *domain.Transaction) *domain.Transaction
}

func NewTransactionService(trx domain.TransactionRepository, validator *validator.Validate) transactionService {
	return transactionService{
		transacationRepo: trx,
		Validate: validator,
	}
}

func (svc *transactionService) Create(ctx context.Context,transaction *domain.Transaction) *domain.Transaction {
	err := svc.Validate.Struct(transaction)
	
	pkg.PanicIfError(err)

	return svc.transacationRepo.Save(ctx, transaction)
}