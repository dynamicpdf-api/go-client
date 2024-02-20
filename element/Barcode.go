package element

import (
	"github.com/dynamicpdf-api/go-client/v2/color"

	"encoding/json"
)

// Base class from which barcode page elements are derived.
type Barcode struct {
	Element
	color color.Color
}

var _ ElementCollector = (*Barcode)(nil)
var _ json.Marshaler = (*Barcode)(nil)

func NewBarcodeWithStringValue(value string, placement elementPlacement, xOffset float32, yOffset float32) *Barcode {
	var p = Barcode{Element: *newElement(value, placement, xOffset, yOffset)}
	p.value = value
	return &p
}

// gets the Color of the barcode.
func (p *Barcode) Color() color.Color {
	return p.color
}

// sets the Color of the barcode.
func (p *Barcode) SetColor(value color.Color) {
	p.color = value
	p.colorName = p.color.ColorString()
}

// Gets the XDimension of the barcode.
func (p *Barcode) XDimension() float32 {
	return p.xDimension
}

// Sets the XDimension of the barcode.
func (p *Barcode) SetXDimension(value float32) {
	p.xDimension = value
}

// Gets the value of the barcode.
func (p *Barcode) Value() string {
	return p.textValue
}

// sets the value of the barcode.
func (p *Barcode) SetValue(value string) {
	p.textValue = value
}
