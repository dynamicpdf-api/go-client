package element

type DataMatrix struct {
	Dim2
}

/**
 * Initializes a new instance of the `DataMatrixBarcodeElement` class.
 * @param {string} value The value of the barcode
 * @param {ElementPlacement} placement The placement of the barcode on the page
 * @param {number} xOffset The X coordinate of the barcode
 * @param {number} yOffset The Y coordinate of the barcode.
 * @param {number} symbolsize The symbol size of the barcode.
 * @param {number} encodingType The encoding type of the barcode.
 * @param {number} functionCharacter The function character of the barcode.
 */
func NewDataMatrix(value string, placement elementPlacement, xOffset float32, yOffset float32, symbolSize DataMatrixSymbolSize, encodingType DataMatrixEncodingType, functionCharacter DataMatrixFunctionCharacter) *DataMatrix {
	var p = DataMatrix{Dim2: *NewDim2BarcodeWithStringValue(value, placement, xOffset, yOffset)}
	p.typeOfElement = p.ElementType()
	p.dataMatrixSymbolSize = symbolSize
	p.dataMatrixFunctionCharacter = functionCharacter
	p.dataMatrixEncodingType = encodingType
	return &p
}

/**
 * Initializes a new instance of the `DataMatrixBarcodeElement` class.
 * @param {Byte[]} value The value of the barcode
 * @param {ElementPlacement} placement The placement of the barcode on the page
 * @param {number} xOffset The X coordinate of the barcode
 * @param {number} yOffset The Y coordinate of the barcode.
 * @param {number} symbolsize The symbol size of the barcode.
 * @param {number} encodingType The encoding type of the barcode.
 * @param {number} functionCharacter The function character of the barcode.
 */
func NewDataMatrixWithByteValue(value []byte, placement elementPlacement, xOffset float32, yOffset float32, symbolSize DataMatrixSymbolSize, encodingType DataMatrixEncodingType, functionCharacter DataMatrixFunctionCharacter) *DataMatrix {
	var p = DataMatrix{Dim2: *NewDim2BarcodeWithByteValue(value, placement, xOffset, yOffset)}
	p.typeOfElement = p.ElementType()
	p.dataMatrixSymbolSize = symbolSize
	p.dataMatrixFunctionCharacter = functionCharacter
	p.dataMatrixEncodingType = encodingType
	return &p
}

func (p *DataMatrix) ElementType() ElementType {
	return DataMatrixBarcodeElement
}
