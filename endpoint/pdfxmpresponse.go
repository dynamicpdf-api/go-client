package endpoint

import (
	"bytes"

	"github.com/google/uuid"
)

// Represnts the pdf resource.
type PdfXmpResponse struct {
	JsonResponse
}

/*
Initializes a new instance of the `XmlResponse` class.
  - @param {string} xmlContent The xml content of the response.
*/
func NewPdfXmpResponse() *PdfXmpResponse {
	var ep PdfXmpResponse
	return &ep
}

// Gets the xml content.
func (pr *PdfXmpResponse) Content() *bytes.Buffer {
	return pr.content
}

func (r *PdfXmpResponse) ClientError() error {
	return r.clientError
}

// Gets the boolean, indicating the response's status.
func (r *PdfXmpResponse) IsSuccessful() bool {
	return r.isSuccessful
}

// Gets the error message.
func (r *PdfXmpResponse) ErrorMessage() string {
	return r.errorMessage
}

// Gets the error id.
func (r *PdfXmpResponse) ErrorId() uuid.UUID {
	return r.errorId
}

// Gets the status code.
func (r *PdfXmpResponse) StatusCode() int {
	return r.statusCode
}

// Gets the error json.
func (r *PdfXmpResponse) ErrorJson() string {
	return r.errorJson
}
