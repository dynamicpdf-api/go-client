package element

// Represents GS1 DataBar types.
type Gs1DataBarType string

const (
	// Omnidirectional GS1 DataBar type.
	Omnidirectional Gs1DataBarType = "Omnidirectional"
	// Limited GS1 DataBar type.
	Limited Gs1DataBarType = "Limited"
	// Expanded GS1 DataBar type.
	Expanded Gs1DataBarType = "Expanded"
)
