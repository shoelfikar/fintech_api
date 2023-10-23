package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/shoelfikar/kreditplus/model/domain"
	"github.com/shoelfikar/kreditplus/model/web"
	"github.com/shoelfikar/kreditplus/pkg"
)

type tenorRepo struct{
	DB *sql.DB
}

func NewTenorRepository(db *sql.DB) domain.TenorReposiroty {
	return &tenorRepo{
		DB: db,
	}
}

func (repo *tenorRepo) Save(ctx context.Context, tnr *domain.LimitTenor) *domain.LimitTenor {
	tx, err := repo.DB.Begin()
	pkg.PanicIfError(err)

	defer tx.Commit()

	query := "INSERT INTO tenor (customer_id, customer_limit, tenor, created_by) "
	query += "VALUES (?,?,?,?) "

	result, errQuery := tx.ExecContext(ctx, query, &tnr.CustomerId, &tnr.Limit, &tnr.Tenor, "admin")

	if errQuery != nil {
		tx.Rollback()
		pkg.PanicIfError(errQuery)
	}

	lastId, err := result.LastInsertId()

	pkg.PanicIfError(err)

	tnr.Id = int(lastId)

	return tnr

}

func (repo *tenorRepo) Update(ctx context.Context, tnr *domain.TenorUpdateRequest) *domain.TenorUpdateRequest {
	tx, err := repo.DB.Begin()

	pkg.PanicIfError(err)

	defer tx.Commit()

	query := "UPDATE tenor SET customer_limit = ?, "
	query += "tenor = ? "
	query += "WHERE customer_id = ? AND id = ?"

	result, errQuery := tx.ExecContext(ctx, query,&tnr.Limit, &tnr.Tenor, &tnr.CustomerId, &tnr.Id)

	if errQuery != nil {
		tx.Rollback()
		pkg.PanicIfError(errQuery)
	}

	rowAffected, err := result.RowsAffected()
	
	if rowAffected == 0 {
		noUpdated := errors.New("no data updated")
		pkg.PanicIfError(noUpdated)
	}

	return tnr
}

func (repo *tenorRepo) FindById(ctx context.Context, custId int, tenorId int) *domain.LimitTenor {
	var tr *domain.LimitTenor
	db := repo.DB

	query := "SELECT id,customer_id,customer_limit,tenor FROM tenor"
	query += "WHERE id = ? AND customer_id = ?"

	errQuery := db.QueryRowContext(ctx, query, tenorId, custId).Scan(&tr.Id, &tr.CustomerId, &tr.Limit, &tr.Tenor)

	if errQuery != nil {
		if errQuery == sql.ErrNoRows {
			panicErr := web.NotFoundError{Error: "tenor not found"}
			panic(panicErr)
		}
		
		pkg.PanicIfError(errQuery)
	}

	return tr
}

func (repo *tenorRepo) GetTenorByCustomer(ctx context.Context, tnr *domain.TenorRequest) []*domain.LimitTenor {
	var t []*domain.LimitTenor
	db := repo.DB

	query := "SELECT id,customer_id, customer_limit, tenor, created_at, updated_at FROM tenor "
	query += "WHERE customer_id = ? AND tenor = ?"

	rows, errQuery := db.QueryContext(ctx, query, &tnr.CustomerId, &tnr.Tenor)

	if errQuery != nil {
		if errQuery == sql.ErrNoRows {
			panicErr := web.NotFoundError{Error: "customer not found"}
			panic(panicErr)
		}
		
		pkg.PanicIfError(errQuery)
	}

	for rows.Next() {
		var limit float64
		ct := domain.LimitTenor{}

		rows.Scan(&ct.Id, &ct.CustomerId, &limit, &ct.Tenor, &ct.CreatedAt, &ct.UpdatedAt)

		ct.Limit = int64(limit)

		t = append(t, &ct)
	}

	return t
}