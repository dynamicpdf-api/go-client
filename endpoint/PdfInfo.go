package endpoint

import (
	"encoding/json"
	"strings"

	"github.com/dynamicpdf-api/go-client/v2/resource"
)

// Represents the pdf info endpoint.
type PdfInfo struct {
	Endpoint
	resource.Resource
	resource resource.PdfResource
}

/*
Initializes a new instance of the `PdfInfo` class.
  - @param {Resource} resource The resource of type `PdfResource`
*/
func NewPdfInfoResource(resource resource.PdfResource) *PdfInfo {
	var ep PdfInfo
	ep.resource = resource
	return &ep
}

var _ EndpointProcessor = (*PdfInfo)(nil)

func (p *PdfInfo) BaseUrl() string {
	if p.Endpoint.BaseUrl != "" {
		return p.Endpoint.BaseUrl
	} else {
		return DefaultBaseUrl
	}
}

func (p *PdfInfo) ApiKey() string {
	return p.Endpoint.ApiKey
}

func (p *PdfInfo) EndpointName() string {
	return "pdf-info"
}

/*
Process the pdf resource to get pdf's information.
  - @returns A Promise of PdfInfo response
*/
func (p *PdfInfo) Process() <-chan PdfInfoResponse {
	restResponse := make(chan PdfInfoResponse)
	postUrl := strings.TrimSuffix(p.BaseUrl(), "/") + "/v1.0/" + p.EndpointName()
	postAuth := "Bearer " + p.Endpoint.ApiKey
	go func() {
		res, err := postHttpRequest(postUrl, p.resource.Data(), postAuth, "application/pdf")
		if err != nil {
			res.clientError = err
		}
		pdfRes := PdfInfoResponse{JsonResponse: res}
		restResponse <- pdfRes
	}()
	return restResponse
}

func (p PdfInfo) MarshalJSON() ([]byte, error) {
	type Alias PdfInfo
	return json.Marshal(&struct {
		ResourceName resource.PdfResource `json:"resourceName,omitempty"`
		Alias
	}{
		ResourceName: p.resource,
		Alias:        (Alias)(p),
	})
}
