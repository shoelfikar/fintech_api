package controller

import (
	"net/http"

	"github.com/shoelfikar/kreditplus/model/domain"
	"github.com/shoelfikar/kreditplus/model/web"
	"github.com/shoelfikar/kreditplus/pkg"
	"github.com/shoelfikar/kreditplus/service"
)

type transactionController struct {
	transactionService service.TransactionService
}

type TransactionController interface {
	Create(w http.ResponseWriter, r *http.Request)
}

func NewTransactionController(trxService service.TransactionService) TransactionController {
	return &transactionController{
		transactionService: trxService,
	}
}

func (ctl *transactionController) Create(w http.ResponseWriter, r *http.Request) {
	var req *domain.Transaction

	pkg.ReadFromRequestBody(r, &req)

	req = ctl.transactionService.Create(r.Context(), req)

	response := web.WebResponse{
		Code: http.StatusCreated,
		Status: "success",
		Message: http.StatusText(http.StatusCreated),
		Data: req,
	}

	pkg.WriteToResponseBody(w, response)
}