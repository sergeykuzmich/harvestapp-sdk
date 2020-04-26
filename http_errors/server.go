package errors

import "fmt"

type Server struct {
	status int
	path   string
	body   []byte
}

func (e *Server) Error() string {
	return fmt.Sprintf("%d %s %s", e.status, e.path, e.body)
}

func (e *Server) Details() []byte {
	return e.body
}

func (e *Server) Status() int {
	return e.status
}

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
