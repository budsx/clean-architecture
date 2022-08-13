package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequest(request *http.Request, result	 interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result	)
	NewPanicError(err)
}

func WriteResponse(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)

	err := encoder.Encode(response)
	NewPanicError(err)
}
