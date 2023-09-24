package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aerostatas/interaction-go-kata/internal/service"
	"github.com/stretchr/testify/assert"
)

func Test_ResponseJSON(t *testing.T) {
	writer := httptest.NewRecorder()
	data := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Success: true,
		Message: "test message",
	}
	code := http.StatusOK

	ResponseJSON(writer, data, code)

	assert.Equal(t, code, writer.Code)
	assert.Equal(t, "{\"success\":true,\"message\":\"test message\"}\n", writer.Body.String())
}

func Test_ErrorJSON_ErrInvalidValue(t *testing.T) {
	writer := httptest.NewRecorder()
	err := fmt.Errorf("test error: %w", service.ErrInvalidValue)
	code := http.StatusBadRequest

	ErrorJSON(writer, err)

	assert.Equal(t, code, writer.Code)
	assert.Equal(t, "{\"code\":400,\"error\":\"test error: invalid value\"}\n", writer.Body.String())
}

func Test_ErrorJSON_ErrInvalidAmount(t *testing.T) {
	writer := httptest.NewRecorder()
	err := fmt.Errorf("test error: %w", service.ErrInvalidAmount)
	code := http.StatusBadRequest

	ErrorJSON(writer, err)

	assert.Equal(t, code, writer.Code)
	assert.Equal(t, "{\"code\":400,\"error\":\"test error: invalid amount\"}\n", writer.Body.String())
}

func Test_ErrorJSON_ErrInvalidJSON(t *testing.T) {
	writer := httptest.NewRecorder()
	err := fmt.Errorf("test error: %w", ErrInvalidJSON)
	code := http.StatusBadRequest

	ErrorJSON(writer, err)

	assert.Equal(t, code, writer.Code)
	assert.Equal(t, "{\"code\":400,\"error\":\"test error: invalid JSON\"}\n", writer.Body.String())
}

func Test_ErrorJSON_Unknown(t *testing.T) {
	writer := httptest.NewRecorder()
	err := fmt.Errorf("test error")
	code := http.StatusInternalServerError

	ErrorJSON(writer, err)

	assert.Equal(t, code, writer.Code)
	assert.Equal(t, "{\"code\":500,\"error\":\"test error\"}\n", writer.Body.String())
}
