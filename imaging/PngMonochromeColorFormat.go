package imaging

// Represents monochrome color format for PNG with black threshold.
type PngMonochromeColorFormat struct {
	PngColorFormat
}

/*
  Initializes a new instance of the PngMonochromeColorFormat class with monochrome color format type.
*/
func NewPngMonochromeColorFormat() *PngMonochromeColorFormat {
	var p = PngMonochromeColorFormat{PngColorFormat: *NewPngColorFormat(Monochrome)}
	return &p
}

func (p *PngMonochromeColorFormat) SetBlackThreshold(threshold int) {
	p.blackThreshold = threshold
}

// GetBlackThreshold returns the black threshold for monochrome PNG images.
func (p *PngMonochromeColorFormat) GetBlackThreshold() int {
	return p.blackThreshold
}

// SetDitheringPercent sets the dithering percentage for PNG images.
func (p *PngMonochromeColorFormat) SetDitheringPercent(percent int) {
	p.ditheringPercent = percent
}

// GetDitheringPercent returns the dithering percentage for PNG images.
func (p *PngMonochromeColorFormat) GetDitheringPercent() int {
	return p.ditheringPercent
}

// SetDitheringAlgorithm sets the dithering algorithm for PNG images.
func (p *PngMonochromeColorFormat) SetDitheringAlgorithm(algorithm DitheringAlgorithm) {
	p.ditheringAlgorithm = algorithm
}

// GetDitheringAlgorithm returns the dithering algorithm for PNG images.
func (p *PngMonochromeColorFormat) GetDitheringAlgorithm() DitheringAlgorithm {
	return p.ditheringAlgorithm
}
