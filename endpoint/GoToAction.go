package endpoint

import "github.com/dynamicpdf-api/go-client/input"

// Represents a goto action in a PDF document that navigates to a specific page using page number and zoom options.
type GoToAction struct {
	Action
	// Gets or sets `PageZoom` to display the destination.
	PageZoom PageZoom `json:"pageZoom,omitempty"`
	// Gets or sets page Offset.
	PageOffset int `json:"pageOffset,omitempty"`
	inputID    int
	input      input.Input
}

/*
 Initializes a new instance of the  `GoToAction` class
  * using an input to create the PDF, page number, and a zoom option.
  * @param {Input} input Any of the `ImageInput`, `DlexInput`, `PdfInput` or `PageInput` objects to create PDF.
  * @param {int} pgOffset Page number to navigate.
  * @param {pageZoom} pgZoom Page Zoom to display the destination.
*/
func NewGoToAction(input input.Input, pageOffset int, pageZoom PageZoom) *GoToAction {
	var ep GoToAction
	ep.input = input
	ep.PageOffset = pageOffset
	ep.PageZoom = pageZoom
	return &ep
}
