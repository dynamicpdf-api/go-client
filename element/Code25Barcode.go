package element

/* Represents a Code 2 of 5 barcode element.
This class can be used to place a Code 2 of 5 barcode on a page.
*/
type Code25Barcode struct {
	TextBarcode
}

var _ ElementCollector = (*Code25Barcode)(nil)

/*
  Initializes a new instance of the `Code25BarcodeElement` class.
   * @param {string} value. The value of the barcode.
   * @param {ElementPlacement} placement. The placement of the barcode on the page.
   * @param {number} height. The height of the barcode.
   * @param {number} xOffset. The X coordinate of the barcode.
   * @param {number} yOffset. The Y coordinate of the barcode.
*/
func NewCode25Barcode(value string, placement elementPlacement, height float32, xOffset float32, yOffset float32) *Code25Barcode {
	var p = Code25Barcode{TextBarcode: *NewTextBarcode(value, placement, xOffset, yOffset)}
	p.typeOfElement = p.ElementType()
	p.SetHeight(height)
	return &p
}

// Gets the height of the barcode.
func (p *Code25Barcode) Height() float32 {
	return p.height
}

// Sets the height of the barcode.
func (p *Code25Barcode) SetHeight(value float32) {
	p.height = value
}

func (p *Code25Barcode) ElementType() ElementType {
	return Code25BarcodeElement
}
