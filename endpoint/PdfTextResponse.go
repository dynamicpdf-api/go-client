package endpoint

import (
	"bytes"
)

// Represnts the pdf resource.
type PdfTextResponse struct {
	JsonResponse
	TextContent []PdfContent
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
