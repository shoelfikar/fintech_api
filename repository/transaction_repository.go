package repository

import (
	"context"
	"database/sql"

	"github.com/shoelfikar/kreditplus/model/domain"
	"github.com/shoelfikar/kreditplus/model/web"
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

func (repo *transaactionRepo) GetTransactionByCustomer(ctx context.Context, custId int, tenor int) []*domain.Transaction {
	var trx []*domain.Transaction
	db := repo.DB

	query := "SELECT id,contract_number,customer_id,otr,admin_fee,interest_fee,installments,asset_name,created_at,created_by FROM transaction "
	query += "WHERE customer_id = ? AND installments = ? AND trx_status = 'active'"

	rows, errQuery := db.QueryContext(ctx, query, &custId, &tenor)
	
	if errQuery != nil {
		if errQuery == sql.ErrNoRows {
			panicErr := web.NotFoundError{Error: "customer not found"}
			panic(panicErr)
		}
		
		pkg.PanicIfError(errQuery)
	}

	for rows.Next() {
		t := domain.Transaction{}

		rows.Scan(&t.Id, &t.ContractNumber, &t.CustomerId, &t.OTR, &t.AdminFee, &t.InterestFee, &t.Installments, &t.AssetName, &t.CreatedAt, &t.CreatedBy)

		trx = append(trx, &t)
	}

	return trx

}