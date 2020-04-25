package http_errors

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUnexpectedError(t *testing.T) {
	status := 418
	path := "/ping"
	body := []byte("Unexpected")
	errorMessage := fmt.Sprintf("Unexpected Error: %d %s %s", status, path, body)


	var err HttpError
	err = createUnexpected(status, path, body)

	converted_err, ok := err.(*Unexpected)
	assert.True(t, ok)

	assert.Equal(t, converted_err.Status(), status)
	assert.Equal(t, converted_err.Path(), path)
	assert.Equal(t, converted_err.Details(), body)
	assert.Equal(t, converted_err.Error(), errorMessage)
}

func TestCreateFromUnexpectedResponse(t *testing.T) {
	req, _ := http.NewRequest("GET", "/unexpected", nil)
	res := &http.Response{
	  Status:        "418 I'm a teapot",
	  StatusCode:    418,
	  Proto:         "HTTP/1.1",
	  Body:          ioutil.NopCloser(bytes.NewBufferString("")),
	  ContentLength: int64(len("")),
	  Request:       req,
	}

	err := CreateFromResponse(res)

	err, ok := err.(*Unexpected)
	assert.True(t, ok)
}
