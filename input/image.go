package input

import (
	"github.com/dynamicpdf-api/go-client/position"
	"github.com/dynamicpdf-api/go-client/resource"
)

// Represents an image input.
type Image struct {
	Input

	// Gets or sets the horizontal scale of the image.
	ScaleX int `json:"scaleX"`

	// Gets or sets the vertical scale of the image.
	ScaleY int `json:"scaleY"`

	// Gets or sets a boolean indicating whether to expand the image.
	ExpandToFit float32 `json:"expandToFit"`

	// Gets or sets a boolean indicating whether to shrink the image.
	ShrinkToFit float32 `json:"shrinkToFit"`

	// Gets or sets the horizontal alignment of the image.
	Align position.Align `json:"align"`

	// Gets or sets the vertical alignment of the image.
	VAlign position.VAlign `json:"vAlign"`
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
	p.inputType = p.InputType()
	return &p
}

func (p *Image) imageType() InputType {
	return ImageInput
}
