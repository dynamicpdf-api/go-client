package color

import (
	"strconv"
)

// Represents a CMYK color.
type CmykColor struct {
	Color
	cyan    float64
	magenta float64
	yellow  float64
	black   float64
}

/*
Initializes a new instance of the 'CmykColor'
  - @param {float} cyan. The cyan intensity.
  - @param {float} magenta. The magenta intensity.
  - @param {float} yellow. The yellow intensity.
  - @param {float} black. The black intensity.
*/
func NewCmykColor(cyan float64, magenta float64, yellow float64, black float64) *CmykColor {
	var p CmykColor
	p.cyan = cyan
	p.magenta = magenta
	p.yellow = yellow
	p.black = black
	p.colorString = p.ColorString()
	return &p
}

// Gets the color black.
func (p *CmykColor) Black() *CmykColor {
	n := NewCmykColor(1, 1, 1, 1)
	return n
}

// Gets the color white.
func (p *CmykColor) White() *CmykColor {
	n := NewCmykColor(0, 0, 0, 0)
	return n
}

func (p *CmykColor) ColorString() string {
	if p.colorString != "" {
		return p.colorString
	} else {
		return "cmyk(" + strconv.FormatFloat(p.cyan, 'f', -1, 64) + "," + strconv.FormatFloat(p.magenta, 'f', -1, 64) + "," + strconv.FormatFloat(p.yellow, 'f', -1, 64) + "," + strconv.FormatFloat(p.black, 'f', -1, 64) + ")"
	}
}

func (p *CmykColor) SetColorString(value string) {
	p.colorString = value
}
