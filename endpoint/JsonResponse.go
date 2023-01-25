package endpoint

import (
	"bytes"

	"github.com/google/uuid"
)

/**Represents the base class for json response. */
type JsonResponse struct {
	Response
	clientError   error
	isSuccessful  bool
	errorMessage  string
	errorId       uuid.UUID
	statusCode    int
	errorJson     string
	content       *bytes.Buffer
	statusMessage string
}

func (r *JsonResponse) ClientError() error {
	return r.clientError
}

// Gets the boolean, indicating the response's status.
func (r *JsonResponse) IsSuccessful() bool {
	return r.isSuccessful
}

// Gets the error message.
func (r *JsonResponse) ErrorMessage() string {
	return r.errorMessage
}

// Gets the error id.
func (r *JsonResponse) ErrorId() uuid.UUID {
	return r.errorId
}

// Gets the status code.
func (r *JsonResponse) StatusCode() int {
	return r.statusCode
}

func (r *JsonResponse) SetStatusMessage(statusMessage string) {
	r.statusMessage = statusMessage
}

// Gets the error json.
func (r *JsonResponse) ErrorJson() string {
	return r.errorJson
}
