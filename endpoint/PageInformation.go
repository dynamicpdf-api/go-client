package endpoint

/**
 * Specifies horizontal alignment. Center alignment is the default value for alignments.
 */
type PageInformation struct {

	/* Gets or sets the page number. */
	PageNumber int `json:"pageNumber,omitempty"`

	/* Gets or sets the width of the page. */
	Width float32 `json:"width,omitempty"`

	/* Gets or sets the height of the page. */
	Height float32 `json:"height,omitempty"`
}
