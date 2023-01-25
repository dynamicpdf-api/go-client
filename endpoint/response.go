// Represents the base class for response.
package endpoint

import (
	"bytes"

	"github.com/google/uuid"
)

type Response interface {
	ClientError() error
	IsSuccessful() bool
	ErrorMessage() string
	ErrorId() uuid.UUID
	StatusCode() int
	ErrorJson() string
}

// Response from server
type response struct {
	Response
	clientError  error
	isSuccessful bool
	errorMessage string
	errorId      uuid.UUID
	statusCode   int
	errorJson    string
	content      *bytes.Buffer
}

var _ Response = (*response)(nil)

// Gets the error json.
func (r *response) ClientError() error {
	return r.clientError
}

// Gets the boolean, indicating the response's status.
func (r *response) IsSuccessful() bool {
	return r.isSuccessful
}

// Gets the error message.
func (r *response) ErrorMessage() string {
	return r.errorMessage
}

// Gets the error id.
func (r *response) ErrorId() uuid.UUID {
	return r.errorId
}

// Gets the status code.
func (r *response) StatusCode() int {
	return r.statusCode
}

// Gets the error json.
func (r *response) ErrorJson() string {
	return r.errorJson
}
