package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON return a response in JSON for request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json") //converter a resposta para o tipo json
	w.WriteHeader(statusCode)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

// Err return a err in format JSON
func Err(w http.ResponseWriter, statusCode int, err error, field string) {
	JSON(w, statusCode, struct {
		Field string `json:"field,omitempty"`
		Err   string `json:"message"`
	}{
		Err:   err.Error(),
		Field: field,
	})
}

func Sucess(w http.ResponseWriter, statusCode int, message string) {
	JSON(w, statusCode, struct {
		Message string `json:"message,omitempty"`
	}{
		Message: message,
	})
}
