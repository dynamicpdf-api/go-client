package endpoint

// Represents the pdf security info response.
type PdfSecurityInfoResponse struct {
	JsonResponse
	Content *PdfSecurityInfo
}

// Initializes a new instance of the `PdfSecurityInfoResponse` class.
func NewPdfSecurityInfoResponse() *PdfSecurityInfoResponse {
	var ep PdfSecurityInfoResponse
	return &ep
}

// Gets the PdfSecurityInfo content.
func (pr *PdfSecurityInfoResponse) JsonContent() *PdfSecurityInfo {
	return pr.Content
}
