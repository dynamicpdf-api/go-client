package endpoint

/**
 * Specifies horizontal alignment. Center alignment is the default value for alignments.
 */
type PdfContent struct {

	/* Gets or sets the page number. */
	PageNumber int `json:"pageNumber,omitempty"`

	/* Gets or sets the text in the pdf. */
	Text string `json:"text,omitempty"`
}
