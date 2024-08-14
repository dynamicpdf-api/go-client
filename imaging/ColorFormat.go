package imaging

import "encoding/json"

// Base class for all color formats.
type ColorFormat struct {

	// Gets or sets the color format type.
	Type ColorFormatType `json:"type,omitempty"`

	quantizationAlgorithm QuantizationAlgorithm
	ditheringPercent      int
	blackThreshold        int
	ditheringAlgorithm    DitheringAlgorithm
	compressionType       CompressionType
}

func (p *ColorFormat) MarshalJSON() ([]byte, error) {
	type Alias ColorFormat
	return json.Marshal(&struct {
		QuantizationAlgorithm QuantizationAlgorithm `json:"quantizationAlgorithm,omitempty"`
		DitheringPercent      int                   `json:"ditheringPercent,omitempty"`
		BlackThreshold        int                   `json:"blackThreshold,omitempty"`
		DitheringAlgorithm    DitheringAlgorithm    `json:"ditheringAlgorithm,omitempty"`
		CompressionType       CompressionType       `json:"compressionType,omitempty"`
		*Alias
	}{
		QuantizationAlgorithm: p.quantizationAlgorithm,
		DitheringPercent:      p.ditheringPercent,
		BlackThreshold:        p.blackThreshold,
		DitheringAlgorithm:    p.ditheringAlgorithm,
		CompressionType:       p.compressionType,
		Alias:                 (*Alias)(p),
	})
}
