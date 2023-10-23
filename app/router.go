package app

import (
	"database/sql"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/shoelfikar/kreditplus/controller"
	"github.com/shoelfikar/kreditplus/exception"
	"github.com/shoelfikar/kreditplus/middleware"
	"github.com/shoelfikar/kreditplus/pkg"
	"github.com/shoelfikar/kreditplus/repository"
	"github.com/shoelfikar/kreditplus/service"
)

func NewRouter(db *sql.DB) *mux.Router {

	router := mux.NewRouter()
	validate := validator.New()

	//middleware
	router.Use(exception.Recovery)
	router.Use(middleware.LoggingMiddleware)
	router.NotFoundHandler = http.HandlerFunc(middleware.NotFoundHandler)


	//router version
	v1 := router.PathPrefix("/api/v1").Subrouter()

	//customer
	customerRepo := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepo, validate)
	customerController := controller.NewCustomerController(customerService)

	//tenor
	tenorRepo := repository.NewTenorRepository(db)
	tenorService := service.NewTenorService(tenorRepo, validate)
	tenorController := controller.NewTenorController(tenorService)

	//transaction
	transactionRepo := repository.NewTransactionRepo(db)
	transactionService := service.NewTransactionService(transactionRepo, validate)
	transactionController := controller.NewTransactionController(&transactionService)


	//endpoint customer
	v1.HandleFunc("/customer", customerController.Create).Methods(http.MethodPost)
	v1.HandleFunc("/customer", customerController.FindAll).Methods(http.MethodGet)
	v1.HandleFunc("/customer/{id}", customerController.Update).Methods(http.MethodPut)
	v1.HandleFunc("/customer/{id}", customerController.FindById).Methods(http.MethodGet)

	//endpoint tenor
	v1.HandleFunc("/tenor", tenorController.Create).Methods(http.MethodPost)
	v1.HandleFunc("/tenor", tenorController.Update).Methods(http.MethodPut)
	v1.HandleFunc("/tenor/customer", tenorController.GetTenorByCustomer).Methods(http.MethodPost)

	//endpoint transaction
	v1.HandleFunc("/transaction", transactionController.Create).Methods(http.MethodPost)


	pkg.DefaultLoggingDebug("server running on port " + pkg.GetViperEnvVariable("PORT"))

	return router
}