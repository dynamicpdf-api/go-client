package endpoint

import (
	"encoding/json"
	"strings"

	"github.com/dynamicpdf-api/go-client/resource"
)

/* Represents the pdfresource. */
type PdfXmp struct {
	Endpoint

	resource resource.PdfResource
}

/**
 * Initializes a new instance of the `PdfXmp` class.
 * @param {PdfResource} resource The image resource of type `PdfResource`.
 */
func NewPdfXmp(resource resource.PdfResource) *PdfXmp {
	var ep PdfXmp
	ep.resource = resource
	return &ep
}

var _ EndpointProcessor = (*PdfXmp)(nil)

func (p *PdfXmp) BaseUrl() string {
	return p.Endpoint.BaseUrl
}

func (p *PdfXmp) ApiKey() string {
	return p.Endpoint.ApiKey
}

func (p *PdfXmp) EndpointName() string {
	return "pdf-xmp"
}

/**  Process the pdf resource to get pdf's xmp data.
 * @returns A Promise of PdfTextResponse callback.
 */
func (p *PdfXmp) Process() <-chan PdfXmpResponse {
	restResponse := make(chan PdfXmpResponse)
	postUrl := strings.TrimSuffix(p.Endpoint.BaseUrl, "/") + "/v1.0/" + p.EndpointName()
	postAuth := "Bearer " + p.Endpoint.ApiKey
	go func() {
		res, err := postHttpRequest(postUrl, p.resource.Data(), postAuth, "application/pdf")
		if err != nil {
			res.clientError = err
		}
		pdfRes := PdfXmpResponse{JsonResponse: res}
		restResponse <- pdfRes
	}()
	return restResponse
}

func (p PdfXmp) MarshalJSON() ([]byte, error) {
	type Alias PdfXmp
	return json.Marshal(&struct {
		ResourceName resource.PdfResource `json:"resourceName,omitempty"`
		Alias
	}{
		ResourceName: p.resource,
		Alias:        (Alias)(p),
	})
}
