package input

type InputType string

// Specifies horizontal alignment. Center alignment is the default value for alignments.
const (

	// Dlex input.
	DlexInput InputType = "dlex"

	// Pdf input.
	PdfInput InputType = "pdf"

	// Image input.
	ImageInput InputType = "image"

	// Page input.
	PageInput InputType = "page"

	// Html input.
	HtmlInput InputType = "html"

	//Word input
	WordInput InputType = "word"
)
