package imaging

// Represents BMP image format with color format.
type BmpImageFormat struct {
	ImageFormat
}

// Initializes a new instance of the BmpImageFormat class.
func NewBmpImageFormat() *BmpImageFormat {
	var p = BmpImageFormat{ImageFormat: *NewImageFormat(BMP)}
	return &p
}

// GetColorFormat returns the color format of the BMP image.
func (b *BmpImageFormat) GetColorFormat() ColorFormat {
	return b.colorFormat
}

// SetColorFormat sets the color format of the BMP image.
func (b *BmpImageFormat) SetColorFormat(colorFormat ColorFormat) {
	b.colorFormat = colorFormat
}
