package errors

import "fmt"

// Unauthorized is returned in case of invalid Harvest credentials.
type Unauthorized struct {
	path string
	body []byte
}

func (e *Unauthorized) Error() string {
	return fmt.Sprintf("Unauthorized: %s", e.path)
}

// Details provides extended info about the error happened.
func (e *Unauthorized) Details() []byte {
	return e.body
}

// Path contains URI the error happened on.
func (e *Unauthorized) Path() string {
	return e.path
}

func createUnauthorized(path string, body []byte) HTTPError {
	return &Unauthorized{
		path: path,
		body: body,
	}
}
