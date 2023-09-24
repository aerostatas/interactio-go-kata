package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/aerostatas/interaction-go-kata/internal/service"
)

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

// ResponseJSON encodes data to JSON and writes it as a response
// with the appropriate Content-Type headers
func ResponseJSON(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// ErrorJSON structures an error as JSON and writes it as a response
func ErrorJSON(w http.ResponseWriter, err error) {
	statusCode := errToStatusCode(err)
	errData := ErrorResponse{statusCode, err.Error()}
	ResponseJSON(w, errData, statusCode)
}

func errToStatusCode(err error) int {
	mapped := map[error]int{
		ErrInvalidJSON:           http.StatusBadRequest,
		service.ErrInvalidAmount: http.StatusBadRequest,
		service.ErrInvalidValue:  http.StatusBadRequest,
	}

	for merr, code := range mapped {
		if errors.Is(err, merr) {
			return code
		}
	}

	return http.StatusInternalServerError
}
