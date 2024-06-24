package input

import (
	"encoding/json"

	"github.com/dynamicpdf-api/go-client/v2/element"
	"github.com/google/uuid"
)

// Represents a page input.
type Page struct {
	Input
	smaller float32
	larger  float32
}

var _ InputCollector = (*Page)(nil)
var _ json.Marshaler = (*Page)(nil)

func NewPage() *Page {
	var p Page
	p.Elements = []element.ElementCollector{}
	p.SetId(uuid.NewString())
	p.inputType = p.InputType()
	return &p
}

/*
Initializes a new instance of the `PageInput` class.
  - @param {float} pageWidth The width of the page.
  - @param {float} pageHeight The height of the page.
*/
func NewPageWithDimension(width float32, height float32) *Page {
	var p Page
	p.Elements = []element.ElementCollector{}
	p.SetId(uuid.NewString())
	p.inputType = p.InputType()
	p.SetPageHeight(height)
	p.SetPageWidth(width)
	return &p
}

/*
Initializes a new instance of the `PageInput` class.
  - @param {PageSize} pageSize The size of the page.
  - @param {Orientation} Orientation The orientation of the page.
  - @param {float32} margins The margins of the page.
*/
func NewPageWithSizeAndOrientation(pageSize PageSize, pageOrientation Orientation, margins float32) *Page {
	var p Page
	p.Elements = []element.ElementCollector{}
	p.SetId(uuid.NewString())
	p.inputType = p.InputType()
	p.SetPageSize(pageSize)
	p.SetPageOrientation(pageOrientation)
	p.SetTopMargin(margins)
	p.SetBottomMargin(margins)
	p.SetRightMargin(margins)
	p.SetLeftMargin(margins)
	return &p
}

func (p *Page) InputType() InputType {
	return PageInput
}
func (p *Page) Element() []element.ElementCollector {
	return p.Elements
}

// Gets the page width.
func (p *Page) PageWidth() float32 {
	return p.pageWidth
}

// sets the page width.
func (p *Page) SetPageWidth(value float32) {
	p.pageWidth = value
}

// Gets the page height.
func (p *Page) PageHeight() float32 {
	return p.pageHeight
}

// sets the page height.
func (p *Page) SetPageHeight(value float32) {
	p.pageHeight = value
}

// Gets the TopMargin.
func (p *Page) TopMargin() float32 {
	return p.topMargin
}

// sets the TopMargin.
func (p *Page) SetTopMargin(value float32) {
	p.topMargin = value
}

// Gets the BottomMargin.
func (p *Page) BottomMargin() float32 {
	return p.bottomMargin
}

// sets the BottomMargin.
func (p *Page) SetBottomMargin(value float32) {
	p.bottomMargin = value
}

// Gets the LeftMargin.
func (p *Page) LeftMargin() float32 {
	return p.leftMargin
}

// sets the LeftMargin.
func (p *Page) SetLeftMargin(value float32) {
	p.leftMargin = value
}

// Gets the RightMargin.
func (p *Page) RightMargin() float32 {
	return p.rightMargin
}

// sets the RightMargin.
func (p *Page) SetRightMargin(value float32) {
	p.rightMargin = value
}

// Gets the PageSize.
func (p *Page) PageSize() PageSize {
	return p.pageSize
}

// sets the PageSize.
func (p *Page) SetPageSize(value PageSize) {
	p.pageSize = value
	p.getPaperSize(value)
	if p.pageOrientation == Landscape {
		p.pageHeight = p.smaller
		p.pageWidth = p.larger
	} else {
		p.pageHeight = p.larger
		p.pageWidth = p.smaller
	}
}

// Gets the PageOrientation.
func (p *Page) PageOrientation() Orientation {
	return p.pageOrientation
}

// sets the PageOrientation.
func (p *Page) SetPageOrientation(value Orientation) {
	p.pageOrientation = value
	if p.pageWidth > p.pageHeight {
		p.smaller = p.pageHeight
		p.larger = p.pageWidth
	} else {
		p.smaller = p.pageWidth
		p.larger = p.pageHeight
	}
	if p.pageOrientation == Landscape {
		p.pageHeight = p.smaller
		p.pageWidth = p.larger
	} else {
		p.pageHeight = p.larger
		p.pageWidth = p.smaller
	}
}

func (p *Page) getPaperSize(size PageSize) {
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

func (p *Page) InchesToPoints(size float32) float32 {
	return size * 72.0
}

func (p *Page) MillimetersToPoints(size float32) float32 {
	return size * 2.8346456692913385826771653543307
}

func (p Page) MarshalJSON() ([]byte, error) {
	type Alias Page
	return json.Marshal(&struct {
		Elements []element.ElementCollector `json:"elements"`
		Alias
	}{
		Elements: p.Elements,
		Alias:    (Alias)(p),
	})
}
