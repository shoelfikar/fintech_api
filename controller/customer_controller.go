package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shoelfikar/kreditplus/model/domain"
	"github.com/shoelfikar/kreditplus/model/web"
	"github.com/shoelfikar/kreditplus/pkg"
	"github.com/shoelfikar/kreditplus/service"
)

type customerController struct {
	CustomerService service.CustomerService
}

type CustomerController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
}

func NewCustomerController(customerService service.CustomerService) CustomerController {
	return &customerController{
		CustomerService: customerService,
	}
}

func (ctl *customerController) Create(w http.ResponseWriter, r *http.Request) {
	var req *domain.Customer

	pkg.ReadFromRequestBody(r, &req)


	req = ctl.CustomerService.Create(r.Context(), req)

	response := web.WebResponse{
		Code: http.StatusCreated,
		Status: "success",
		Message: http.StatusText(http.StatusCreated),
		Data: req,
	}

	pkg.WriteToResponseBody(w, response)
}

func (ctl *customerController) Update(w http.ResponseWriter, r *http.Request) {
	var req domain.CustomerUpdateRequest

	params := mux.Vars(r)
	paramId := params["id"]

	custId, err := strconv.Atoi(paramId)

	pkg.PanicIfError(err)

	pkg.ReadFromRequestBody(r, &req)

	reqByte, err := json.Marshal(req)

	pkg.PanicIfError(err)

	customer := ctl.CustomerService.Update(r.Context(), reqByte, custId)

	response := web.WebResponse{
		Code: http.StatusOK,
		Status: "success",
		Message: http.StatusText(http.StatusOK),
		Data: customer,
	}

	pkg.WriteToResponseBody(w, response)
}

func (ctl *customerController) FindById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramId := params["id"]

	custId, err := strconv.Atoi(paramId)

	pkg.PanicIfError(err)

	customer := ctl.CustomerService.FindById(r.Context(), custId)

	response := web.WebResponse{
		Code: http.StatusOK,
		Status: "success",
		Message: http.StatusText(http.StatusOK),
		Data: customer,
	}

	pkg.WriteToResponseBody(w, response)
}

func (ctl *customerController) FindAll(w http.ResponseWriter, r *http.Request) {
	
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	search := r.URL.Query().Get("search")
	
	customers := ctl.CustomerService.FindAll(r.Context(), limit, page, search)

	response := web.WebResponse{
		Code: http.StatusOK,
		Status: "success",
		Message: http.StatusText(http.StatusOK),
		Data: customers,
	}

	pkg.WriteToResponseBody(w, response)
}