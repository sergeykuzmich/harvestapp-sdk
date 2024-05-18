package errors

import "fmt"

// NotFound is returned in case of requesting non-existing resource.
type NotFound struct {
	path string
	body []byte
}

func (e *NotFound) Error() string {
	return fmt.Sprintf("Not Found: %s", e.path)
}

// Details provides extended info about the error happened.
func (e *NotFound) Details() []byte {
	return e.body
}

// Path contains URI the error happened on.
func (e *NotFound) Path() string {
	return e.path
}

func createNotFound(path string, body []byte) HTTPError {
	return &NotFound{
		path: path,
		body: body,
	}
}
