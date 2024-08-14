package imaging

// Enum representing color formats.
type ColorFormatType string

const (
	// Rgb color format.
	Rgb ColorFormatType = "RGB"

	// RgbA color format.
	RgbA ColorFormatType = "RGBA"

	// Grayscale color format.
	Grayscale ColorFormatType = "Grayscale"

	// Monochrome color format.
	Monochrome ColorFormatType = "Monochrome"

	// Indexed color format.
	Indexed ColorFormatType = "Indexed"
)
