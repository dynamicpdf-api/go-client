package element

/*
  Represents a Code 128 barcode element.
 This class can be used to place a Code 128 barcode on a page.
*/
type Code128Barcode struct {
	TextBarcode
}

var _ ElementCollector = (*Code128Barcode)(nil)

/*
   Initializes a new instance of the `Code128BarcodeElement` class.
Code sets can be specified along with data, in order to do this `ProcessTilde` property needs to be set to `true`
Example value: "~ BHello ~ AWORLD 1~ C2345", where ~A, ~B and ~C representing code sets A, B and C respectively.
However if any inline code set has invalid characters it will be shifted to an appropriate code set.
    * @param {string} value The value of the barcode.
    * @param {elementPlacement} placement The placement of the barcode on the page.
    * @param {number} height The height of the barcode.
    * @param {number} xOffset The X coordinate of the barcode.
    * @param {number} yOffset The Y coordinate of the barcode.
*/
func NewCode128Barcode(value string, placement elementPlacement, height float32, xOffset float32, yOffset float32) *Code128Barcode {
	var p = Code128Barcode{TextBarcode: *NewTextBarcode(value, placement, xOffset, yOffset)}
	p.typeOfElement = p.ElementType()
	p.SetHeight(height)
	return &p
}

// Gets the height of the barcode.
func (p *Code128Barcode) Height() float32 {
	return p.height
}

// sets the height of the barcode.
func (p *Code128Barcode) SetHeight(value float32) {
	p.height = value
}

func (p *Code128Barcode) ElementType() ElementType {
	return Code128BarcodeElement
}
