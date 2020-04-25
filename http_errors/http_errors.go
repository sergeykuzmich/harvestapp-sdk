package http_errors

import (
	"io/ioutil"
	"net/http"
)

type HttpError interface {
	error

	// RAW response body
	Details() []byte
	Path() string
}

func CreateFromResponse(response *http.Response) HttpError {
	body, _ := ioutil.ReadAll(response.Body)
	path := response.Request.URL.Path

	var error HttpError

	switch status := response.StatusCode; {
		case status == 404:
			error = createNotFound(path, body)
		case status == 423:
			error = createUnprocessableEntity(path, body)
		default:
			error = createUnexpected(status, path, body)
	}

	return error
}


