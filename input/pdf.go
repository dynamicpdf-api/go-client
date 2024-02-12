package input

import (
	"encoding/json"

	"github.com/dynamicpdf-api/go-client/element"
	"github.com/dynamicpdf-api/go-client/resource"
	"github.com/google/uuid"
)

// Represents a pdf input.
type Pdf struct {
	Input

	pdfresource []resource.PdfResource
	Elements    []element.ElementCollector `json:"elements"` /*TODO:Type*/
}

var _ InputCollector = (*Pdf)(nil)
var _ json.Marshaler = (*Pdf)(nil)

/*
Initializes a new instance of the `PdfInput` class.
  - @param { string } The resource path in cloud resource manager.
  - @param {MergeOptions} mergeOptions The merge options for the pdf.
*/
func NewPdfWithCloudPath(cloudResourcePath string, option MergeOptions) *Pdf {
	var p Pdf
	p.Elements = []element.ElementCollector{}
	p = Pdf{Input: *newInput(cloudResourcePath)}
	p.inputType = p.InputType()
	p.MergeOption = option
	return &p
}

/*
Initializes a new instance of the `PdfInput` class.
  - @param { PdfResource} resource The resource of type `PdfResource`.
  - @param {MergeOptions} mergeOptions The merge options for the pdf.
*/
func NewPdfWithResource(resource resource.PdfResource) *Pdf {
	p := Pdf{Input: *newInputWithResource(resource.Resource)}
	p.SetId(uuid.NewString())
	p.Elements = []element.ElementCollector{}
	p.inputType = p.InputType()
	return &p
}

/*
Initializes a new instance of the `PdfInput` class.
  - @param { PdfResource} resource The resource of type `PdfResource`.
  - @param {MergeOptions} mergeOptions The merge options for the pdf.
*/
func NewPdfWithResourceWithMergerOption(resource resource.PdfResource, option MergeOptions) *Pdf {
	var p Pdf
	p.SetId(uuid.NewString())
	p = Pdf{Input: *newInputWithResource(resource.Resource)}
	p.Elements = []element.ElementCollector{}
	p.inputType = p.InputType()
	p.MergeOption = option
	return &p
}

func (p *Pdf) InputType() InputType {
	return PdfInput
}

func (p *Pdf) Resources() []resource.Resource {
	return p.resources
}
