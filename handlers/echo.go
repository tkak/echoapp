package handlers

import (
	"encoding/json"
	"net/http"
)

type EchoResponse struct {
	Message string `json:"message"`
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	response := EchoResponse{
		Message: q.Get("message"),
	}
	json.NewEncoder(w).Encode(response)
	return
}
