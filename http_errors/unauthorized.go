package errors

import "fmt"

type Unauthorized struct {
	path string
	body []byte
}

func (e *Unauthorized) Error() string {
	return fmt.Sprintf("Unauthorized: %s", e.path)
}

func (e *Unauthorized) Details() []byte {
	return e.body
}

func (e *Unauthorized) Path() string {
	return e.path
}

func createUnauthorized(path string, body []byte) *Unauthorized {
	return &Unauthorized{
		path: path,
		body: body,
	}
}
