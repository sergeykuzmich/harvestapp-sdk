package sdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiError interface {
	error
	StatusCode() int
	ResponseBody() []byte
	Path() string
}

func CreateFromResponse(response *http.Response) ApiError {

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic("Could't read response body")
	}

	path := response.Request.URL.Path

	switch status := response.StatusCode; {
	case status == 404:
		return &NotFoundError{body: body, path: path}
	case status == 422:
		return &UnprocessableEntityError{body: body, path: path}
	case status == 429:
		return &UnprocessableEntityError{body: body, path: path}
	case status >= 500:
		return &ServerError{statusCode: status, body: body, path: path}
	default:
		return &UnknownError{statusCode: status, body: body, path: path}
	}
}

type NotFoundError struct {
	body []byte
	path string
}

func (err *NotFoundError) StatusCode() int {
	return 404
}

func (err *NotFoundError) ResponseBody() []byte {
	return err.body
}

func (err *NotFoundError) Path() string {
	return err.path
}

func (err *NotFoundError) Error() string {
	return fmt.Sprintf("The object you requested can’t be found: %s", err.Path())
}

type ForbiddenError struct {
	body []byte
	path string
}

func (err *ForbiddenError) StatusCode() int {
	return 403
}

func (err *ForbiddenError) ResponseBody() []byte {
	return err.body
}

func (err *ForbiddenError) Error() string {
	return fmt.Sprintf("The object you requested was found but you don’t have authorization to perform your request: %s", err.path)
}

func (err *ForbiddenError) Path() string {
	return err.path
}

type UnprocessableEntityError struct {
	body []byte
	path string
}

func (err UnprocessableEntityError) Error() string {
	return fmt.Sprintf("There were errors processing your request (%s): %s", err.path, err.body)
}

func (err UnprocessableEntityError) StatusCode() int {
	return 422
}

func (err UnprocessableEntityError) ResponseBody() []byte {
	return err.body
}

func (err UnprocessableEntityError) Path() string {
	return err.path
}

type TooManyRequestsError struct {
	body []byte
	path string
}

func (err TooManyRequestsError) Error() string {
	return fmt.Sprintf("Your request has been throttled: %s", err.path)
}

func (err TooManyRequestsError) StatusCode() int {
	return 429
}

func (err TooManyRequestsError) ResponseBody() []byte {
	return err.body
}

func (err TooManyRequestsError) Path() string {
	return err.path
}

type ServerError struct {
	statusCode int
	body       []byte
	path       string
}

func (err ServerError) Error() string {
	return fmt.Sprintf("There was a server error: %s", err.path)
}

func (err ServerError) StatusCode() int {
	return err.statusCode
}

func (err ServerError) ResponseBody() []byte {
	return err.body
}

func (err ServerError) Path() string {
	return err.path
}

type UnknownError struct {
	statusCode int
	body       []byte
	path       string
}

func (err UnknownError) Error() string {
	return fmt.Sprintf("Unknown error: %s", err.path)
}

func (err UnknownError) StatusCode() int {
	return err.statusCode
}

func (err UnknownError) ResponseBody() []byte {
	return err.body
}

func (err UnknownError) Path() string {
	return err.path
}
