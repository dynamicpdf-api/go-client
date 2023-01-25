package element

import (
	"github.com/dynamicpdf-api/go-client/resource"
)

/** Represents an image element.
 * This class can be used to place images on a page.
 */
type Image struct {
	Element
}

var _ ElementCollector = (*Image)(nil)

/**
 * *Initializes a new instance of the `ImageElement` class.     *
 * @param {string} resource The name of the image resource. | object containing the image resource.
 * @param {ElementPlacement} placement The placement of the image on the page.
 * @param {float} xOffset X coordinate of the image.
 * @param {float} yOffset Y coordinate of the image.
 */
func NewImageWithResourcePath(value string, placement elementPlacement, xOffset float32, yOffset float32) *Image {
	var p = Image{Element: *newElement(value, placement, xOffset, yOffset)}
	p.typeOfElement = p.ElementType()
	return &p
}

/**
 * *Initializes a new instance of the `ImageElement` class.     *
 * @param {Resource} resource The name of the image resource. | object containing the image resource.
 * @param {ElementPlacement} placement The placement of the image on the page.
 * @param {float} xOffset X coordinate of the image.
 * @param {float} yOffset Y coordinate of the image.
 */
func NewImagewithImageResource(value resource.ImageResource, placement elementPlacement, xOffset float32, yOffset float32) *Image {
	var p Image
	p.typeOfElement = p.ElementType()
	p.resource = value.Resource
	p.resourceName = value.ResourceName
	p.Placement = placement
	p.XOffset = xOffset
	p.YOffset = yOffset
	return &p
}

func (p *Image) ElementType() ElementType {
	return ImageElement
}

/** Gets the horizontal scale of the image. */
func (p *Image) ScaleX() float32 {
	return p.scaleX
}

/** Sets the horizontal scale of the image. */
func (p *Image) SetScaleX(value float32) {
	p.scaleX = value
}

/** Gets the vertical scale of the image. */
func (p *Image) ScaleY() float32 {
	return p.scaleY
}

/** Sets the vertical scale of the image. */
func (p *Image) SetScaleY(value float32) {
	p.scaleY = value
}

/** Gets the maximum height of the image. */
func (p *Image) MaxHeight() float32 {
	return p.maxHeight
}

/** Sets the maximum height of the image. */
func (p *Image) SetMaxHeight(value float32) {
	p.maxHeight = value
}

/** Gets the maximum width of the image.  */
func (p *Image) MaxWidth() float32 {
	return p.maxWidth
}

/** Sets the maximum width of the image.  */
func (p *Image) SetMaxWidth(value float32) {
	p.maxWidth = value
}
