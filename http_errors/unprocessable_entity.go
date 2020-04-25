package http_errors

import "fmt"

type UnprocessableEntity struct {
    path string
    body []byte
}

func (e *UnprocessableEntity) Error() string {
    return fmt.Sprintf("Unprocessable Entity: %s %s", e.path, e.body)
}

func (e *UnprocessableEntity) Details() []byte {
    return e.body
}

func (e *UnprocessableEntity) Path() string {
    return e.path
}

func createUnprocessableEntity(path string, body []byte) *UnprocessableEntity {
    return &UnprocessableEntity{
        path: path,
        body: body,
    }
}
