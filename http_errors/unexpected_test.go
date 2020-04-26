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
	path := "/418"
	body := []byte("Unexpected")
	errorMessage := fmt.Sprintf("Unexpected Error: %d %s %s", status, path, body)

	var err *Unexpected
	err = createUnexpected(status, path, body)

	assert.Equal(t, err.Status(), status)
	assert.Equal(t, err.Path(), path)
	assert.Equal(t, err.Details(), body)
	assert.Equal(t, err.Error(), errorMessage)
}

func TestCreateFromUnexpectedResponse(t *testing.T) {

	status := 418
	path := "/418"
	body := "Unexpected"
	errorMessage := fmt.Sprintf("Unexpected Error: %d %s %s", status, path, body)

	req, _ := http.NewRequest("GET", path, nil)
	res := &http.Response{
		Status:        "418 I'm a teapot",
		StatusCode:    status,
		Proto:         "HTTP/1.1",
		Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
		ContentLength: int64(len("")),
		Request:       req,
	}

	err := CreateFromResponse(res)

	asUnexpected, ok := err.(*Unexpected)
	assert.True(t, ok)

	assert.Equal(t, asUnexpected.Status(), status)
	assert.Equal(t, asUnexpected.Path(), path)
	assert.Equal(t, asUnexpected.Details(), []byte(body))
	assert.Equal(t, asUnexpected.Error(), errorMessage)
}
