package imaging

// Represents monochrome color format for BMP images.
type BmpMonochromeColorFormat struct {
	BmpColorFormat
}

/*
  Creates object for monochrome color format for BMP image format.
*/
func NewBmpMonochromeColorFormat() *BmpMonochromeColorFormat {
	var p = BmpMonochromeColorFormat{BmpColorFormat: *NewBmpColorFormat(Monochrome)}
	return &p
}

// SetBlackThreshold sets the black threshold for monochrome BMP images.
func (b *BmpMonochromeColorFormat) SetBlackThreshold(threshold int) {
	b.blackThreshold = threshold
}

// GetBlackThreshold returns the black threshold for monochrome BMP images.
func (b *BmpMonochromeColorFormat) GetBlackThreshold() int {
	return b.blackThreshold
}

// SetDitheringPercent sets the dithering percentage for BMP images.
func (b *BmpMonochromeColorFormat) SetDitheringPercent(percent int) {
	b.ditheringPercent = percent
}

// GetDitheringPercent returns the dithering percentage for BMP images.
func (b *BmpMonochromeColorFormat) GetDitheringPercent() int {
	return b.ditheringPercent
}

// SetDitheringAlgorithm sets the dithering algorithm for BMP images.
func (b *BmpMonochromeColorFormat) SetDitheringAlgorithm(algorithm DitheringAlgorithm) {
	b.ditheringAlgorithm = algorithm
}

// GetDitheringAlgorithm returns the dithering algorithm for BMP images.
func (b *BmpMonochromeColorFormat) GetDitheringAlgorithm() DitheringAlgorithm {
	return b.ditheringAlgorithm
}
