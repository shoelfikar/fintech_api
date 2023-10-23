package repository

import (
	"context"
	"database/sql"
	"math"

	"github.com/shoelfikar/kreditplus/model/domain"
	"github.com/shoelfikar/kreditplus/model/web"
	"github.com/shoelfikar/kreditplus/pkg"
)

type customerRepo struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) domain.CustomerRepository {
	return &customerRepo{
		DB: db,
	}
}

func (repo *customerRepo) Save(ctx context.Context, cust *domain.Customer) *domain.Customer {
	tx, err := repo.DB.Begin()
	pkg.PanicIfError(err)

	
	defer tx.Commit()

	query := "INSERT INTO customers (nik,full_name, legal_name, place_of_birth, date_of_birth, salary, id_card_image, selfie_image, created_by) VALUES (?,?,?,?,?,?,?,?,?) "

	result, errQuery := tx.ExecContext(ctx, query, &cust.NIK, &cust.FullName, &cust.LegalName, &cust.PlaceOfBirth, &cust.DateOfBrith, &cust.Salary, &cust.Idcard_image, &cust.SelfieImage, &cust.CreatedBy)

	if errQuery != nil {
		tx.Rollback()
		pkg.PanicIfError(errQuery)
	}


	lastid, err := result.LastInsertId()

	pkg.PanicIfError(err)

	cust.Id = int(lastid)

	return cust
}

func (repo *customerRepo) Update(ctx context.Context, cust *domain.Customer) *domain.Customer {
	tx, err := repo.DB.Begin()

	pkg.PanicIfError(err)

	defer tx.Commit()

	query := "UPDATE customers SET " 
	query += "nik = ?, "
	query += "full_name = ?, "
	query += "legal_name = ?, "
	query += "place_of_birth = ?, "
	query += "date_of_birth = ?, "
	query += "id_card_image = ?, "
	query += "salary = ?, "
	query += "selfie_image = ?, "
	query += "updated_by = 'admin', "
	query += "updated_at = NOW() "
	query += "WHERE id = ?"

	result, errQuery := tx.ExecContext(ctx, query, &cust.NIK, &cust.FullName, &cust.LegalName, &cust.PlaceOfBirth, &cust.DateOfBrith, &cust.Idcard_image, &cust.Salary, &cust.SelfieImage, &cust.Id)

	if errQuery != nil {
		tx.Rollback()
		pkg.PanicIfError(errQuery)
	}

	rowAffected, err := result.RowsAffected()

	pkg.PanicIfError(err)

	if int(rowAffected) == 0 {
		tx.Rollback()
		panic("no content updated")
	}

	return cust

}

func (repo *customerRepo) FindById(ctx context.Context, custId int) *domain.Customer {
	db := repo.DB
	var (
		cust domain.Customer
		salary float64
	)

	query := "SELECT id,nik,full_name,legal_name,place_of_birth,date_of_birth,id_card_image,salary,selfie_image "
	query += "FROM customers "
	query += "WHERE id = ? "

	errQuery := db.QueryRowContext(ctx, query, &custId).Scan(&cust.Id, &cust.NIK, &cust.FullName, &cust.LegalName, &cust.PlaceOfBirth, &cust.DateOfBrith, &cust.Idcard_image, &salary, &cust.SelfieImage)

	if errQuery != nil {
		if errQuery == sql.ErrNoRows {
			panicErr := web.NotFoundError{Error: "customer not found"}
			panic(panicErr)
		}
		
		pkg.PanicIfError(errQuery)
	}

	cust.Salary = int64(salary)

	return &cust
}

func (repo *customerRepo) FindAll(ctx context.Context, limit int, page int, search string) *domain.CustomerSummary {
	var (
		rows *sql.Rows
		errQuery error
	)
	db := repo.DB
	customers := []*domain.Customer{}
	summary := domain.CustomerSummary{}
	totalData := countCustomers(db, ctx)

	
	query := "SELECT id,nik,full_name,legal_name,place_of_birth,date_of_birth,id_card_image,salary,selfie_image "
	query += "FROM customers "

	switch keyword := search; {
	case keyword != "":
		searchQuery := "WHERE nik LIKE '%?%' OR full_name LIKE '%?%' OR legal_name LIKE'%?%' "

		query += searchQuery
		if page != 0 {
			paging := pkg.Pagination(limit, page, totalData)
			query += paging
		}

		query += "ORDER BY created_at DESC "
	
		rows, errQuery = db.QueryContext(ctx, query, search, search, search)
	default:
		if page != 0 {
			paging := pkg.Pagination(limit, page, totalData)
			query += paging
		}

		query += "ORDER BY created_at DESC "
		rows, errQuery = db.QueryContext(ctx, query)
	}
	

	if errQuery != nil {
		if errQuery == sql.ErrNoRows {
			panicErr := web.NotFoundError{Error: "customers not found"}
			panic(panicErr)
		}
		
		pkg.PanicIfError(errQuery)
	}

	for rows.Next() {
		var salary float64
		cust := domain.Customer{}
		rows.Scan(&cust.Id, &cust.NIK, &cust.FullName, &cust.LegalName, &cust.PlaceOfBirth, &cust.DateOfBrith, &cust.Idcard_image, &salary, &cust.SelfieImage)
		cust.Salary = int64(salary)

		customers = append(customers, &cust)
	}

	if len(customers) == 0 {
		panicErr := web.NotFoundError{Error: "customers not found"}
		panic(panicErr)
	}

	summary.Customers = customers
	summary.Count = totalData
	summary.TotalPage = int(math.Ceil(float64(totalData)/100))

	return &summary

}

func countCustomers(db *sql.DB, ctx context.Context) int {
	cnt := 0
	query := "SELECT COUNT(id) AS cnt FROM customers"

	err := db.QueryRowContext(ctx, query).Scan(&cnt)

	pkg.PanicIfError(err)

	return cnt
}