package input

import (
	"github.com/dynamicpdf-api/go-client/v2/resource"
	"github.com/google/uuid"
)

// Represents a Excel input.
type Excel struct {
	Converter
}

/*
Initializes a new instance of the `ExcelInput` class.
  - @param { string} input The Excel embeded as a string.
*/
func NewExcelInputWithResource(resources resource.ExcelResource) *Excel {
	var p Excel
	p.resources = append(p.resources, resources.Resource)
	p.SetId(uuid.NewString())
	p.inputType = p.InputType()
	p.getPaperSize(Letter)
	p.SetPageHeight(p.larger)
	p.SetPageWidth(p.smaller)
	p.ResourceName = resources.ResourceName
	return &p
}

func (p *Excel) Resources() []resource.Resource {
	return p.resources
}

func (p *Excel) InputType() InputType {
	return ExcelInput
}
