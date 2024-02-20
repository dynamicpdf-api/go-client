package input

import (
	"github.com/dynamicpdf-api/go-client/v2/resource"
	"github.com/google/uuid"
)

// Represents an image input.
type Image struct {
	Input
}

/*
Initializes a new instance of the `ImageInput` class.
  - @param {string} resource object to create ImageInput. | The image file path present in cloud resource manager.
*/
func NewImageWithResourcePath(value string) *Image {
	var p Image
	p.ResourceName = value
	p.inputType = p.InputType()
	return &p
}

/*
Initializes a new instance of the `ImageInput` class.
  - @param {ImageResource} resource object to create ImageInput. | The image file path present in cloud resource manager.
*/
func NewImagewithImageResource(value resource.ImageResource) *Image {
	var p Image
	p.resources = append(p.resources, value.Resource)
	p.SetId(uuid.NewString())
	p.inputType = p.InputType()
	p.ResourceName = value.ResourceName
	return &p
}

func (p *Image) InputType() InputType {
	return ImageInput
}

// Gets the pageWidth.
func (p *Image) PageWidth() float32 {
	return p.pageWidth
}

// sets the pageWidth.
func (p *Image) SetPageWidth(value float32) {
	p.pageWidth = value
}

// Gets the PageHeight.
func (p *Image) PageHeight() float32 {
	return p.pageHeight
}

// sets the PageHeight.
func (p *Image) SetPageHeight(value float32) {
	p.pageHeight = value
}

// Gets the TopMargin.
func (p *Image) TopMargin() float32 {
	return p.topMargin
}

// sets the TopMargin.
func (p *Image) SetTopMargin(value float32) {
	p.topMargin = value
}

// Gets the BottomMargin.
func (p *Image) BottomMargin() float32 {
	return p.bottomMargin
}

// sets the BottomMargin.
func (p *Image) SetBottomMargin(value float32) {
	p.bottomMargin = value
}

// Gets the LeftMargin.
func (p *Image) LeftMargin() float32 {
	return p.leftMargin
}

// sets the LeftMargin.
func (p *Image) SetLeftMargin(value float32) {
	p.leftMargin = value
}

// Gets the RightMargin.
func (p *Image) RightMargin() float32 {
	return p.rightMargin
}

// sets the RightMargin.
func (p *Image) SetRightMargin(value float32) {
	p.rightMargin = value
}
