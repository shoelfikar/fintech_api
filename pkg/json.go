package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/shoelfikar/kreditplus/model/web"
)

func ReadFromRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	code := response.(web.WebResponse).Code
	w.WriteHeader(code)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfError(err)
}