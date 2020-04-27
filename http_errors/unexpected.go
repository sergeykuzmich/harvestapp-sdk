package errors

import "fmt"

// Unexpected represents all other non 2** HTTP responses which are not covered
// with custom error handlers.
type Unexpected struct {
	status int
	path   string
	body   []byte
}

func (e *Unexpected) Error() string {
	return fmt.Sprintf("Unexpected Error: %d %s %s", e.status, e.path, e.body)
}

// Details provides extended info about the error happened.
func (e *Unexpected) Details() []byte {
	return e.body
}

// Status shows exact HTTP status which was returned on request.
func (e *Unexpected) Status() int {
	return e.status
}

// Path contains URI the error happened on.
func (e *Unexpected) Path() string {
	return e.path
}

func createUnexpected(status int, path string, body []byte) *Unexpected {
	return &Unexpected{
		status: status,
		path:   path,
		body:   body,
	}
}
