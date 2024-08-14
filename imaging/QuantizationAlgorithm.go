package imaging

// Enum representing quantization algorithms.
type QuantizationAlgorithm string

const (
	// Octree quantization algorithm.
	Octree QuantizationAlgorithm = "Octree"

	// WebSafe quantization algorithm.
	WebSafe QuantizationAlgorithm = "WebSafe"

	// Werner quantization algorithm.
	Werner QuantizationAlgorithm = "Werner"

	// WU quantization algorithm.
	WU QuantizationAlgorithm = "WU"
)
