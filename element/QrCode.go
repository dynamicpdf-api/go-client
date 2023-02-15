package element

// Represents a QR code barcode element.
type QrCode struct {
	Dim2
}

var _ ElementCollector = (*QrCode)(nil)

/*
  Initializes a new instance of the `QrCodeElement` class.
   * @param {string} value The value of the QR code.
   * @param {ElementPlacement} placement The placement of the barcode on the page.
   * @param {float} xOffset The X coordinate of the QR code.
   * @param {float} yOffset The Y coordinate of the QR code.
*/
func NewQrCodeElement(value string, placement elementPlacement, xOffset float32, yOffset float32) *QrCode {
	var p = QrCode{Dim2: *NewDim2BarcodeWithStringValue(value, placement, xOffset, yOffset)}
	p.typeOfElement = p.ElementType()
	return &p
}

/*
  Initializes a new instance of the `QrCodeElement` class.
   * @param {byte[]} value The value of the QR code.
   * @param {ElementPlacement} placement The placement of the barcode on the page.
   * @param {float} xOffset The X coordinate of the QR code.
   * @param {float} yOffset The Y coordinate of the QR code.
*/
func NewQrCodeElementWithByteArray(value []byte, placement elementPlacement, xOffset float32, yOffset float32) *QrCode {
	var p = QrCode{Dim2: *NewDim2BarcodeWithByteValue(value, placement, xOffset, yOffset)}
	p.typeOfElement = p.ElementType()
	return &p
}

// Gets the QR code version.
func (a *QrCode) Fnc1() QrCodeFnc1 {
	return a.fnc1
}

// Sets the QR code version.
func (a *QrCode) SetFnc1(value QrCodeFnc1) {
	a.fnc1 = value
}

// Gets or sets the QR code version.
func (a *QrCode) Version() int {
	return a.version
}

// Gets or sets the QR code version.
func (a *QrCode) SetVersion(value int) {
	a.version = value
}

func (p *QrCode) ElementType() ElementType {
	return QrCodeElement
}
