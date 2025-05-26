package endpoint

// Specifies text extraction order.
type TextOrder string

const (
	// Represents a Stream text order.
	Stream TextOrder = "Stream"

	// Represents a Visible text order.
	Visible TextOrder = "Visible"

	// Represents a VisibleExtraSpace text order.
	VisibleExtraSpace TextOrder = "VisibleExtraSpace"
)
