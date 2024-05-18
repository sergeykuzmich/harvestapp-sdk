package errors

import (
	"bytes"
	"encoding/json"
	"errors"
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

	var err HTTPError
	err = createUnprocessableEntity(path, []byte(body))

	var asUnprocessableEntity *UnprocessableEntity
	ok := errors.As(err, &asUnprocessableEntity)
	assert.True(t, ok)

	assert.Equal(t, asUnprocessableEntity.Path(), path)
	assert.Equal(t, asUnprocessableEntity.Details(), []byte(body))
	assert.Equal(t, asUnprocessableEntity.Error(), errorMessage)
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
	_ = json.NewEncoder(buffer).Encode(reqBody)

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

	var asUnprocessableEntity *UnprocessableEntity
	ok := errors.As(err, &asUnprocessableEntity)
	assert.True(t, ok)

	assert.Equal(t, asUnprocessableEntity.Path(), path)
	assert.Equal(t, asUnprocessableEntity.Details(), []byte(resBody))
}
