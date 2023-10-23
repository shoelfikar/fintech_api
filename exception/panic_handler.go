package exception

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/shoelfikar/kreditplus/model/web"
	"github.com/shoelfikar/kreditplus/pkg"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {

				webResponse := errorHandler(err)
				jsonBody, _ := json.Marshal(webResponse)


				pkg.DefaultLoggingWarning(webResponse.Message)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(webResponse.Code)
				w.Write(jsonBody)
			}

		}()

		next.ServeHTTP(w, r)

	})
}

func errorHandler(err interface{}) *web.WebResponse {

	if notFoundError(err) {
		response := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "failed",
			Message: err.(web.NotFoundError).Error,
		}
		return &response
	}

	if validationErrors(err) {
		errors := getErrorMessagevalidator(err)
		response := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "failed",
			Message: "Validation Error",
			Data:    errors,
		}
		return &response
	}

	return internalServerError(err)
}

func getErrorMessagevalidator(err interface{}) web.ValidationError {
	errors := err.(validator.ValidationErrors)

	var messages web.ValidationError

	for _, e := range errors {
		message := web.ValidationMessage{
			Field: e.Field(),
			Tag:   e.Tag(),
			Value: e.Value(),
			Error: fmt.Sprintf("Field %s %s %s", e.Field(), e.Tag(), e.Param()),
		}

		messages.Errors = append(messages.Errors, &message)
	}

	return messages
}

func validationErrors(err interface{}) bool {
	_, ok := err.(validator.ValidationErrors)
	if ok {
		return true
	}

	return false
}

func notFoundError(err interface{}) bool {
	_, ok := err.(web.NotFoundError)
	if ok {
		return true
	}

	return false
}

func internalServerError(err interface{}) *web.WebResponse {

	response := web.WebResponse{
		Code:    http.StatusInternalServerError,
		Status:  "failed",
		Message: err.(error).Error(),
	}
	return &response

}
