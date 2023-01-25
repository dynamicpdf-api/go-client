package element

/** Represents an IATA 2 of 5 barcode element.
 * This class can be used to place an IATA 2 of 5 barcode on a page.
 */
type Iata25Barcode struct {
	TextBarcode
}

var _ ElementCollector = (*Iata25Barcode)(nil)

/**
 * Initializes a new instance of the `Iata25BarcodeElement` class.
 * @param {string} value The value of the barcode.
 * @param {ElementPlacement} placement The placement of the barcode on the page.
 * @param {float} height The height of the barcode.
 * @param {float} xOffset The X coordinate of the barcode.
 * @param {float} yOffset The Y coordinate of the barcode.
 */
func NewIata25Barcode(value string, placement elementPlacement, height float32, xOffset float32, yOffset float32) *Iata25Barcode {
	var p = Iata25Barcode{TextBarcode: *NewTextBarcode(value, placement, xOffset, yOffset)}
	p.typeOfElement = p.ElementType()
	p.SetHeight(height)
	return &p
}

/** Gets the height of the barcode. */
func (p *Iata25Barcode) Height() float32 {
	return p.height
}

/** Sets the height of the barcode. */
func (p *Iata25Barcode) SetHeight(value float32) {
	p.height = value
}

/** Gets a value indicating if the check digit should be added to the value. */
func (p *Iata25Barcode) IncludeCheckDigit() bool {
	return p.includeCheckDigit
}

/** Sets a value indicating if the check digit should be added to the value. */
func (p *Iata25Barcode) SetIncludeCheckDigit(value bool) {
	p.includeCheckDigit = value
}
func (p *Iata25Barcode) ElementType() ElementType {
	return Iata25BarcodeElement
}
