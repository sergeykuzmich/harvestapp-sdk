package errors

import "fmt"

// UnprocessableEntity resresents any kind of resource validation errors.
type UnprocessableEntity struct {
	path string
	body []byte
}

func (e *UnprocessableEntity) Error() string {
	return fmt.Sprintf("Unprocessable Entity: %s %s", e.path, e.body)
}

// Details provides extended info about the error happened.
func (e *UnprocessableEntity) Details() []byte {
	return e.body
}

// Path contains URI the error happened on.
func (e *UnprocessableEntity) Path() string {
	return e.path
}

func createUnprocessableEntity(path string, body []byte) *UnprocessableEntity {
	return &UnprocessableEntity{
		path: path,
		body: body,
	}
}
