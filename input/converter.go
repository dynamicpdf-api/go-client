package input

// Represents the base class for inputs.
type Converter struct {
	Input
	smaller float32
	larger  float32
}

// Gets the pageWidth.
func (p *Converter) PageWidth() float32 {
	return p.pageWidth
}

// sets the pageWidth.
func (p *Converter) SetPageWidth(value float32) {
	p.pageWidth = value
}

// Gets the PageHeight.
func (p *Converter) PageHeight() float32 {
	return p.pageHeight
}

// sets the PageHeight.
func (p *Converter) SetPageHeight(value float32) {
	p.pageHeight = value
}

// Gets the TopMargin.
func (p *Converter) TopMargin() float32 {
	return p.topMargin
}

// sets the TopMargin.
func (p *Converter) SetTopMargin(value float32) {
	p.topMargin = value
}

// Gets the BottomMargin.
func (p *Converter) BottomMargin() float32 {
	return p.bottomMargin
}

// sets the BottomMargin.
func (p *Converter) SetBottomMargin(value float32) {
	p.bottomMargin = value
}

// Gets the LeftMargin.
func (p *Converter) LeftMargin() float32 {
	return p.leftMargin
}

// sets the LeftMargin.
func (p *Converter) SetLeftMargin(value float32) {
	p.leftMargin = value
}

// Gets the RightMargin.
func (p *Converter) RightMargin() float32 {
	return p.rightMargin
}

// sets the RightMargin.
func (p *Converter) SetRightMargin(value float32) {
	p.rightMargin = value
}

// Gets the PageSize.
func (p *Converter) PageSize() PageSize {
	return p.pageSize
}

// sets the PageSize.
func (p *Converter) SetPageSize(value PageSize) {
	p.pageSize = value
	p.getPaperSize(value)
	if p.pageOrientation == Landscape {
		p.applyOrientation(Landscape)
	}
	p.pageHeight = p.larger
	p.pageWidth = p.smaller
}

// Gets the PageOrientation.
func (p *Converter) PageOrientation() Orientation {
	return p.pageOrientation
}

// sets the PageOrientation.
func (p *Converter) SetPageOrientation(value Orientation) {
	p.pageOrientation = value
	if p.pageOrientation == Landscape {
		p.applyOrientation(Landscape)
	}
	p.pageHeight = p.larger
	p.pageWidth = p.smaller
}

func (p *Converter) applyOrientation(value Orientation) {
	if value == Landscape {
		var temp = p.smaller
		p.smaller = p.larger
		p.larger = temp
	}
}
func (p *Converter) getPaperSize(size PageSize) {
	switch size {
	case Letter:
		p.smaller = p.InchesToPoints(8.5)
		p.larger = p.InchesToPoints(11)
		break
	case Legal:
		p.smaller = p.InchesToPoints(8.5)
		p.larger = p.InchesToPoints(14)
		break
	case Executive:
		p.smaller = p.InchesToPoints(7.25)
		p.larger = p.InchesToPoints(10.5)
		break
	case Tabloid:
		p.smaller = p.InchesToPoints(11)
		p.larger = p.InchesToPoints(17)
		break
	case Envelope10:
		p.smaller = p.InchesToPoints(4.125)
		p.larger = p.InchesToPoints(9.5)
		break
	case EnvelopeMonarch:
		p.smaller = p.InchesToPoints(3.875)
		p.larger = p.InchesToPoints(7.5)
		break
	case Folio:
		p.smaller = p.InchesToPoints(8.5)
		p.larger = p.InchesToPoints(13)
		break
	case Statement:
		p.smaller = p.InchesToPoints(5.5)
		p.larger = p.InchesToPoints(8.5)
		break
	case A4:
		p.smaller = p.MillimetersToPoints(210)
		p.larger = p.MillimetersToPoints(297)
		break
	case A5:
		p.smaller = p.MillimetersToPoints(148)
		p.larger = p.MillimetersToPoints(210)
		break
	case B4:
		p.smaller = p.MillimetersToPoints(250)
		p.larger = p.MillimetersToPoints(353)
		break
	case B5:
		p.smaller = p.MillimetersToPoints(176)
		p.larger = p.MillimetersToPoints(250)
		break
	case A3:
		p.smaller = p.MillimetersToPoints(297)
		p.larger = p.MillimetersToPoints(420)
		break
	case B3:
		p.smaller = p.MillimetersToPoints(353)
		p.larger = p.MillimetersToPoints(500)
		break
	case A6:
		p.smaller = p.MillimetersToPoints(105)
		p.larger = p.MillimetersToPoints(148)
		break
	case B5JIS:
		p.smaller = p.MillimetersToPoints(182)
		p.larger = p.MillimetersToPoints(257)
		break
	case EnvelopeDL:
		p.smaller = p.MillimetersToPoints(110)
		p.larger = p.MillimetersToPoints(220)
		break
	case EnvelopeC5:
		p.smaller = p.MillimetersToPoints(162)
		p.larger = p.MillimetersToPoints(229)
		break
	case EnvelopeB5:
		p.smaller = p.MillimetersToPoints(176)
		p.larger = p.MillimetersToPoints(250)
		break
	case PRC16K:
		p.smaller = p.MillimetersToPoints(146)
		p.larger = p.MillimetersToPoints(215)
		break
	case PRC32K:
		p.smaller = p.MillimetersToPoints(97)
		p.larger = p.MillimetersToPoints(151)
		break
	case Quatro:
		p.smaller = p.MillimetersToPoints(215)
		p.larger = p.MillimetersToPoints(275)
		break
	case DoublePostcard:
		p.smaller = p.MillimetersToPoints(148.0)
		p.larger = p.MillimetersToPoints(200.0)
		break
	case Postcard:
		p.smaller = p.InchesToPoints(3.94)
		p.larger = p.InchesToPoints(5.83)
		break
	default:
		p.smaller = p.InchesToPoints(8.5)
		p.larger = p.InchesToPoints(11)
		break
	}
}

func (p *Converter) InchesToPoints(size float32) float32 {
	return size * 72.0
}

func (p *Converter) MillimetersToPoints(size float32) float32 {
	return size * 2.8346456692913385826771653543307
}
