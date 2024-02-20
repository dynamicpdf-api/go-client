package input

import (
	"github.com/dynamicpdf-api/go-client/v2/resource"
	"github.com/google/uuid"
)

// Represents a HTML input.
type Html struct {
	Input
	smaller float32
	larger  float32
}

var _ InputCollector = (*Html)(nil)

/*
Initializes a new instance of the `HTMLInput` class.
  - @param { string} input The html embeded as a string.
*/
func NewHtmlInputWithString(inputString string) *Html {
	var p Html
	var ep resource.HtmlResource
	var name = uuid.NewString()
	resourceData := []byte(inputString)
	ep = resource.HtmlResource{Resource: *resource.NewResourceWithByteValue(resourceData, name)}
	p.resources = append(p.resources, ep.Resource)
	p.inputType = p.InputType()
	p.getPaperSize(Letter)
	p.SetPageHeight(p.larger)
	p.SetPageWidth(p.smaller)
	p.ResourceName = ep.ResourceName
	return &p
}

/*
Initializes a new instance of the `HTMLInput` class.
  - @param { string} input The html embeded as a string.
*/
func NewHtmlInputWithResource(resource resource.HtmlResource) *Html {
	var p Html
	p.resources = append(p.resources, resource.Resource)
	p.SetId(uuid.NewString())
	p.inputType = p.InputType()
	p.getPaperSize(Letter)
	p.SetPageHeight(p.larger)
	p.SetPageWidth(p.smaller)
	p.ResourceName = resource.ResourceName
	return &p
}

func (p *Html) Resources() []resource.Resource {
	return p.resources
}

func (p *Html) InputType() InputType {
	return HtmlInput
}

// Gets the pageWidth.
func (p *Html) PageWidth() float32 {
	return p.pageWidth
}

// sets the pageWidth.
func (p *Html) SetPageWidth(value float32) {
	p.pageWidth = value
}

// Gets the PageHeight.
func (p *Html) PageHeight() float32 {
	return p.pageHeight
}

// sets the PageHeight.
func (p *Html) SetPageHeight(value float32) {
	p.pageHeight = value
}

// Gets the BasePath.
func (p *Html) BasePath() string {
	return p.basePath
}

// sets the BasePath.
func (p *Html) SetBasePath(value string) {
	p.basePath = value
}

// Gets the TopMargin.
func (p *Html) TopMargin() float32 {
	return p.topMargin
}

// sets the TopMargin.
func (p *Html) SetTopMargin(value float32) {
	p.topMargin = value
}

// Gets the BottomMargin.
func (p *Html) BottomMargin() float32 {
	return p.bottomMargin
}

// sets the BottomMargin.
func (p *Html) SetBottomMargin(value float32) {
	p.bottomMargin = value
}

// Gets the LeftMargin.
func (p *Html) LeftMargin() float32 {
	return p.leftMargin
}

// sets the LeftMargin.
func (p *Html) SetLeftMargin(value float32) {
	p.leftMargin = value
}

// Gets the RightMargin.
func (p *Html) RightMargin() float32 {
	return p.rightMargin
}

// sets the RightMargin.
func (p *Html) SetRightMargin(value float32) {
	p.rightMargin = value
}

// Gets the PageSize.
func (p *Html) PageSize() PageSize {
	return p.pageSize
}

// sets the PageSize.
func (p *Html) SetPageSize(value PageSize) {
	p.pageSize = value
	p.getPaperSize(value)
	if p.pageOrientation == Landscape {
		p.applyOrientation(Landscape)
	}
	p.pageHeight = p.larger
	p.pageWidth = p.smaller
}

// Gets the PageOrientation.
func (p *Html) PageOrientation() Orientation {
	return p.pageOrientation
}

// sets the PageOrientation.
func (p *Html) SetPageOrientation(value Orientation) {
	p.pageOrientation = value
	if p.pageOrientation == Landscape {
		p.applyOrientation(Landscape)
	}
	p.pageHeight = p.larger
	p.pageWidth = p.smaller
}
func (p *Html) applyOrientation(value Orientation) {
	if value == Landscape {
		var temp = p.smaller
		p.smaller = p.larger
		p.larger = temp
	}
}
func (p *Html) getPaperSize(size PageSize) {
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

func (p *Html) InchesToPoints(size float32) float32 {
	return size * 72.0
}

func (p *Html) MillimetersToPoints(size float32) float32 {
	return size * 2.8346456692913385826771653543307
}
