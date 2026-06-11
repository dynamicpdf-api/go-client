package element

import "encoding/base64"

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
	var p Dim2
	p.valueType = Base64EncodedBytes
	p.value = base64.StdEncoding.EncodeToString(value)
	p.Placement = placement
	p.XOffset = xOffset
	p.YOffset = yOffset
	return &p
}
