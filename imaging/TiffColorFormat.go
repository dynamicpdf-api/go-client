package imaging

// Base class for Tiff color formats, used for RGB, RGBA, and Grayscale color formats.
type TiffColorFormat struct {
	ColorFormat
}

/*
  Initializes a new instance of the TiffColorFormat class.
   * @param {ColorFormatType} colorFormatType. The value of the barcode.
*/
func NewTiffColorFormat(colorFormatType ColorFormatType) *TiffColorFormat {
	var ep TiffColorFormat
	ep.Type = colorFormatType
	return &ep
}
