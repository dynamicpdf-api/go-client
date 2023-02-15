package endpoint

type ImageInformation struct {
	Endpoint
	// Gets page number of the pdf where the image is present.
	PageNumber int `json:"pageNumber,omitempty"`

	// Gets the width of the image.
	Width float32 `json:"width,omitempty"`

	// Gets the height of the image.
	Height float32 `json:"height,omitempty"`

	// Gets the horizontalDpi of the image.
	HorizontalDpi float32 `json:"horizontalDpi,omitempty"`

	// Gets the verticalDpi of the image.
	VerticalDpi float32 `json:"verticalDpi,omitempty"`

	// Gets the number of color components present in the image.
	NumberOfComponents float32 `json:"numberOfComponents,omitempty"`

	// Gets the bits per component of the image.
	BitsPerComponent float32 `json:"bitsPerComponent,omitempty"`

	// Gets the color space of the image.
	ColorSpace int `json:"colorSpace,omitempty"`
}
