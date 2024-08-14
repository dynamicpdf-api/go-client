package imaging

// Represents indexed color format for TIFF.
type TiffIndexedColorFormat struct {
	TiffColorFormat
}

/*
  Initializes a new instance of the TiffIndexedColorFormat class  with indexed color format type.
*/
func NewTiffIndexedColorFormat() *TiffIndexedColorFormat {
	var p = TiffIndexedColorFormat{TiffColorFormat: *NewTiffColorFormat(Indexed)}
	return &p
}

// SetQuantizationAlgorithm sets the quantization algorithm for TIFF images.
func (t *TiffIndexedColorFormat) SetQuantizationAlgorithm(algorithm QuantizationAlgorithm) {
	t.quantizationAlgorithm = algorithm
}

// GetQuantizationAlgorithm returns the quantization algorithm for TIFF images.
func (t *TiffIndexedColorFormat) GetQuantizationAlgorithm() QuantizationAlgorithm {
	return t.quantizationAlgorithm
}

// SetDitheringPercent sets the dithering percentage for TIFF images.
func (t *TiffIndexedColorFormat) SetDitheringPercent(percent int) {
	t.ditheringPercent = percent
}

// GetDitheringPercent returns the dithering percentage for TIFF images.
func (t *TiffIndexedColorFormat) GetDitheringPercent() int {
	return t.ditheringPercent
}

// SetDitheringAlgorithm sets the dithering algorithm for TIFF images.
func (t *TiffIndexedColorFormat) SetDitheringAlgorithm(algorithm DitheringAlgorithm) {
	t.ditheringAlgorithm = algorithm
}

// GetDitheringAlgorithm returns the dithering algorithm for TIFF images.
func (t *TiffIndexedColorFormat) GetDitheringAlgorithm() DitheringAlgorithm {
	return t.ditheringAlgorithm
}
