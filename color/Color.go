package color

// Base class representing a color.
type Color struct {
	colorString string
}

func NewColor() *Color {
	var p Color
	return &p
}

func (p *Color) ColorString() string {
	return p.colorString
}

func (p *Color) SetColorString(value string) {
	p.colorString = value
}
