package repository

import (
	"context"
	"database/sql"

	"github.com/shoelfikar/kreditplus/model/domain"
	"github.com/shoelfikar/kreditplus/pkg"
)

type transaactionRepo struct {
	DB *sql.DB
}

func NewTransactionRepo(db *sql.DB) domain.TransactionRepository {
	return &transaactionRepo{
		DB: db,
	}
}


func (repo *transaactionRepo) Save(ctx context.Context, trx *domain.Transaction) *domain.Transaction {
	tx, err := repo.DB.Begin()

	pkg.PanicIfError(err)

	query := "INSERT INTO transaction (customer_id, contract_number, otr, admin_fee, interest_fee, installments, asset_name, created_by) "
	query += "VALUES (?,?,?,?,?,?,?,?)"

	result, errQuery := tx.ExecContext(ctx, query, &trx.CustomerId, &trx.ContractNumber, &trx.OTR , &trx.AdminFee, &trx.InterestFee, &trx.Installments, &trx.AssetName, "admin")

	pkg.PanicIfError(errQuery)

	lastId, err := result.LastInsertId()
	pkg.PanicIfError(err)

	trx.Id = int(lastId)
	return trx
}