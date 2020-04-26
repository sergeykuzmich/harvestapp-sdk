package http_errors

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTooManyRequestsError(t *testing.T) {
	path := "/429"
	errorMessage := fmt.Sprintf("Too Many Requests: %s", path)
	var body []byte

	var err *TooManyRequests
	err = createTooManyRequests(path, body)

	assert.Equal(t, err.Path(), path)
	assert.Equal(t, err.Details(), body)
	assert.Equal(t, err.Error(), errorMessage)
}

func TestCreateFromTooManyRequestsResponse(t *testing.T) {
	path := "/429"
	req, _ := http.NewRequest("GET", path, nil)
	res := &http.Response{
		Status:        "429 Too Many Requests",
		StatusCode:    429,
		Proto:         "HTTP/1.1",
		Body:          ioutil.NopCloser(bytes.NewBufferString("")),
		ContentLength: int64(len("")),
		Request:       req,
	}

	err := CreateFromResponse(res)

	asTooManyRequests, ok := err.(*TooManyRequests)
	assert.True(t, ok)

	assert.Equal(t, asTooManyRequests.Path(), path)
}
