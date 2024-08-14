package imaging

// Represents indexed color format for PNG.
type PngIndexedColorFormat struct {
	PngColorFormat
}

/*
  Initializes a new instance of the PngIndexedColorFormat class with indexed color format type.
*/
func NewPngIndexedColorFormat() *PngIndexedColorFormat {
	var p = PngIndexedColorFormat{PngColorFormat: *NewPngColorFormat(Indexed)}
	return &p
}

// SetQuantizationAlgorithm sets the quantization algorithm for PNG images.
func (p *PngIndexedColorFormat) SetQuantizationAlgorithm(algorithm QuantizationAlgorithm) {
	p.quantizationAlgorithm = algorithm
}

// GetQuantizationAlgorithm returns the quantization algorithm for PNG images.
func (p *PngIndexedColorFormat) GetQuantizationAlgorithm() QuantizationAlgorithm {
	return p.quantizationAlgorithm
}

// SetDitheringPercent sets the dithering percentage for PNG images.
func (p *PngIndexedColorFormat) SetDitheringPercent(percent int) {
	p.ditheringPercent = percent
}

// GetDitheringPercent returns the dithering percentage for PNG images.
func (p *PngIndexedColorFormat) GetDitheringPercent() int {
	return p.ditheringPercent
}

// SetDitheringAlgorithm sets the dithering algorithm for PNG images.
func (p *PngIndexedColorFormat) SetDitheringAlgorithm(algorithm DitheringAlgorithm) {
	p.ditheringAlgorithm = algorithm
}

// GetDitheringAlgorithm returns the dithering algorithm for PNG images.
func (p *PngIndexedColorFormat) GetDitheringAlgorithm() DitheringAlgorithm {
	return p.ditheringAlgorithm
}
