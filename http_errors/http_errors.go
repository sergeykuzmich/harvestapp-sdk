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

type httpErrorCallback func(string, []byte) HTTPError

// CreateFromResponse converts http.Response with non 2** status code
// to Error with HTTPError interface.
func CreateFromResponse(response *http.Response) HTTPError {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	path := response.Request.URL.Path

	errorsMap := map[int]httpErrorCallback{
		401: createUnauthorized,
		403: createForbidden,
		404: createNotFound,
		422: createUnprocessableEntity,
		429: createTooManyRequests,
	}

	status := response.StatusCode
	if errorHandler, ok := errorsMap[status]; ok {
		return errorHandler(path, body)
	}

	return createUnknown(status, path, body)
}
