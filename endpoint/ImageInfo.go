package endpoint

import (
	"encoding/json"
	"strings"

	"github.com/dynamicpdf-api/go-client/v2/resource"
	"github.com/gabriel-vasile/mimetype"
)

// Represents an image information endpoint.
type ImageInfo struct {
	Endpoint
	resource resource.ImageResource
}

/*
Initializes a new instance of the `ImageInfo` class.
  - @param {ImageResource} resource The image resource of type `ImageResource`.
*/
func NewImageInfo(image resource.ImageResource) *ImageInfo {
	var ep ImageInfo
	ep.resource = image
	return &ep
}

var _ EndpointProcessor = (*ImageInfo)(nil)

func (p *ImageInfo) BaseUrl() string {
	return p.Endpoint.BaseUrl
}

func (p *ImageInfo) ApiKey() string {
	return p.Endpoint.ApiKey
}

func (p *ImageInfo) EndpointName() string {
	return "image-info"
}

/*
Process the image resource to get image's information.
  - @returns A Promise of ImageResponse callback.
*/
func (p *ImageInfo) Process() <-chan ImageResponse {
	restResponse := make(chan ImageResponse)
	postUrl := strings.TrimSuffix(p.Endpoint.BaseUrl, "/") + "/v1.0/" + p.EndpointName()
	postAuth := "Bearer " + p.Endpoint.ApiKey
	mtype := mimetype.Detect(p.resource.Data())
	go func() {
		res, err := postHttpRequest(postUrl, p.resource.Data(), postAuth, mtype.String())
		if err != nil {
			res.clientError = err
		}
		pdfRes := ImageResponse{JsonResponse: res}
		restResponse <- pdfRes
	}()
	return restResponse
}

func (p ImageInfo) MarshalJSON() ([]byte, error) {
	type Alias ImageInfo
	return json.Marshal(&struct {
		ResourceName resource.ImageResource `json:"resourceName,omitempty"`
		Alias
	}{
		ResourceName: p.resource,
		Alias:        (Alias)(p),
	})
}
