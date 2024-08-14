package imaging

// PdfImageResponse represents the JSON structure in Go
type PdfImageResponse struct {
	response
	ImageFormat   string  `json:"imageFormat"`
	Images        []Image `json:"images"`
	ContentType   string  `json:"contentType"`
	HorizontalDpi int     `json:"horizontalDpi"`
	VerticalDpi   int     `json:"verticalDpi"`
}

// Image represents an individual image in the PdfImageResponse
type Image struct {
	PageNumber  string `json:"pageNumber"`
	Data        string `json:"data"`
	BilledPages int    `json:"billedPages"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
}

// Initializes a new instance of the PdfImageResponse class.
func NewPdfImageResponse() *PdfImageResponse {
	var p = PdfImageResponse{}
	return &p
}
