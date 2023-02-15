package element

import (
	"encoding/json"

	"github.com/dynamicpdf-api/go-client/color"
	"github.com/dynamicpdf-api/go-client/font"
)

// Base class from which barcode page elements that display text are derived.
type TextBarcode struct {
	Barcode

	font font.Font

	textColor color.Color
}

var _ ElementCollector = (*TextBarcode)(nil)
var _ json.Marshaler = (*TextBarcode)(nil)

func NewTextBarcode(value string, placement elementPlacement, xOffset float32, yOffset float32) *TextBarcode {
	var p = TextBarcode{Barcode: *NewBarcodeWithStringValue(value, placement, xOffset, yOffset)}
	p.Placement = TopLeft
	p.typeOfElement = p.ElementType()
	return &p
}

func (p *TextBarcode) ElementType() ElementType {
	return TextElement
}

// Gets a value indicating if the value should be placed as text below the barcode.
func (p *TextBarcode) ShowText() bool {
	return p.showText
}

// Sets a value indicating if the value should be placed as text below the barcode.
func (p *TextBarcode) SetShowText(value bool) {
	p.showText = value
}

// Gets the font size to use when displaying the text.
func (p *TextBarcode) FontSize() float32 {
	return p.fontSize
}

// Sets the font size to use when displaying the text.
func (p *TextBarcode) SetFontSize(value float32) {
	p.fontSize = value
}
func (p *TextBarcode) Text() string {
	return p.textValue
}

func (p *TextBarcode) SetText(value string) {
	p.textValue = value
}

// Gets the color of the text.
func (p *TextBarcode) TextColor() color.Color {
	return p.textColor
}

// Sets the color of the text.
func (p *TextBarcode) SetTextColor(color color.Color) {
	p.textColor = color
	p.textColorName = color.ColorString()
}

// Gets  the font to use when displaying the text.
func (p *TextBarcode) Font() font.Font {
	return p.font
}

// Sets  the font to use when displaying the text.
func (p *TextBarcode) SetFont(font font.Font) {
	p.font = font
	p.fontName = font.Name
	p.resource = font.Resource
}
