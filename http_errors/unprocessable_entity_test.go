package http_errors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUnprocessableEntityError(t *testing.T) {
	path := "/ping"
	body := "{" +
		"\"default_hourly_rate\":120.0" +
		"}"
	errorMessage := fmt.Sprintf("Unprocessable Entity: %s %s", path, []byte(body))

	err := createUnprocessableEntity(path, []byte(body))

	assert.Equal(t, err.Path(), path)
	assert.Equal(t, err.Details(), []byte(body))
	assert.Equal(t, err.Error(), errorMessage)
}

func TestCreateFromUnprocessableEntityResponse(t *testing.T) {
	path := "/tasks"
	req_body := "{" +
		"\"default_hourly_rate\":120.0" +
		"}"
	res_body := "{" +
		"\"message\": \"Name can't be blank\"" +
		"}"

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(req_body)

	req, _ := http.NewRequest("POST", path, buffer)

	res := &http.Response{
		Status:        "422 Unprocessable Entity",
		StatusCode:    422,
		Proto:         "HTTP/1.1",
		Body:          ioutil.NopCloser(bytes.NewBufferString(string(res_body))),
		ContentLength: int64(len(string(res_body))),
		Request:       req,
	}

	err := CreateFromResponse(res)

	err, ok := err.(*UnprocessableEntity)
	assert.True(t, ok)

	assert.Equal(t, err.Path(), path)
	assert.Equal(t, err.Details(), []byte(res_body))
}
