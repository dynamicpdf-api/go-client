package imaging

// Represents TIFF image format with color format.
type TiffImageFormat struct {
	ImageFormat // Embedded struct ImageFormat
}

// NewTiffImageFormat initializes a new instance of the TiffImageFormat class.
func NewTiffImageFormat() *TiffImageFormat {
	var p = TiffImageFormat{ImageFormat: *NewImageFormat(TIFF)}
	return &p
}

func (t *TiffImageFormat) SetMultiPage(multiPage bool) {
	t.multiPage = multiPage
}

// GetMultiPage returns the multi-page support flag for TIFF images.
func (t *TiffImageFormat) GetMultiPage() bool {
	return t.multiPage
}

// SetColorFormat sets the color format for TIFF images.
func (t *TiffImageFormat) SetColorFormat(colorFormat ColorFormat) {
	t.colorFormat = colorFormat
}

// GetColorFormat returns the color format for TIFF images.
func (t *TiffImageFormat) GetColorFormat() ColorFormat {
	return t.colorFormat
}
