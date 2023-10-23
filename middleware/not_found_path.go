package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/shoelfikar/kreditplus/model/web"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	// Set a custom 404 response status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	// Write a custom message or page for the 404 response
	response := web.WebResponse{
		Code: http.StatusNotFound,
		Status: "failed",
		Message: http.StatusText(http.StatusNotFound),
	}
	
	json.NewEncoder(w).Encode(response)
}