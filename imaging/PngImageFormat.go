package imaging

// Represents PNG image format with color format.
type PngImageFormat struct {
	ImageFormat // Embedded struct ImageFormat
}

// Initializes a new instance of the PngImageFormat class.
func NewPngImageFormat() *PngImageFormat {
	var p = PngImageFormat{ImageFormat: *NewImageFormat(PNG)}
	return &p
}

func (p *PngImageFormat) GetColorFormat() ColorFormat {
	return p.colorFormat
}

// SetColorFormat sets the color format for PNG images.
func (p *PngImageFormat) SetColorFormat(colorFormat ColorFormat) {
	p.colorFormat = colorFormat
}
