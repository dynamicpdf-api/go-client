package imaging

// Represents GIF image format with dithering properties.
type GifImageFormat struct {
	ImageFormat
}

// Initializes a new instance of the GifImageFormat class and sets the image format type to GIF.
func NewGifImageFormat() *GifImageFormat {
	var p = GifImageFormat{ImageFormat: *NewImageFormat(GIF)}
	return &p
}

// GetDitheringPercent returns the dithering percentage.
func (gif *GifImageFormat) GetDitheringPercent() int {
	return gif.ditheringPercent
}

// SetDitheringPercent sets the dithering percentage.
func (gif *GifImageFormat) SetDitheringPercent(value int) {
	gif.ditheringPercent = value
}

// GetDitheringAlgorithm returns the dithering algorithm.
func (gif *GifImageFormat) GetDitheringAlgorithm() DitheringAlgorithm {
	return gif.ditheringAlgorithm
}

// SetDitheringAlgorithm sets the dithering algorithm.
func (gif *GifImageFormat) SetDitheringAlgorithm(value DitheringAlgorithm) {
	gif.ditheringAlgorithm = value
}
