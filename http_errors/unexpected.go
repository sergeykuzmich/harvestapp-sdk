package errors

import "fmt"

type Unexpected struct {
	status int
	path   string
	body   []byte
}

func (e *Unexpected) Error() string {
	return fmt.Sprintf("Unexpected Error: %d %s %s", e.status, e.path, e.body)
}

func (e *Unexpected) Details() []byte {
	return e.body
}

func (e *Unexpected) Status() int {
	return e.status
}

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
