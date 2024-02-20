package element

import (
	"encoding/json"

	"github.com/dynamicpdf-api/go-client/v2/color"
	"github.com/dynamicpdf-api/go-client/v2/font"
)

type Text struct {
	Element
	font  font.Font
	color color.Color
}

var _ ElementCollector = (*Text)(nil)
var _ json.Marshaler = (*Text)(nil)

func NewText(value string, placement elementPlacement, xOffset float32, yOffset float32) *Text {
	var p = Text{Element: *newElement(value, placement, xOffset, yOffset)}
	p.typeOfElement = p.ElementType()
	return &p
}

func (p *Text) ElementType() ElementType {
	return TextElement
}

// Gets the font size for the text of the text element.
func (p *Text) FontSize() float32 {
	return p.fontSize
}

// Sets the font size for the text of the text element.
func (p *Text) SetFontSize(value float32) {
	p.fontSize = value
}

func (p *Text) Text() string {
	return p.textValue
}

func (p *Text) SetText(value string) {
	p.textValue = value
}

// Gets `Color` object to use for the text of the text element.
func (p *Text) Color() color.Color {
	return p.color
}

// Sets `Color` object to use for the text of the text element.
func (p *Text) SetColor(color color.Color) {
	p.color = color
	p.colorName = color.ColorString()
}

// Gets the `Font` object used to specify the font of the text for the text element.
func (p *Text) Font() font.Font {
	return p.font
}

// Sets the `Font` object used to specify the font of the text for the text element.
func (p *Text) SetFont(font font.Font) {
	p.font = font
	p.fontName = font.Name
	p.resource = font.Resource
}
