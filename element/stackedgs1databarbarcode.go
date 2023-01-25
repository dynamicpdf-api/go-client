package element

/** Represents a StackedGS1DataBar barcode element. */
type StackedGs1DataBarBarcode struct {
	TextBarcode
}

var _ ElementCollector = (*StackedGs1DataBarBarcode)(nil)

/**
 * Initializes a new instance of the `StackedGs1DataBarBarcodeElement` class.
 * @param {string} value The value of the barcode.
 * @param {ElementPlacement} placement The placement of the barcode on the page.
 * @param {StackedGs1DataBarType} stackedGs1DataBarType The StackedGS1DataBarType of the barcode.
 * @param {float} rowHeight The row height of the barcode.
 * @param {float} xOffset The X coordinate of the barcode.
 * @param {float} yOffset The Y coordinate of the barcode.
 */
func NewStackedGs1DataBarBarcode(value string, placement elementPlacement, stackedGs1BarType stackedGs1DataBarType, rowHeight float32, xOffset float32, yOffset float32) *StackedGs1DataBarBarcode {
	var p = StackedGs1DataBarBarcode{TextBarcode: *NewTextBarcode(value, placement, xOffset, yOffset)}
	p.typeOfElement = p.ElementType()
	p.stackedGs1DataBarType = stackedGs1BarType
	p.rowHeight = rowHeight
	return &p
}

/** Gets the segment count of the Expanded Stacked barcode. */

func (p *StackedGs1DataBarBarcode) ExpandedStackedSegmentCount() int {
	return p.expandedStackedSegmentCount
}

/** Sets the segment count of the Expanded Stacked barcode. */

func (p *StackedGs1DataBarBarcode) SetExpandedStackedSegmentCount(value int) {
	p.expandedStackedSegmentCount = value
}

/** Gets the row height of the barcode. */
func (p *StackedGs1DataBarBarcode) RowHeight() float32 {
	return p.rowHeight
}

/** Sets the row height of the barcode. */
func (p *StackedGs1DataBarBarcode) SetRowHeight() float32 {
	return float32(p.rowHeight)
}

func (p *StackedGs1DataBarBarcode) ElementType() ElementType {
	return StackedGs1DataBarBarcodeElement
}
