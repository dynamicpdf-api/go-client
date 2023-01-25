package element

/**
 * Represents Pdf417 barcode element.
 * This class can be used to generate Pdf417 barcode symbol.
 */
type Pdf417Barcode struct {
	Dim2
}

var _ ElementCollector = (*Pdf417Barcode)(nil)

/**
 * Initializes a new instance of the `Pdf417BarcodeElement` class.
 * @param {string} value String to be encoded | byte array.
 * @param {ElementPlacement} placement The placement of the barcode on the page.
 * @param {int} columns Columns of the PDF417 barcode.
 * @param {float} xOffset The X coordinate of the PDF417 barcode.
 * @param {float} yOffset The Y coordinate of the PDF417 barcode.
 */
func NewPdf417Barcode(value string, placement elementPlacement, columns int, xOffset float32, yOffset float32) *Pdf417Barcode {
	var p = Pdf417Barcode{Dim2: *NewDim2BarcodeWithStringValue(value, placement, xOffset, yOffset)}
	p.typeOfElement = p.ElementType()
	p.SetColumns(columns)
	return &p
}

/**
 * Initializes a new instance of the `Pdf417BarcodeElement` class.
 * @param {Byte[]} value String to be encoded | byte array.
 * @param {ElementPlacement} placement The placement of the barcode on the page.
 * @param {int} columns Columns of the PDF417 barcode.
 * @param {float} xOffset The X coordinate of the PDF417 barcode.
 * @param {float} yOffset The Y coordinate of the PDF417 barcode.
 */
func NewPdf417BarcodeWithByteValue(value []byte, placement elementPlacement, columns int, xOffset float32, yOffset float32) *Pdf417Barcode {
	var p = Pdf417Barcode{Dim2: *NewDim2BarcodeWithByteValue(value, placement, xOffset, yOffset)}
	p.typeOfElement = p.ElementType()
	p.SetColumns(columns)
	return &p
}

/** Gets the columns of the barcode. */
func (p *Pdf417Barcode) Columns() int {
	return p.columns
}

/*Sets the columns of the barcode. */
func (p *Pdf417Barcode) SetColumns(value int) {
	p.columns = value
}

/** Gets a boolean indicating whether to process the tilde character. */
func (a *Pdf417Barcode) ProcessTilde() bool {
	return a.processTilde
}

/** Sets a boolean indicating whether to process the tilde character. */
func (a *Pdf417Barcode) SetProcessTilde(value bool) {
	a.processTilde = value
}

/** Gets the YDimension of the barcode. */
func (a *Pdf417Barcode) YDimension() float32 {
	return a.yDimension
}

/** Sets the YDimension of the barcode. */
func (a *Pdf417Barcode) SetYDimension(value float32) {
	a.yDimension = value
}

/** Gets the Compact Pdf417. */
func (a *Pdf417Barcode) CompactPdf417() bool {
	return a.compactPdf417
}

/** Sets the Compact Pdf417. */
func (a *Pdf417Barcode) SetCompactPdf417(value bool) {
	a.compactPdf417 = value
}

/** Gets the type of compaction. */
func (a *Pdf417Barcode) Compaction() compaction {
	return a.compaction
}

/** sets the type of compaction. */
func (a *Pdf417Barcode) SetCompaction(value compaction) {
	a.compaction = value
}

/** Gets or sets the error correction level for the PDF417 barcode. */
func (a *Pdf417Barcode) ErrorCorrection() ErrorCorrection {
	return a.errorCorrection
}

/** sets the error correction level for the PDF417 barcode. */
func (a *Pdf417Barcode) SetErrorCorrection(value ErrorCorrection) {
	a.errorCorrection = value
}
func (p *Pdf417Barcode) ElementType() ElementType {
	return Pdf417BarcodeElement
}
