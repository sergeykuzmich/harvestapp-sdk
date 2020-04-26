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
	case status == 401:
		error = createUnauthorized(path, body)
	case status == 403:
		error = createForbidden(path, body)
	case status == 404:
		error = createNotFound(path, body)
	case status == 422:
		error = createUnprocessableEntity(path, body)
	case status >= 500:
		error = createServerError(status, path, body)
	default:
		error = createUnexpected(status, path, body)
	}

	return error
}
