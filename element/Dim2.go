package element

// The base class for 2 dimensional bar codes (Aztec, Pdf417, DataMatrixBarcode and QrCode).
type Dim2 struct {
	Barcode
}

var _ ElementCollector = (*Dim2)(nil)

func NewDim2BarcodeWithStringValue(value string, placement elementPlacement, xOffset float32, yOffset float32) *Dim2 {
	var p = Dim2{Barcode: *NewBarcodeWithStringValue(value, placement, xOffset, yOffset)}
	return &p
}

func NewDim2BarcodeWithByteValue(value []byte, placement elementPlacement, xOffset float32, yOffset float32) *Dim2 {
	var p = Dim2{Barcode: *NewBarcodeWithStringValue(string(value), placement, xOffset, yOffset)}
	p.valueType = Base64EncodedBytes
	return &p
}
