package errors

import "fmt"

// Forbidden is returned on when client can't get a resource.
type Forbidden struct {
	path string
	body []byte
}

func (e *Forbidden) Error() string {
	return fmt.Sprintf("Forbidden: %s", e.path)
}

// Details provides extended info about the error happened.
func (e *Forbidden) Details() []byte {
	return e.body
}

// Path contains URI the error happened on.
func (e *Forbidden) Path() string {
	return e.path
}

func createForbidden(path string, body []byte) *Forbidden {
	return &Forbidden{
		path: path,
		body: body,
	}
}
