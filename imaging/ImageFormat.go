package imaging

import "encoding/json"

// Base class for image formats.
type ImageFormat struct {
	// Type ImageSizeType `json:"type,omitempty"`

	Type               ImageFormatType
	ditheringPercent   int
	ditheringAlgorithm DitheringAlgorithm
	colorFormat        ColorFormat
	multiPage          bool
	quality            int
}

// Initializes a new instance of the ImageFormat class with the specified type.
func NewImageFormat(imageFormatType ImageFormatType) *ImageFormat {
	var ep ImageFormat
	ep.Type = imageFormatType
	return &ep
}

func (p *ImageFormat) MarshalJSON() ([]byte, error) {
	type Alias ImageFormat
	return json.Marshal(&struct {
		DitheringPercent   int                `json:"ditheringPercent,omitempty"`
		DitheringAlgorithm DitheringAlgorithm `json:"ditheringAlgorithm,omitempty"`
		ColorFormat        ColorFormat        `json:"colorFormat,omitempty"`
		MultiPage          bool               `json:"multiPage"`
		Quality            int                `json:"quality,omitempty"`
		*Alias
	}{
		DitheringPercent:   p.ditheringPercent,
		DitheringAlgorithm: p.ditheringAlgorithm,
		ColorFormat:        p.colorFormat,
		MultiPage:          p.multiPage,
		Quality:            p.quality,
		Alias:              (*Alias)(p),
	})
}
