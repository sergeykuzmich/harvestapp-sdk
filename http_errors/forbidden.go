package errors

import "fmt"

type Forbidden struct {
	path string
	body []byte
}

func (e *Forbidden) Error() string {
	return fmt.Sprintf("Forbidden: %s", e.path)
}

func (e *Forbidden) Details() []byte {
	return e.body
}

func (e *Forbidden) Path() string {
	return e.path
}

func createForbidden(path string, body []byte) *Forbidden {
	return &Forbidden{
		path: path,
		body: body,
	}
}
