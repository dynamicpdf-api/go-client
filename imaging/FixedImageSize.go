package imaging

// Represents an image size with fixed dimensions.
type FixedImageSize struct {
	ImageSize
}

// Initializes a new instance of FixedImageSize and sets the image size type to Fixed.
func NewFixedImageSize() *FixedImageSize {
	var ep FixedImageSize
	ep.Type = Fixed
	return &ep
}

//Gets the width of the image.
func (p *FixedImageSize) Width() int {
	return p.width
}

//Sets the width of the image.
func (p *FixedImageSize) SetWidth(value int) {
	p.width = value
}

// Gets the height of the image.
func (p *FixedImageSize) Height() int {
	return p.height
}

// Sets the height of the image.
func (p *FixedImageSize) SetHeight(value int) {
	p.height = value
}

// Gets the unit of measurement `ImageSizeUnit` for the width and height.
func (p *FixedImageSize) Unit() ImageSizeUnit {
	return p.unit
}

// Sets the unit of measurement `ImageSizeUnit` for the width and height.
func (p *FixedImageSize) SetUnit(value ImageSizeUnit) {
	p.unit = value
}
