package element

import (
	"encoding/json"

	"github.com/dynamicpdf-api/go-client/v2/color"

	"github.com/dynamicpdf-api/go-client/v2/font"
)

// Represents a page numbering label page element.
type PageNumbering struct {
	Element

	color color.Color
}

var _ ElementCollector = (*PageNumbering)(nil)
var _ json.Marshaler = (*PageNumbering)(nil)

/*
Initializes a new instance of the `PageNumberingElement` class.
  - @param {string} text Text to display in the label.
  - @param {ElementPlacement} placement The placement of the page numbering element on the page.
  - @param {float} xOffset X coordinate of the label.
  - @param {float} yOffset Y coordinate of the label.
*/
func NewPageNumberingElement(value string, placement elementPlacement, xOffset float32, yOffset float32) *PageNumbering {
	var p = PageNumbering{Element: *newElement(value, placement, xOffset, yOffset)}
	p.typeOfElement = p.ElementType()
	return &p
}
func (p *PageNumbering) ElementType() ElementType {
	return PageNumberingElement
}

// Gets `Color` object to use for the text of the label.
func (p *PageNumbering) Color() color.Color {
	return p.color
}

// sets `Color` object to use for the text of the label.
func (p *PageNumbering) SetColor(color color.Color) {
	p.color = color
	p.colorName = color.ColorString()
}

// Gets the `Font` object to use for the text of the label.
func (p *PageNumbering) Font() font.Font {
	return p.font
}

// sets the `Font` object to use for the text of the label.
func (p *PageNumbering) SetFont(font font.Font) {
	p.font = font
	p.fontName = font.Name
	p.resource = font.Resource
}

// Gets the text to display in the label.
func (p *PageNumbering) Text() string {
	return p.textValue
}

// sets the text to display in the label.
func (p *PageNumbering) SetText(value string) {
	p.textValue = value
}

// Gets the font size for the text of the label.
func (p *PageNumbering) FontSize() float32 {
	return p.fontSize
}

// Sets the font size for the text of the label.
func (p *PageNumbering) SetFontSize(value float32) {
	p.fontSize = value
}
