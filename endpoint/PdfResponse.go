// Represents the pdf response.
package endpoint

import "bytes"

/*
 Initializes a new instance of the `PdfResponse` class.
  * @param {Buffer[]} pdfContent The byte array of pdf content.
*/
type PdfResponse struct {
	response
}

// Gets the content of pdf.
func (pr *PdfResponse) Content() *bytes.Buffer {
	return pr.content
}
