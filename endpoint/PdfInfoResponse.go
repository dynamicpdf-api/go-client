package endpoint

import "bytes"

// Represents the pdf information response.
type PdfInfoResponse struct {
	JsonResponse
}

// Initializes a new instance of the `PdfInfoResponse` class.
func NewPdfInfoResponse() *PdfInfoResponse {
	var ep PdfInfoResponse
	return &ep
}

// Gets the pdf information content.
func (pr *PdfInfoResponse) Content() *bytes.Buffer {
	return pr.content
}
