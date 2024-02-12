package input

// Represents different options for merging PDF documents.
type MergeOptions struct {

	/* Gets or sets a boolean indicating whether to import document information when merging. */
	DocumentInfo *bool `json:"documentInfo,omitempty"`

	// Gets or sets a boolean indicating whether to import document level JavaScript when merging.
	DocumentJavaScript *bool `json:"documentJavaScript,omitempty"`

	// Gets or sets a boolean indicating whether to import document properties when merging.
	DocumentProperties *bool `json:"documentProperties,omitempty"`

	// Gets or sets a boolean indicating whether to import embedded files when merging.
	EmbeddedFiles *bool `json:"embeddedFiles,omitempty"`

	// Gets or sets a boolean indicating whether to import form fields when merging.
	FormFields *bool `json:"formFields,omitempty"`

	// Gets or sets a boolean indicating whether to import XFA form data when merging.
	FormsXfaData *bool `json:"formsXfaData,omitempty"`

	// Gets or sets a boolean indicating whether to import logical structure (tagging information) when merging.
	LogicalStructure *bool `json:"logicalStructure,omitempty"`

	// Gets or sets a boolean indicating whether to import document's opening action (initial page and zoom settings) when merging.
	OpenAction *bool `json:"openAction,omitempty"`

	// Gets or sets a boolean indicating whether to import optional content when merging.
	OptionalContentInfo *bool `json:"optionalContentInfo,omitempty"`

	// Gets or sets a boolean indicating whether to import outlines and bookmarks when merging.
	Outlines *bool `json:"outlines,omitempty"`

	// Gets or sets a boolean indicating whether to import OutputIntent when merging.
	OutputIntent *bool `json:"outputIntent,omitempty"`

	// Gets or sets a boolean indicating whether to import PageAnnotations when merging.
	PageAnnotations *bool `json:"pageAnnotations,omitempty"`

	// Gets or sets a boolean indicating whether to import PageLabelsAndSections when merging.
	PageLabelsAndSections *bool `json:"pageLabelsAndSections,omitempty"`

	/*
		Gets or sets the root form field for imported form fields.
		Useful when merging a PDF repeatedly to have a better
		control on the form field names.
	*/
	RootFormField string `json:"rootFormField,omitempty"`

	// Gets or sets a boolean indicating whether to import XmpMetadata when merging.
	XmpMetadata *bool `json:"xmpMetadata,omitempty"`
}

func NewMergeOptions() MergeOptions {
	var p MergeOptions
	return p
}
