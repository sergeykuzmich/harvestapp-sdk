package http_errors

import "fmt"

type TooManyRequests struct {
	path string
	body []byte
}

func (e *TooManyRequests) Error() string {
	return fmt.Sprintf("Too Many Requests: %s", e.path)
}

func (e *TooManyRequests) Details() []byte {
	return e.body
}

func (e *TooManyRequests) Path() string {
	return e.path
}

func createTooManyRequests(path string, body []byte) *TooManyRequests {
	return &TooManyRequests{
		path: path,
		body: body,
	}
}
