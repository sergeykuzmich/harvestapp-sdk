package errors

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateForbiddenError(t *testing.T) {
	path := "/403"
	errorMessage := fmt.Sprintf("Forbidden: %s", path)
	var body []byte

	var err *Forbidden
	err = createForbidden(path, body)

	assert.Equal(t, err.Path(), path)
	assert.Equal(t, err.Details(), body)
	assert.Equal(t, err.Error(), errorMessage)
}

func TestCreateFromForbiddenResponse(t *testing.T) {
	path := "/403"
	errorMessage := fmt.Sprintf("Forbidden: %s", path)

	req, _ := http.NewRequest("GET", path, nil)
	res := &http.Response{
		Status:        "403 Forbidden",
		StatusCode:    403,
		Proto:         "HTTP/1.1",
		Body:          ioutil.NopCloser(bytes.NewBufferString("")),
		ContentLength: int64(len("")),
		Request:       req,
	}

	err := CreateFromResponse(res)

	asForbidden, ok := err.(*Forbidden)
	assert.True(t, ok)

	assert.Equal(t, asForbidden.Path(), path)
	assert.Equal(t, asForbidden.Error(), errorMessage)
}
