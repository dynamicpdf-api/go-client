package imaging

// Enum representing dithering algorithms.
type DitheringAlgorithm string

const (
	// FloydSteinberg dithering algorithm.
	FloydSteinberg DitheringAlgorithm = "Floyd-Steinberg"

	// Bayer dithering algorithm.
	Bayer DitheringAlgorithm = "Bayer"

	// No dithering.
	None DitheringAlgorithm = "None"
)
