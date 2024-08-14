package imaging

// Base class for Png color formats, used for RGB, RGBA, and Grayscale color formats.
type PngColorFormat struct {
	ColorFormat
}

/*
  Initializes a new instance of the PngColorFormat class.
   * @param {ColorFormatType} colorFormatType. The value of the barcode.
*/
func NewPngColorFormat(colorFormatType ColorFormatType) *PngColorFormat {
	var ep PngColorFormat
	ep.Type = colorFormatType
	return &ep
}
