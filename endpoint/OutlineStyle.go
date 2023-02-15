package endpoint

type OutlineStyle int

// Specifies horizontal alignment. Center alignment is the default value for alignments.

const (

	// Bold.
	Bold OutlineStyle = 0

	// BoldItalic.
	BoldItalic OutlineStyle = 1

	// Italic.
	Italic OutlineStyle = 2

	// Regular.
	Regular OutlineStyle = 3
)
