package imaging

// Represents an image size defined by DPI (Dots Per Inch).
type DpiImageSize struct {
	ImageSize
}

// NewDpiImageSize creates a new instance of DpiImageSize and sets the image size type to DPI.
func NewDpiImageSize() *DpiImageSize {
	var ep DpiImageSize
	ep.Type = Dpi
	return &ep
}

// GetHorizontalDpi returns the horizontal DPI of the image.
func (d *DpiImageSize) GetHorizontalDpi() int {
	return d.horizontalDpi
}

// SetHorizontalDpi sets the horizontal DPI of the image.
func (d *DpiImageSize) SetHorizontalDpi(value int) {
	d.horizontalDpi = value
}

// GetVerticalDpi returns the vertical DPI of the image.
func (d *DpiImageSize) GetVerticalDpi() int {
	return d.verticalDpi
}

// SetVerticalDpi sets the vertical DPI of the image.
func (d *DpiImageSize) SetVerticalDpi(value int) {
	d.verticalDpi = value
}
