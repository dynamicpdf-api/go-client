package element

import (
	"github.com/dynamicpdf-api/go-client/color"
)

/** Represents a rectangle page element.
 * This class can be used to place rectangles of any size or color on a page.
 */
type Rectangle struct {
	Element

	fillColor color.Color

	borderColor color.Color

	borderStyle LineStyle
}

var _ ElementCollector = (*Rectangle)(nil)

func NewRectangle(placement elementPlacement, width float32, height float32) *Rectangle {
	var p Rectangle
	p.Placement = placement
	p.SetWidth(width)
	p.SetHeight(height)
	p.typeOfElement = p.ElementType()
	return &p
}

func (p *Rectangle) ElementType() ElementType {
	return RectangleElement
}

/** Gets the height of the rectangle. */
func (p *Rectangle) Height() float32 {
	return p.height
}

/** Sets the height of the rectangle. */
func (p *Rectangle) SetHeight(value float32) {
	p.height = value
}

/** Gets the width of the rectangle. */
func (p *Rectangle) Width() float32 {
	return p.width
}

/** Gets the border width of the rectangle. */
func (p *Rectangle) BorderWidth() float32 {
	return p.borderWidth
}

/** sets the border width of the rectangle. */
func (p *Rectangle) SetBorderWidth(value float32) {
	p.borderWidth = value
}

/** Sets the width of the rectangle. */
func (p *Rectangle) SetWidth(value float32) {
	p.width = value
}

/** Gets the corner radius of the rectangle. */
func (p *Rectangle) CornerRadius() float32 {
	return p.cornerRadius
}

/** Sets the corner radius of the rectangle. */
func (p *Rectangle) SetCornerRadius(value float32) {
	p.cornerRadius = value
}

/**
 * Gets the `Color` object to use for the fill of the rectangle.
 */
func (p *Rectangle) FillColor() color.Color {
	return p.fillColor
}

/**
 * sets the `Color` object to use for the fill of the rectangle.
 */
func (p *Rectangle) SetFillColor(fillColor color.Color) {
	p.fillColor = fillColor
	p.fillColorName = fillColor.ColorString()
}

/**
 * Gets the `Color` object to use for the border of the rectangle.
 */
func (p *Rectangle) BorderColor() color.Color {
	return p.borderColor
}

/** sets the `Color` object to use for the border of the rectangle. */
func (p *Rectangle) SetBorderColor(borderColor color.Color) {
	p.borderColor = borderColor
	p.borderColorName = borderColor.ColorString()
}

/** Gets the `BorderStyle` object used to specify the border style of the rectangle. */
func (p *Rectangle) BorderStyle() LineStyle {
	return p.borderStyle
}

/**sets the `BorderStyle` object used to specify the border style of the rectangle. */
func (p *Rectangle) SetBorderStyle(style LineStyle) {
	p.borderStyle = style
	p.borderStyleName = style.lineStyleString
}
