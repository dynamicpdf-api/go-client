package element

import (
	"encoding/json"

	"github.com/dynamicpdf-api/go-client/color"
)

/*
	Represents a line page element.

This class can be used to place lines of different length, width, color and patterns on a page.
*/
type Line struct {
	Element

	lineStyle LineStyle

	color color.Color
}

var _ ElementCollector = (*Line)(nil)
var _ json.Marshaler = (*Line)(nil)

/*
Initializes a new instance of the `LineElement` class.
  - @param {ElementPlacement} placement The placement of the line on the page.
  - @param {float} x2Offset X2 coordinate of the line.
  - @param {float} y2Offset Y2 coordinate of the line.
*/
func NewLine(placement elementPlacement, x2OffSet float32, y2OffSet float32) *Line {
	var p Line
	p.typeOfElement = p.ElementType()
	p.Placement = placement
	p.SetX2Offset(x2OffSet)
	p.SetY2Offset(y2OffSet)
	p.SetLineStyle(*p.lineStyle.Dash())
	return &p
}

func (p *Line) ElementType() ElementType {
	return LineElement
}

func (p *Line) Color() color.Color {
	return p.color
}

func (p *Line) SetColor(value color.Color) {
	p.color = value
	p.colorName = value.ColorString()
}

func (p *Line) LineStyle() LineStyle {
	return p.lineStyle
}

func (p *Line) SetLineStyle(value LineStyle) {
	p.lineStyle = value
	p.lineStyleName = value.lineStyleString
}

// Gets the X2 coordinate of the line.
func (p *Line) X2Offset() float32 {
	return p.x2Offset
}

// Sets the X2 coordinate of the line.
func (p *Line) SetX2Offset(value float32) {
	p.x2Offset = value
}

// Sets the X2 coordinate of the line.
func (p *Line) Y2Offset() float32 {
	return p.y2Offset
}

// sets the X2 coordinate of the line.
func (p *Line) SetY2Offset(value float32) {
	p.y2Offset = value
}

// Gets the width of the line.
func (p *Line) Width() float32 {
	return p.width
}

// Sets the width of the line.
func (p *Line) SetWidth(value float32) {
	p.width = value
}
