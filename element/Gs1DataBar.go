package element

/** Represents a GS1DataBar barcode element. */
type Gs1DataBar struct {
	TextBarcode
}

var _ ElementCollector = (*Gs1DataBar)(nil)

/**
 * Initializes a new instance of the `Gs1DataBarBarcodeElement` class.
 * @param {string} value The value of the barcode.
 * @param {ElementPlacement} placement The placement of the barcode on the page.
 * @param {float} height The height of the barcode.
 * @param {barType} type The GS1DataBarType of the barcode.
 * @param {float} xOffset The X coordinate of the barcode.
 * @param {float} yOffset The Y coordinate of the barcode.
 */
func NewGs1DataBarBarcode(value string, placement elementPlacement, height float32, bartype Gs1DataBarType, xOffset float32, yOffset float32) *Gs1DataBar {
	var p = Gs1DataBar{TextBarcode: *NewTextBarcode(value, placement, xOffset, yOffset)}
	p.typeOfElement = p.ElementType()
	p.SetHeight(height)
	p.gs1DataBarType = bartype
	return &p
}

/** Gets the height of the barcode. */
func (p *Gs1DataBar) Height() float32 {
	return p.height
}

/** Sets the height of the barcode. */
func (p *Gs1DataBar) SetHeight(value float32) {
	p.height = value
}

func (p *Gs1DataBar) ElementType() ElementType {
	return Gs1DataBarBarcodeElement
}
