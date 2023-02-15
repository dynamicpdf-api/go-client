package endpoint

import (
	"bytes"
)

// Represents an image response.
type ImageResponse struct {
	JsonResponse
}

/*
Initializes a new instance of the `ImageResponse` class.
  - @param {string} imageResponse The image content of the response.
*/
func NewImageResponse(jsonContent string) *ImageResponse {
	var pr ImageResponse
	return &pr
}

func (pr *ImageResponse) Content() *bytes.Buffer {
	return pr.content
}
