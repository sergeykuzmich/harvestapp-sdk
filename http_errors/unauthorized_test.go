package errors

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUnauthorizedError(t *testing.T) {
	path := "/401"
	errorMessage := fmt.Sprintf("Unauthorized: %s", path)
	var body []byte

	var err HTTPError
	err = createUnauthorized(path, body)

	var asUnauthorized *Unauthorized
	ok := errors.As(err, &asUnauthorized)
	assert.True(t, ok)

	assert.Equal(t, asUnauthorized.Path(), path)
	assert.Equal(t, asUnauthorized.Details(), body)
	assert.Equal(t, asUnauthorized.Error(), errorMessage)
}

func TestCreateFromUnauthorizedResponse(t *testing.T) {
	path := "/401"
	errorMessage := fmt.Sprintf("Unauthorized: %s", path)

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

	var asUnauthorized *Unauthorized
	ok := errors.As(err, &asUnauthorized)
	assert.True(t, ok)

	assert.Equal(t, asUnauthorized.Path(), path)
	assert.Equal(t, asUnauthorized.Error(), errorMessage)
}
