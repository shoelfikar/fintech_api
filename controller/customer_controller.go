package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
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
	GetFile(w http.ResponseWriter, r *http.Request)
}

func NewCustomerController(customerService service.CustomerService) CustomerController {
	return &customerController{
		CustomerService: customerService,
	}
}

func (ctl *customerController) Create(w http.ResponseWriter, r *http.Request) {
	var req *domain.Customer

	body := r.FormValue("body_json")
	name := r.ParseForm()

	// bodyByt, _ := json.Marshal(body)

	json.Unmarshal([]byte(body), &req)

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, "Failed to get the file", http.StatusBadRequest)
	}

	file, handler, err := r.FormFile("ktp")
	if err != nil {
		http.Error(w, "Failed to get the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	path := "./public/ktp/"

	// Check if the destination folder exists, and if not, create it
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			http.Error(w, "Failed to create the destination folder", http.StatusInternalServerError)
			return
		}
	}

	// Create a new file on the server to save the uploaded file
	outputFile, err := os.Create(filepath.Join(path, handler.Filename))
	if err != nil {
		http.Error(w, "Failed to create the file on the server", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// Copy the uploaded file to the server file
	_, err = io.Copy(outputFile, file)
	if err != nil {
		http.Error(w, "Failed to copy the file", http.StatusInternalServerError)
		return
	}

	// req = ctl.CustomerService.Create(r.Context(), req)

	response := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "success",
		Message: http.StatusText(http.StatusCreated),
		Data:    req,
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
		Code:    http.StatusOK,
		Status:  "success",
		Message: http.StatusText(http.StatusOK),
		Data:    customer,
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
		Code:    http.StatusOK,
		Status:  "success",
		Message: http.StatusText(http.StatusOK),
		Data:    customer,
	}

	pkg.WriteToResponseBody(w, response)
}

func (ctl *customerController) FindAll(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	search := r.URL.Query().Get("search")

	customers := ctl.CustomerService.FindAll(r.Context(), limit, page, search)

	response := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: http.StatusText(http.StatusOK),
		Data:    customers,
	}

	pkg.WriteToResponseBody(w, response)
}

func (ctl *customerController) GetFile(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("./public/ktp/", "maki_zenin.jpg")
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer file.Close()

	// Read the file's content
	stat, _ := file.Stat()
	fileSize := stat.Size()
	fileContent := make([]byte, fileSize)
	_, err = file.Read(fileContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header based on the file type
	contentType := http.DetectContentType(fileContent)
	w.Header().Set("Content-Type", contentType)

	// Write the file content to the response
	w.Write(fileContent)
}
