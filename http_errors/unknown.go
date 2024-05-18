package errors

import "fmt"

// Unknown represents all other non 2** HTTP responses which are not covered
// with custom error handlers.
type Unknown struct {
	status int
	path   string
	body   []byte
}

func (e *Unknown) Error() string {
	return fmt.Sprintf("Unknown Error: %d %s %s", e.status, e.path, e.body)
}

// Details provides extended info about the error happened.
func (e *Unknown) Details() []byte {
	return e.body
}

// Status shows exact HTTP status which was returned on request.
func (e *Unknown) Status() int {
	return e.status
}

// Path contains URI the error happened on.
func (e *Unknown) Path() string {
	return e.path
}

func createUnknown(status int, path string, body []byte) HTTPError {
	return &Unknown{
		status: status,
		path:   path,
		body:   body,
	}
}
