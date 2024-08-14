package imaging

import "encoding/json"

// Base class for image size types.

type ImageSize struct {
	Type ImageSizeType `json:"type,omitempty"`

	width                int
	height               int
	unit                 ImageSizeUnit
	horizontalDpi        int
	verticalDpi          int
	maxWidth             int
	maxHeight            int
	horizontalPercentage int
	verticalPercentage   int
}

func (p *ImageSize) MarshalJSON() ([]byte, error) {
	type Alias ImageSize
	return json.Marshal(&struct {
		Width                int           `json:"width,omitempty"`
		Height               int           `json:"height,omitempty"`
		Unit                 ImageSizeUnit `json:"unit,omitempty"`
		HorizontalDpi        int           `json:"horizontalDpi,omitempty"`
		VerticalDpi          int           `json:"verticalDpi,omitempty"`
		MaxWidth             int           `json:"maxWidth,omitempty"`
		MaxHeight            int           `json:"maxHeight,omitempty"`
		HorizontalPercentage int           `json:"horizontalPercentage,omitempty"`
		VerticalPercentage   int           `json:"verticalPercentage,omitempty"`
		*Alias
	}{
		Width:                p.width,
		Height:               p.height,
		Unit:                 p.unit,
		HorizontalDpi:        p.horizontalDpi,
		VerticalDpi:          p.verticalDpi,
		MaxWidth:             p.maxWidth,
		MaxHeight:            p.maxHeight,
		HorizontalPercentage: p.horizontalDpi,
		VerticalPercentage:   p.verticalDpi,
		Alias:                (*Alias)(p),
	})
}
