package element

/** Represents a MSI Barcode element (also known as Modified Plessey). */
type MsiBarcode struct {
	TextBarcode
}

var _ ElementCollector = (*MsiBarcode)(nil)

/**
 * Initializes a new instance of the `MsiBarcodeElement` class.
 * @param {string} value The value of the barcode.
 * @param {ElementPlacement} placement The placement of the barcode on the page.
 * @param {float} height The height of the barcode.
 * @param {float} xOffset The X coordinate of the barcode.
 * @param {float} yOffset The Y coordinate of the barcode.
 */
func NewMsiBarcode(value string, placement elementPlacement, height float32, xOffset float32, yOffset float32) *MsiBarcode {
	var p = MsiBarcode{TextBarcode: *NewTextBarcode(value, placement, xOffset, yOffset)}
	p.typeOfElement = p.ElementType()
	p.SetHeight(height)
	return &p
}

/** Gets the height of the barcode. */
func (p *MsiBarcode) Height() float32 {
	return p.height
}

/**Sets the height of the barcode. */
func (p *MsiBarcode) SetHeight(value float32) {
	p.height = value
}

func (p *MsiBarcode) ElementType() ElementType {
	return MsiBarcodeElement
}
