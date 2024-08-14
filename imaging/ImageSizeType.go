package imaging

// Specifies the type of image size.
type ImageSizeType string

const (
	// Dpi image size type.
	Dpi ImageSizeType = "Dpi"

	// Fixed image size type.
	Fixed ImageSizeType = "Fixed"

	// Max image size type.
	Max ImageSizeType = "Max"

	// Percentage image size type.
	Percentage ImageSizeType = "Percentage"
)
