package element

/**
 * Represents an Aztec barcode element.
 */
type AztecBarcode struct {
	Dim2
}

var _ ElementCollector = (*AztecBarcode)(nil)

/**
 * Initializes a new instance of the `AztecBarcodeElement` class.
 * @param {string} value The value of the barcode.
 * @param {ElementPlacement} placement The placement of the barcode on the page.
 * @param {float} xOffset The X coordinate of the barcode.
 * @param {float} yOffset The Y coordinate of the barcode.
 */
func NewAztecBarcodeElement(value string, placement elementPlacement, xOffset float32, yOffset float32) *AztecBarcode {
	var p = AztecBarcode{Dim2: *NewDim2BarcodeWithStringValue(value, placement, xOffset, yOffset)}
	p.typeOfElement = p.ElementType()
	return &p
}

/**
 * Initializes a new instance of the `AztecBarcodeElement` class.
 * @param {Buffer []} value The value of the barcode.
 * @param {ElementPlacement} placement The placement of the barcode on the page.
 * @param {float} xOffset The X coordinate of the barcode.
 * @param {float} yOffset The Y coordinate of the barcode.
 */
func NewAztecBarcodeElementWithByteValue(value []byte, placement elementPlacement, xOffset float32, yOffset float32) *AztecBarcode {
	var p AztecBarcode
	p.typeOfElement = p.ElementType()
	p = AztecBarcode{Dim2: *NewDim2BarcodeWithByteValue(value, placement, xOffset, yOffset)}
	return &p
}

/** Gets a boolean indicating whether to process tilde symbol in the input.
 */
func (a *AztecBarcode) ProcessTilde() bool {
	return a.processTilde
}

/** Sets a boolean indicating whether to process tilde symbol in the input.
* Setting <b>true</b> will check for ~ character and processes it for FNC1 or ECI characters.
 */
func (a *AztecBarcode) SetProcessTilde(value bool) {
	a.processTilde = value
}

/**Gets the barcode size, `AztecSymbolSize`. */
func (a *AztecBarcode) SymbolSize() int {
	return a.symbolSize
}

/**Sets the barcode size, `AztecSymbolSize`. */
func (a *AztecBarcode) SetSymbolSize(value int) {
	a.symbolSize = value
}

/** Gets the error correction value.
 * Error correction value may be between 5% to 95%.
 */
func (a *AztecBarcode) AztecErrorCorrection() int {
	return a.aztecErrorCorrection
}

/**Sets the error correction value.
 * Error correction value may be between 5% to 95%.
 */
func (a *AztecBarcode) SetAztecErrorCorrection(value int) {
	a.aztecErrorCorrection = value
}

/**
 * Gets a boolean representing if the barcode is a reader initialization symbol.
 * Setting <b>true</b> will mark the symbol as reader initialization symbol
 * and the size of the symbol should be one of the following, R15xC15 Compact, R19xC19, R23xC23, R27xC27, R31xC31, R37xC37, R41xC41, R45xC45, R49xC49, R53xC53, R57xC57, R61xC61, R67xC67, R71xC71, R75xC75,
 * R79xC79, R83xC83, R87xC87, R91xC91, R95xC95, R101xC101, R105xC105, R109xC109, however it is recommended to set Auto.
 */
func (a *AztecBarcode) ReaderInitialization() bool {
	return a.readerInitialization
}

/**
 * Sets a boolean representing if the barcode is a reader initialization symbol.
 * Setting <b>true</b> will mark the symbol as reader initialization symbol
 * and the size of the symbol should be one of the following, R15xC15 Compact, R19xC19, R23xC23, R27xC27, R31xC31, R37xC37, R41xC41, R45xC45, R49xC49, R53xC53, R57xC57, R61xC61, R67xC67, R71xC71, R75xC75,
 * R79xC79, R83xC83, R87xC87, R91xC91, R95xC95, R101xC101, R105xC105, R109xC109, however it is recommended to set Auto.
 */
func (a *AztecBarcode) SetReaderInitialization(value bool) {
	a.readerInitialization = value
}
func (p *AztecBarcode) ElementType() ElementType {
	return AztecBarcodeElement
}
