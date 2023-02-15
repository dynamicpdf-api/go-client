package position

type VAlign int

// Specifies horizontal alignment. Center alignment is the default value for alignments.
const (

	// Align left.
	Top VAlign = 0

	// Align center.
	Middle VAlign = 1

	// Align right.
	Bottom VAlign = 2
)
