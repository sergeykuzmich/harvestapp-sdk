package errors

import "fmt"

// TooManyRequests is returned after exceeding API rate limits.
// > https://help.getharvest.com/api-v2/introduction/overview/general/#rate-limiting
type TooManyRequests struct {
	path string
	body []byte
}

func (e *TooManyRequests) Error() string {
	return fmt.Sprintf("Too Many Requests: %s", e.path)
}

// Details provides extended info about the error happened.
func (e *TooManyRequests) Details() []byte {
	return e.body
}

// Path contains URI the error happened on.
func (e *TooManyRequests) Path() string {
	return e.path
}

func createTooManyRequests(path string, body []byte) *TooManyRequests {
	return &TooManyRequests{
		path: path,
		body: body,
	}
}
