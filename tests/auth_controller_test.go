package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	parameters := []byte(`
	{
		"email":"jeanmichel.lecul@gmail.com",
		"password":"strongPassword"
	}`)

	resp := SendRequest(parameters, "POST", "/v1/auth/")
	assert.Equal(t, http.StatusOK, resp.Code)
}
