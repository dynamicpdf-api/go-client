package input

import (
	"github.com/dynamicpdf-api/go-client/v2/resource"
	"github.com/google/uuid"
)

// Represents a HTML input.
type Html struct {
	Converter
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
	p.ResourceName = resource.ResourceName
	return &p
}

func (p *Html) Resources() []resource.Resource {
	return p.resources
}

func (p *Html) InputType() InputType {
	return HtmlInput
}

// Gets the BasePath.
func (p *Html) BasePath() string {
	return p.basePath
}

// sets the BasePath.
func (p *Html) SetBasePath(value string) {
	p.basePath = value
}
