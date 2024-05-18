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

func TestCreateUnknownError(t *testing.T) {
	status := 418
	path := "/418"
	body := []byte("Unknown")
	errorMessage := fmt.Sprintf("Unknown Error: %d %s %s", status, path, body)

	var err HTTPError
	err = createUnknown(status, path, body)

	var asUnknown *Unknown
	ok := errors.As(err, &asUnknown)
	assert.True(t, ok)

	assert.Equal(t, asUnknown.Status(), status)
	assert.Equal(t, asUnknown.Path(), path)
	assert.Equal(t, asUnknown.Details(), body)
	assert.Equal(t, asUnknown.Error(), errorMessage)
}

func TestCreateFromUnknownResponse(t *testing.T) {

	status := 418
	path := "/418"
	body := "Unknown"
	errorMessage := fmt.Sprintf("Unknown Error: %d %s %s", status, path, body)

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

	var asUnknown *Unknown
	ok := errors.As(err, &asUnknown)
	assert.True(t, ok)

	assert.Equal(t, asUnknown.Status(), status)
	assert.Equal(t, asUnknown.Path(), path)
	assert.Equal(t, asUnknown.Details(), []byte(body))
	assert.Equal(t, asUnknown.Error(), errorMessage)
}
