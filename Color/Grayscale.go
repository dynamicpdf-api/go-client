package color

import (
	"strconv"
)

// <summary>
// Represents a grayscale color.
// </summary>
type Grayscalecolor struct {
	Color
	grayLevel float64
}

func newGrayscalecolor(color string) *Grayscalecolor {
	var ep Grayscalecolor
	ep.colorString = color
	return &ep
}

/// <summary>
/// Initializes a new instance of the  <see cref="Grayscale"/> class.
/// </summary>
/// <param name="grayLevel">The gray level for the color.</param>

func NewGrayscalecolor(greyLevel float64) *Grayscalecolor {
	var ep Grayscalecolor
	ep.grayLevel = greyLevel
	ep.colorString = ep.ColorString()
	return &ep
}

// / <summary>Gets the color black.</summary>
func (p *Grayscalecolor) Black() string {
	return p.colorString
}

// / <summary>Gets the color white.</summary>
func (p *Grayscalecolor) White() string {
	return p.colorString
}

func (p *Grayscalecolor) ColorString() string {

	if p.colorString != "" {
		return p.colorString
	} else {
		return "gray(" + strconv.FormatFloat(p.grayLevel, 'f', -1, 64) + ")"
	}
}

func (p *Grayscalecolor) SetColorString(value string) {
	p.colorString = value
}

func (p *Grayscalecolor) GrayLevel() float64 {
	return p.grayLevel
}

func (p *Grayscalecolor) SetGrayLevel(grayLevel float64) {
	p.grayLevel = grayLevel
}
