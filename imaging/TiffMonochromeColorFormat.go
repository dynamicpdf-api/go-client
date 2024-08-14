package imaging

// Represents monochrome color format for TIFF with black threshold and compression type.
type TiffMonochromeColorFormat struct {
	TiffColorFormat
}

/*
  Initializes a new instance of the TiffMonochromeColorFormat class with monochrome color format type.
*/
func NewTiffMonochromeColorFormat() *TiffMonochromeColorFormat {
	var p = TiffMonochromeColorFormat{TiffColorFormat: *NewTiffColorFormat(Monochrome)}
	return &p
}

// SetBlackThreshold sets the black threshold for monochrome TIFF images.
func (t *TiffMonochromeColorFormat) SetBlackThreshold(threshold int) {
	t.blackThreshold = threshold
}

// GetBlackThreshold returns the black threshold for monochrome TIFF images.
func (t *TiffMonochromeColorFormat) GetBlackThreshold() int {
	return t.blackThreshold
}

// SetCompressionType sets the compression type for monochrome TIFF images.
func (t *TiffMonochromeColorFormat) SetCompressionType(compression CompressionType) {
	t.compressionType = compression
}

// GetCompressionType returns the compression type for monochrome TIFF images.
func (t *TiffMonochromeColorFormat) GetCompressionType() CompressionType {
	return t.compressionType
}

// SetDitheringPercent sets the dithering percentage for TIFF images.
func (t *TiffMonochromeColorFormat) SetDitheringPercent(percent int) {
	t.ditheringPercent = percent
}

// GetDitheringPercent returns the dithering percentage for TIFF images.
func (t *TiffMonochromeColorFormat) GetDitheringPercent() int {
	return t.ditheringPercent
}

// SetDitheringAlgorithm sets the dithering algorithm for TIFF images.
func (t *TiffMonochromeColorFormat) SetDitheringAlgorithm(algorithm DitheringAlgorithm) {
	t.ditheringAlgorithm = algorithm
}

// GetDitheringAlgorithm returns the dithering algorithm for TIFF images.
func (t *TiffMonochromeColorFormat) GetDitheringAlgorithm() DitheringAlgorithm {
	return t.ditheringAlgorithm
}
