package element

type Code93Barcode struct {
	TextBarcode
}

var _ ElementCollector = (*Code93Barcode)(nil)

/**
 * Initializes a new instance of the `Code93BarcodeElement` class.
 * @param {string} value The value of the barcode.
 * @param {ElementPlacement} placement The placement of the barcode on the page.
 * @param {number} height The height of the barcode.
 * @param {number} xOffset The X coordinate of the barcode.
 * @param {number} yOffset The Y coordinate of the barcode.
 */
func NewCode93Barcode(value string, placement elementPlacement, height float32, xOffset float32, yOffset float32) *Code93Barcode {
	var p = Code93Barcode{TextBarcode: *NewTextBarcode(value, placement, xOffset, yOffset)}
	p.typeOfElement = p.ElementType()
	p.SetHeight(height)
	return &p
}

/** Gets the height of the barcode. */
func (p *Code93Barcode) Height() float32 {
	return p.height
}

/**sets the height of the barcode. */
func (p *Code93Barcode) SetHeight(value float32) {
	p.height = value
}

func (p *Code93Barcode) ElementType() ElementType {
	return Code93BarcodeElement
}
