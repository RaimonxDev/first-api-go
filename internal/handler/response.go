package handler

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

func NewResponse(message string, status int, err interface{}, data interface{}) response {
	//return struct response
	return response{
		Message: message,
		Status:  status,
		Error:   err,
		Data:    data,
	}
}

func (r *response) ToJSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	err := json.NewEncoder(w).Encode(&r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return nil
}
