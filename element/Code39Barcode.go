package element

/*
 Represents a Code 3 of 9 barcode element.
This class can be used to place a Code 3 of 9 barcode on a page.
*/
type Code39Barcode struct {
	TextBarcode
}

var _ ElementCollector = (*Code39Barcode)(nil)

func NewCode39Barcode(value string, placement elementPlacement, height float32, xOffset float32, yOffset float32) *Code39Barcode {
	var p = Code39Barcode{TextBarcode: *NewTextBarcode(value, placement, xOffset, yOffset)}
	p.typeOfElement = p.ElementType()
	p.SetHeight(height)
	return &p
}

// Gets the height of the barcode.
func (p *Code39Barcode) Height() float32 {
	return p.height
}

// sets the height of the barcode.
func (p *Code39Barcode) SetHeight(value float32) {
	p.height = value
}

func (p *Code39Barcode) ElementType() ElementType {
	return Code39BarcodeElement
}
