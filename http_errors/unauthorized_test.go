package http_errors

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUnauthorizedError(t *testing.T) {
	path := "/tasks/401"
	errorMessage := fmt.Sprintf("Unauthorized: %s", path)
	var body []byte

	var err *Unauthorized
	err = createUnauthorized(path, body)

	assert.Equal(t, err.Path(), path)
	assert.Equal(t, err.Details(), body)
	assert.Equal(t, err.Error(), errorMessage)
}

func TestCreateFromUnauthorizedResponse(t *testing.T) {
	path := "/tasks/401"
	req, _ := http.NewRequest("GET", path, nil)
	res := &http.Response{
		Status:        "401 Unauthorized",
		StatusCode:    401,
		Proto:         "HTTP/1.1",
		Body:          ioutil.NopCloser(bytes.NewBufferString("")),
		ContentLength: int64(len("")),
		Request:       req,
	}

	err := CreateFromResponse(res)

	asUnauthorized, ok := err.(*Unauthorized)
	assert.True(t, ok)

	assert.Equal(t, asUnauthorized.Path(), path)
	assert.Equal(t, asUnauthorized.Error(), fmt.Sprintf("Unauthorized: %s", path))
}
