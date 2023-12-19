package endpoint

import (
	"bytes"
	"io"
	"mime/multipart"

	"github.com/dynamicpdf-api/go-client/resource"
)

// Represents a Dlex layout endpoint.
type Dlex struct {
	Endpoint `json:"-"` // This will bekipped in JSON

	resource resource.LayoutDataResource

	// Gets or sets the DLEX file path present in the resource manager.
	DlexPath string `json:"dlexPath,omitempty"`
	Resources []resource.Resource
}

/*
	Initializes a new instance of the <see cref="DlexLayout"/> class using the

DLEX file path present in the cloud environment and the JSON data for the PDF report.
  - @param {string} cloudDlexPath The DLEX file path present in the resource manager
  - @param {LayoutDataResource} layoutData The `LayoutDataResource` json data file used to create the PDF report.
*/
func NewDlexEndpoint(cloudDlexPath string, resource resource.LayoutDataResource) *Dlex {
	var ep Dlex
	ep.resource = resource
	ep.DlexPath = cloudDlexPath
	return &ep
}

/*
	Initializes a new instance of the <see cref="DlexInput"/> class by posting the

DLEX file and the JSON data file from the client to the API to create the PDF report.
  - @param {DlexResource} dlexResource The DLEX file created as per the desired PDF report layout design.
  - @param {LayoutDataResource} layoutData The `LayoutDataResource` json data file used to create the PDF report.
*/
func NewDlexEndpointWithResource(dlexResource resource.DlexResource, resource resource.LayoutDataResource) *Dlex {
	var ep Dlex
	ep.Resources = append(ep.Resources, dlexResource.Resource)
	ep.resource = resource
	return &ep
}

/*
Adds additional resource to the endpoint.
  - @param {resourcePath} The resource file path.
  - @param {resourceName} The name of the resource.
*/
func NewDlexWithAdditionalResource(resourcePath string, resourceName string) resource.Resource {
	additionalResource := resource.NewResourceWithPath(resourcePath, resourceName)
	return additionalResource
}

var _ EndpointProcessor = (*Dlex)(nil)

func (p *Dlex) BaseUrl() string {
	return p.Endpoint.BaseUrl
}

func (p *Dlex) ApiKey() string {
	return p.Endpoint.ApiKey
}

func (p *Dlex) EndpointName() string {
	return "dlex-layout"
}

/*
Process the DLEX and layout data to create PDF report.
  - @returns A Promise of DlexResponse callback.
*/
func (p *Dlex) Process() <-chan PdfResponse {
	retResponse := make(chan PdfResponse)

	go func() {
		var form formData
		form.content = bytes.NewBuffer(nil)
		formWriter := multipart.NewWriter(form.content)
		form.contentType = formWriter.FormDataContentType()
		part, err := formWriter.CreateFormFile("LayoutData", p.resource.LayoutDataResourceName)
		if err != nil {
			return
		}
		resourceData := bytes.NewBuffer(p.resource.Data())
		io.Copy(part, resourceData)
		formWriter.WriteField("DlexPath", p.DlexPath)

		for _, finalResource1 := range p.Resources {
			part, err := formWriter.CreateFormFile("Resource", finalResource1.ResourceName)
			if err != nil {
				return
			}
			resourceData := bytes.NewBuffer(finalResource1.Data())
			io.Copy(part, resourceData)
		}

		err = formWriter.Close()
		if err != nil {
			return
		}

		res, err := postForm(form, p)
		if err != nil {
			res.clientError = err
		}
		dlexRes := PdfResponse{response: res}
		retResponse <- dlexRes
	}()

	return retResponse
}
