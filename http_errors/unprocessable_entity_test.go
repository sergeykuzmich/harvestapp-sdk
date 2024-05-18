package errors

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
	path := "/422"
	body := `{
		"default_hourly_rate":120.0
	}`
	errorMessage := fmt.Sprintf("Unprocessable Entity: %s %s", path, []byte(body))

	var err *UnprocessableEntity
	err = createUnprocessableEntity(path, []byte(body))

	assert.Equal(t, err.Path(), path)
	assert.Equal(t, err.Details(), []byte(body))
	assert.Equal(t, err.Error(), errorMessage)
}

func TestCreateFromUnprocessableEntityResponse(t *testing.T) {
	path := "/422"
	reqBody := `{
		"default_hourly_rate":120.0
	}`
	resBody := `{
		"message": "Name can't be blank"
	}`

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(reqBody)

	req, _ := http.NewRequest("POST", path, buffer)

	res := &http.Response{
		Status:        "422 Unprocessable Entity",
		StatusCode:    422,
		Proto:         "HTTP/1.1",
		Body:          ioutil.NopCloser(bytes.NewBufferString(string(resBody))),
		ContentLength: int64(len(string(resBody))),
		Request:       req,
	}

	err := CreateFromResponse(res)

	asUnprocessableEntity, ok := err.(*UnprocessableEntity)
	assert.True(t, ok)

	assert.Equal(t, asUnprocessableEntity.Path(), path)
	assert.Equal(t, asUnprocessableEntity.Details(), []byte(resBody))
}
