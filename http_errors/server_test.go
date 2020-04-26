package http_errors

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateServerError(t *testing.T) {
	status := 500
	path := "/500"
	body := []byte("Internal Server Error")
	errorMessage := fmt.Sprintf("%d %s %s", status, path, body)

	var err *Server
	err = createServerError(status, path, body)

	assert.Equal(t, err.Status(), status)
	assert.Equal(t, err.Path(), path)
	assert.Equal(t, err.Details(), body)
	assert.Equal(t, err.Error(), errorMessage)
}

func TestCreateFromServerErrordResponse(t *testing.T) {
	status := 500
	path := "/500"
	body := "Internal Server Error"
	errorMessage := fmt.Sprintf("%d %s %s", status, path, body)

	req, _ := http.NewRequest("GET", path, nil)
	res := &http.Response{
		Status:        "500 Internal Server Error",
		StatusCode:    status,
		Proto:         "HTTP/1.1",
		Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
		ContentLength: int64(len("")),
		Request:       req,
	}

	err := CreateFromResponse(res)

	asServer, ok := err.(*Server)
	assert.True(t, ok)

	assert.Equal(t, asServer.Status(), status)
	assert.Equal(t, asServer.Path(), path)
	assert.Equal(t, asServer.Details(), []byte(body))
	assert.Equal(t, asServer.Error(), errorMessage)
}
