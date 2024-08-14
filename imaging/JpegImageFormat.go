package imaging

// JpegImageFormat represents JPEG image format with quality.
type JpegImageFormat struct {
	ImageFormat
}

// SetQuality sets the quality of the JPEG image.
func (j *JpegImageFormat) SetQuality(quality int) {
	j.quality = quality
}

// GetQuality returns the quality of the JPEG image.
func (j *JpegImageFormat) GetQuality() int {
	return j.quality
}

// Initializes a new instance of the JpegImageFormat class.
func NewJpegImageFormat() *JpegImageFormat {
	var p = JpegImageFormat{ImageFormat: *NewImageFormat(JPEG)}
	return &p
}
