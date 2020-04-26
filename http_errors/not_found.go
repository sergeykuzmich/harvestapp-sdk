package errors

import "fmt"

type NotFound struct {
	path string
	body []byte
}

func (e *NotFound) Error() string {
	return fmt.Sprintf("Not Found: %s", e.path)
}

func (e *NotFound) Details() []byte {
	return e.body
}

func (e *NotFound) Path() string {
	return e.path
}

func createNotFound(path string, body []byte) *NotFound {
	return &NotFound{
		path: path,
		body: body,
	}
}
