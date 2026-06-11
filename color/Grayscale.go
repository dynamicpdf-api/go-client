package color

import (
	"strconv"
)

// Represents a grayscale color.
type Grayscalecolor struct {
	Color
	grayLevel float64
}

/*
Initializes a new instance of the 'Grayscalecolor' class.
  - @param {string} color. The color of the Grayscalecolor.
*/
func newGrayscalecolor(color string) *Grayscalecolor {
	var ep Grayscalecolor
	ep.colorString = color
	return &ep
}

/*
Initializes a new instance of the 'Grayscale' class.
  - @param {float} grayLevel. The gray level for the color.
*/
func NewGrayscalecolor(greyLevel float64) *Grayscalecolor {
	var ep Grayscalecolor
	ep.grayLevel = greyLevel
	ep.colorString = ep.ColorString()
	return &ep
}

// Gets the color black.
func (p *Grayscalecolor) Black() string {
	return p.colorString
}

// Gets the color white.
func (p *Grayscalecolor) White() string {
	return p.colorString
}

/*
Gets the color string
  - @return The color string
*/
func (p *Grayscalecolor) ColorString() string {

	if p.colorString != "" {
		return p.colorString
	} else {
		return "gray(" + strconv.FormatFloat(p.grayLevel, 'f', -1, 64) + ")"
	}
}

/*
Sets the color string
  - @param value The color string
*/
func (p *Grayscalecolor) SetColorString(value string) {
	p.colorString = value
}

func (p *Grayscalecolor) GrayLevel() float64 {
	return p.grayLevel
}

func (p *Grayscalecolor) SetGrayLevel(grayLevel float64) {
	p.grayLevel = grayLevel
}
