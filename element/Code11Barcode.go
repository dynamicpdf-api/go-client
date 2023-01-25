package element

/**Represents a Code 11 barcode element.
 * This class can be used to place a Code 11 barcode on a page.
 */
type Code11Barcode struct {
	TextBarcode
}

var _ ElementCollector = (*Code11Barcode)(nil)

/**
 * Initializes a new instance of the `Code11BarcodeElement` class.
 * @param {string} value The value of the barcode.
 * @param {ElementPlacement} placement The placement of the barcode on the page.
 * @param {number} height The height of the barcode.
 * @param {number} xOffset The X coordinate of the barcode.
 * @param {number} yOffset The Y coordinate of the barcode.
 */
func NewCode11Barcode(value string, placement elementPlacement, height float32, xOffset float32, yOffset float32) *Code11Barcode {
	var p = Code11Barcode{TextBarcode: *NewTextBarcode(value, placement, xOffset, yOffset)}
	p.typeOfElement = p.ElementType()
	p.SetHeight(height)
	return &p
}

/** Gets or sets the height of the barcode. */
func (p *Code11Barcode) Height() float32 {
	return p.height
}

/** Sets or sets the height of the barcode. */
func (p *Code11Barcode) SetHeight(value float32) {
	p.height = value
}
func (p *Code11Barcode) ElementType() ElementType {
	return Code11BarcodeElement
}
