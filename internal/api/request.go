package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ParseJSON attempts to decode request body to the provided var
func ParseJSON(r *http.Request, v any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&v); err != nil {
		return fmt.Errorf("decode request to JSON: %w", err)
	}

	return nil
}
