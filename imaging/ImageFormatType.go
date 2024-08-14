package imaging

// Enum representing supported image formats.
type ImageFormatType string

const (
	// JPEG image format.
	JPEG ImageFormatType = "Jpeg"

	// GIF image format.
	GIF ImageFormatType = "Gif"

	// BMP image format.
	BMP ImageFormatType = "Bmp"

	// PNG image format.
	PNG ImageFormatType = "Png"

	// TIFF image format.
	TIFF ImageFormatType = "Tiff"
)
