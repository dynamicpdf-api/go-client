package endpoint

import (
	"bytes"

	"github.com/google/uuid"
)

// Represnts the pdf resource.
type PdfTextResponse struct {
	JsonResponse
}

// Initializes a new instance of the `PdfResponse` class.
func NewPdfTextResponse() *PdfTextResponse {
	var ep PdfTextResponse
	return &ep
}

// Gets the collection of PdfContent.
func (pr *PdfTextResponse) Content() *bytes.Buffer {
	return pr.content
}

func (r *PdfTextResponse) ClientError() error {
	return r.clientError
}

// Gets the boolean, indicating the response's status.
func (r *PdfTextResponse) IsSuccessful() bool {
	return r.isSuccessful
}

// Gets the error message.
func (r *PdfTextResponse) ErrorMessage() string {
	return r.errorMessage
}

// Gets the error id.
func (r *PdfTextResponse) ErrorId() uuid.UUID {
	return r.errorId
}

// Gets the status code.
func (r *PdfTextResponse) StatusCode() int {
	return r.statusCode
}

// Gets the error json.
func (r *PdfTextResponse) ErrorJson() string {
	return r.errorJson
}
