package imaging

// Base class for BMP color formats
type BmpColorFormat struct {
	ColorFormat
}

// Creates BmpColorFormat object with the given type.
func NewBmpColorFormat(colorFormatType ColorFormatType) *BmpColorFormat {
	var ep BmpColorFormat
	if colorFormatType != Monochrome {
		ep.Type = Rgb
	} else {
		ep.Type = colorFormatType
	}
	return &ep
}
