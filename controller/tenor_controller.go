package controller

import (
	"net/http"

	"github.com/shoelfikar/kreditplus/model/domain"
	"github.com/shoelfikar/kreditplus/model/web"
	"github.com/shoelfikar/kreditplus/pkg"
	"github.com/shoelfikar/kreditplus/service"
)

type tenorController struct {
	tenorService service.TenorService
}

type TenorController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	GetTenorByCustomer(w http.ResponseWriter, r *http.Request)
}

func NewTenorController(tenorService service.TenorService) TenorController {
	return &tenorController{
		tenorService: tenorService,
	}
}

func (ctl *tenorController) Create(w http.ResponseWriter, r *http.Request) {
	var req *domain.LimitTenor
	
	pkg.ReadFromRequestBody(r, &req)

	req = ctl.tenorService.Create(r.Context(), req)

	response := web.WebResponse{
		Code: http.StatusCreated,
		Status: "success",
		Message: http.StatusText(http.StatusCreated),
		Data: req,
	}

	pkg.WriteToResponseBody(w, response)
}

func (ctl *tenorController) Update(w http.ResponseWriter, r *http.Request) {
	var req *domain.TenorUpdateRequest

	pkg.ReadFromRequestBody(r, &req)

	req = ctl.tenorService.Update(r.Context(), req)

	response := web.WebResponse{
		Code: http.StatusCreated,
		Status: "success",
		Message: http.StatusText(http.StatusCreated),
		Data: req,
	}

	pkg.WriteToResponseBody(w, response)
}

func (ctl *tenorController) GetTenorByCustomer(w http.ResponseWriter, r *http.Request) {
	var req *domain.TenorRequest

	pkg.ReadFromRequestBody(r, &req)

	custtenor := ctl.tenorService.GetTenorByCustomer(r.Context(), req)

	response := web.WebResponse{
		Code: http.StatusCreated,
		Status: "success",
		Message: http.StatusText(http.StatusCreated),
		Data: custtenor,
	}

	pkg.WriteToResponseBody(w, response)
}