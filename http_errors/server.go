package errors

import "fmt"

// Server represents any `server` errors with status codes of 5** series.
type Server struct {
	status int
	path   string
	body   []byte
}

func (e *Server) Error() string {
	return fmt.Sprintf("%d %s %s", e.status, e.path, e.body)
}

// Details provides extended info about the error happened.
func (e *Server) Details() []byte {
	return e.body
}

// Status shows exact HTTP status which was returned on request.
func (e *Server) Status() int {
	return e.status
}

// Path contains URI the error happened on.
func (e *Server) Path() string {
	return e.path
}

func createServerError(status int, path string, body []byte) *Server {
	return &Server{
		status: status,
		path:   path,
		body:   body,
	}
}
