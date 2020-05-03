package errors

import (
	"io/ioutil"
	"net/http"
)

// HTTPError is an interface for errors which can be returned.
type HTTPError interface {
	error
	Details() []byte
	Path() string
}

// CreateFromResponse converts http.Response with non 2** status code
// to Error with HTTPError interface.
func CreateFromResponse(response *http.Response) HTTPError {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	path := response.Request.URL.Path

	var error HTTPError

	switch status := response.StatusCode; {
	case status == 401:
		error = createUnauthorized(path, body)
	case status == 403:
		error = createForbidden(path, body)
	case status == 404:
		error = createNotFound(path, body)
	case status == 422:
		error = createUnprocessableEntity(path, body)
	case status == 429:
		error = createTooManyRequests(path, body)
	case status >= 500:
		error = createServerError(status, path, body)
	default:
		error = createUnexpected(status, path, body)
	}

	return error
}
