package endpoint

import (
	"encoding/json"
	"strings"

	"github.com/dynamicpdf-api/go-client/v2/resource"
)

// Represents the pdf info endpoint.
type PdfSecurityInfoEndpoint struct {
	Endpoint
	resource.Resource
	resource resource.PdfResource
}

/*
Initializes a new instance of the `PdfSecurityInfo` class.
  - @param {Resource} resource The resource of type `PdfResource`
*/
func NewPdfSecurityInfoResource(resource resource.PdfResource) *PdfSecurityInfoEndpoint {
	var ep PdfSecurityInfoEndpoint
	ep.resource = resource
	return &ep
}

var _ EndpointProcessor = (*PdfSecurityInfoEndpoint)(nil)

func (p *PdfSecurityInfoEndpoint) BaseUrl() string {
	if p.Endpoint.BaseUrl != "" {
		return p.Endpoint.BaseUrl
	} else {
		return DefaultBaseUrl
	}
}

func (p *PdfSecurityInfoEndpoint) ApiKey() string {
	return p.Endpoint.ApiKey
}

func (p *PdfSecurityInfoEndpoint) EndpointName() string {
	return "pdf-security-info"
}

/*
Process the pdf resource to get pdf's security information.
  - @returns A Promise of PdfSecurityInfo response
*/
func (p *PdfSecurityInfoEndpoint) Process() <-chan PdfSecurityInfoResponse {
	restResponse := make(chan PdfSecurityInfoResponse)
	postUrl := strings.TrimSuffix(p.BaseUrl(), "/") + "/v1.0/" + p.EndpointName()
	postAuth := "Bearer " + p.Endpoint.ApiKey
	go func() {
		res, err := postHttpRequest(postUrl, p.resource.Data(), postAuth, "application/pdf")
		if err != nil {
			res.clientError = err
		}
		var pdfRes PdfSecurityInfoResponse
		pdfRes.JsonResponse = res

		if err != nil {
			pdfRes.clientError = err
		} else if res.content != nil {
			var parsed PdfSecurityInfo
			if e := json.Unmarshal(res.content.Bytes(), &parsed); e == nil {
				parsed.EncryptionType = parsed.GetEncryptionType()
				pdfRes.Content = &parsed
			}
		}
		restResponse <- pdfRes
	}()
	return restResponse
}

func (p PdfSecurityInfoEndpoint) MarshalJSON() ([]byte, error) {
	type Alias PdfSecurityInfoEndpoint
	return json.Marshal(&struct {
		ResourceName resource.PdfResource `json:"resourceName,omitempty"`
		Alias
	}{
		ResourceName: p.resource,
		Alias:        (Alias)(p),
	})
}
