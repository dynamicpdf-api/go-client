package input

import (
	"github.com/dynamicpdf-api/go-client/v2/resource"
	"github.com/google/uuid"
)

// Represents a Word input.
type Word struct {
	Converter
}

/*
Initializes a new instance of the `WordInput` class.
  - @param { string} input The Word embeded as a string.
*/
func NewWordInputWithResource(resources resource.WordResource) *Word {
	var p Word
	p.resources = append(p.resources, resources.Resource)
	p.SetId(uuid.NewString())
	p.inputType = p.InputType()
	p.ResourceName = resources.ResourceName
	return &p
}

func (p *Word) Resources() []resource.Resource {
	return p.resources
}

func (p *Word) InputType() InputType {
	return WordInput
}
