package utils

import (
	"encoding/json"
	"net/http"
)

// JSONResponse - Sends JSON Response in the responseWriter Stream
func JSONResponse(w http.ResponseWriter, response interface{}, statusCode int) {
	msg, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application-json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(msg)
}
