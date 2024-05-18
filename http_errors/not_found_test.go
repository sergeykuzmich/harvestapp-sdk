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

func TestCreateNotFoundError(t *testing.T) {
	path := "/404"
	errorMessage := fmt.Sprintf("Not Found: %s", path)
	var body []byte

	var err HTTPError
	err = createNotFound(path, body)

	var asNotFound *NotFound
	ok := errors.As(err, &asNotFound)
	assert.True(t, ok)

	assert.Equal(t, asNotFound.Path(), path)
	assert.Equal(t, asNotFound.Details(), body)
	assert.Equal(t, asNotFound.Error(), errorMessage)
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

	var asNotFound *NotFound
	ok := errors.As(err, &asNotFound)
	assert.True(t, ok)

	assert.Equal(t, asNotFound.Path(), path)
}
