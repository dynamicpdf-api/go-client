package input

import (
	"encoding/json"

	"github.com/dynamicpdf-api/go-client/v2/resource"
)

// Represents a Dlex input.
type Dlex struct {
	Input
}

var _ InputCollector = (*Dlex)(nil)
var _ json.Marshaler = (*Dlex)(nil)

/*
	Initializes a new instance of the `DlexInput` class by posting the

DLEX file and the JSON data file or
DLEX file path that is present in the cloud environment and the JSON data file or
DLEX file path and DLEX data file path that is present in the cloud environment
from the client to the API to create the PDF report.
  - @param {DlexResource} resource dlex file created as per the desired PDF report layout design.
  - @param {LayoutDataResource} layoutData json data file used to create the PDF report.
*/
func NewDlexWithDlexNLayoutResources(dlexResource resource.DlexResource, layoutData resource.LayoutDataResource) *Dlex {
	var p Dlex
	p.inputType = p.InputType()
	p.ResourceName = dlexResource.ResourceName
	p.LayoutDataResourceName = layoutData.LayoutDataResourceName
	p.resources = append(p.resources, dlexResource.Resource)
	p.resources = append(p.resources, layoutData.Resource)
	return &p
}

/*
	Initializes a new instance of the `DlexInput` class by posting the

DLEX file and the JSON data file or
DLEX file path that is present in the cloud environment and the JSON data file or
DLEX file path and DLEX data file path that is present in the cloud environment
from the client to the API to create the PDF report.
  - @param {DlexResource} resource dlex file created as per the desired PDF report layout design.
  - @param {string} The JSON data file path present in the resource manager used to create the PDF report.
*/
func NewDlexWithDlexResourceNLayoutDataPath(dlexResource resource.DlexResource, layoutData string) *Dlex {
	var p Dlex
	p.inputType = p.InputType()
	p.ResourceName = dlexResource.ResourceName
	layoutDataResource := resource.NewLayoutDataResource(layoutData, "")
	p.LayoutDataResourceName = layoutDataResource.LayoutDataResourceName
	p.resources = append(p.resources, dlexResource.Resource)
	p.resources = append(p.resources, layoutDataResource.Resource)
	return &p
}

/*
	Initializes a new instance of the `DlexInput` class by posting the

DLEX file and the JSON data file or
DLEX file path that is present in the cloud environment and the JSON data file or
DLEX file path and DLEX data file path that is present in the cloud environment
from the client to the API to create the PDF report.
  - @param {string}The DLEX file path present in the resource manager.
  - @param {LayoutDataResource} layoutData json data file used to create the PDF report.
*/
func NewDlexWithCloudResourceNLayoutData(cloudResourcePath string, layoutData resource.LayoutDataResource) *Dlex {
	var p Dlex
	p.inputType = p.InputType()
	p.ResourceName = cloudResourcePath
	p.LayoutDataResourceName = layoutData.LayoutDataResourceName
	p.resources = append(p.resources, layoutData.Resource)
	return &p
}

/*
	Initializes a new instance of the `DlexInput` class by posting the

DLEX file and the JSON data file or
DLEX file path that is present in the cloud environment and the JSON data file or
DLEX file path and DLEX data file path that is present in the cloud environment
from the client to the API to create the PDF report.
  - @param {string} The DLEX file path present in the resource manager.
  - @param {string} The JSON data file path present in the resource manager used to create the PDF report.
*/
func NewDlexWithCloudResourceNLayoutDataPath(cloudResourcePath string, layoutData string) *Dlex {
	var p Dlex
	p.inputType = p.InputType()
	p.ResourceName = cloudResourcePath
	layoutDataResource := resource.NewLayoutDataResource(layoutData, "")
	p.LayoutDataResourceName = layoutDataResource.LayoutDataResourceName
	p.resources = append(p.resources, layoutDataResource.Resource)
	return &p
}

func (p *Dlex) InputType() InputType {
	return DlexInput
}
