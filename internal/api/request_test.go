package api

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseJSON_rejectsInvalidJson(t *testing.T) {
	r := httptest.NewRequest("POST", "/api/test", strings.NewReader("test"))

	var s struct {
		Test bool `json:"test"`
	}

	err := ParseJSON(r, &s)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "decode request to JSON")
}

func Test_ParseJSON(t *testing.T) {
	r := httptest.NewRequest("POST", "/api/test", strings.NewReader("{\"test\":true}"))

	var s struct {
		Test bool `json:"test"`
	}

	err := ParseJSON(r, &s)

	assert.NoError(t, err)
	assert.True(t, s.Test)
}
