package http_errors

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNotFoundError(t *testing.T) {
	path := "/ping"
	errorMessage := fmt.Sprintf("Not Found: %s", path)
	var body []byte


	var err HttpError
	err = createNotFound(path, body)

	err, ok := err.(*NotFound)
	assert.True(t, ok)

	assert.Equal(t, err.Path(), path)
	assert.Equal(t, err.Details(), body)
	assert.Equal(t, err.Error(), errorMessage)
}

func TestCreateFromNotFoundResponse(t *testing.T) {
	req, _ := http.NewRequest("GET", "/tasks", nil)
	res := &http.Response{
	  Status:        "404 Not Found",
	  StatusCode:    404,
	  Proto:         "HTTP/1.1",
	  Body:          ioutil.NopCloser(bytes.NewBufferString("")),
	  ContentLength: int64(len("")),
	  Request:       req,
	}

	err := CreateFromResponse(res)

	err, ok := err.(*NotFound)
	assert.True(t, ok)
}
