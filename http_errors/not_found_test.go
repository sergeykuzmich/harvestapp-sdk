package errors

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNotFoundError(t *testing.T) {
	path := "/404"
	errorMessage := fmt.Sprintf("Not Found: %s", path)
	var body []byte

	var err *NotFound
	err = createNotFound(path, body)

	assert.Equal(t, err.Path(), path)
	assert.Equal(t, err.Details(), body)
	assert.Equal(t, err.Error(), errorMessage)
}

func TestCreateFromNotFoundResponse(t *testing.T) {
	path := "/404"
	req, _ := http.NewRequest("GET", path, nil)
	res := &http.Response{
		Status:        "404 Not Found",
		StatusCode:    404,
		Proto:         "HTTP/1.1",
		Body:          ioutil.NopCloser(bytes.NewBufferString("")),
		ContentLength: int64(len("")),
		Request:       req,
	}

	err := CreateFromResponse(res)

	asNotFound, ok := err.(*NotFound)
	assert.True(t, ok)

	assert.Equal(t, asNotFound.Path(), path)
}
