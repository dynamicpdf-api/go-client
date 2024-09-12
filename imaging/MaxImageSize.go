package imaging

// MaxImageSize represents an image size that fits within a specified maximum width and height.
type MaxImageSize struct {
	ImageSize
}

// Initializes a new instance of the MaxImageSize class.
func NewMaxImageSize() *MaxImageSize {
	var ep MaxImageSize
	ep.Type = Max
	return &ep
}

// GetMaxWidth returns the maximum width of the image.
func (m *MaxImageSize) GetMaxWidth() int {
	return m.maxWidth
}

// SetMaxWidth sets the maximum width of the image.
func (m *MaxImageSize) SetMaxWidth(value int) {
	m.maxWidth = value
}

// GetMaxHeight returns the maximum height of the image.
func (m *MaxImageSize) GetMaxHeight() int {
	return m.maxHeight
}

// SetMaxHeight sets the maximum height of the image.
func (m *MaxImageSize) SetMaxHeight(value int) {
	m.maxHeight = value
}

// GetUnit returns the unit of measurement `ImageSizeUnit` for the maximum width and height.
func (m *MaxImageSize) GetUnit() ImageSizeUnit {
	return m.unit
}

// SetUnit sets the unit of measurement `ImageSizeUnit` for the maximum width and height.
func (m *MaxImageSize) SetUnit(value ImageSizeUnit) {
	m.unit = value
}
