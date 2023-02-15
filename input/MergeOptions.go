package input

// Represents different options for merging PDF documents.
type MergeOptions struct {

	/* Gets or sets a boolean indicating whether to import document information when merging. */
	DocumentInfo bool

	// Gets or sets a boolean indicating whether to import document level JavaScript when merging.
	DocumentJavaScript bool

	// Gets or sets a boolean indicating whether to import document properties when merging.
	DocumentProperties bool

	// Gets or sets a boolean indicating whether to import embedded files when merging.
	EmbeddedFiles bool

	// Gets or sets a boolean indicating whether to import form fields when merging.
	FormFields bool

	// Gets or sets a boolean indicating whether to import XFA form data when merging.
	FormsXfaData bool

	// Gets or sets a boolean indicating whether to import logical structure (tagging information) when merging.
	LogicalStructure bool

	// Gets or sets a boolean indicating whether to import document's opening action (initial page and zoom settings) when merging.
	OpenAction bool

	// Gets or sets a boolean indicating whether to import optional content when merging.
	OptionalContentInfo bool

	// Gets or sets a boolean indicating whether to import outlines and bookmarks when merging.
	Outlines bool

	// Gets or sets a boolean indicating whether to import OutputIntent when merging.
	OutputIntent bool

	// Gets or sets a boolean indicating whether to import PageAnnotations when merging.
	PageAnnotations bool

	// Gets or sets a boolean indicating whether to import PageLabelsAndSections when merging.
	PageLabelsAndSections bool

	/*
		Gets or sets the root form field for imported form fields.
		Useful when merging a PDF repeatedly to have a better
		control on the form field names.
	*/
	RootFormField bool

	// Gets or sets a boolean indicating whether to import XmpMetadata when merging.
	XmpMetadata bool
}

func NewMergeOptions() MergeOptions {
	var p MergeOptions
	return p
}
