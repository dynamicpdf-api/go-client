package imaging

// Represents an image size defined by Percentage scaling.
type PercentageImageSize struct {
	ImageSize
}

// Initializes a new instance of the `PercentageImageSize` class and sets the image size type to Percentage.
func NewPercentageImageSize() *PercentageImageSize {
	var ep PercentageImageSize
	ep.Type = Percentage
	return &ep
}

// GetType returns the type of the image size.
func (p *PercentageImageSize) GetType() ImageSizeType {
	return p.Type
}

// GetHorizontalPercentage returns the horizontal scaling percentage.
func (p *PercentageImageSize) GetHorizontalPercentage() int {
	return p.horizontalPercentage
}

// SetHorizontalPercentage sets the horizontal scaling percentage.
func (p *PercentageImageSize) SetHorizontalPercentage(value int) {
	p.horizontalPercentage = value
}

// GetVerticalPercentage returns the vertical scaling percentage.
func (p *PercentageImageSize) GetVerticalPercentage() int {
	return p.verticalPercentage
}

// SetVerticalPercentage sets the vertical scaling percentage.
func (p *PercentageImageSize) SetVerticalPercentage(value int) {
	p.verticalPercentage = value
}
