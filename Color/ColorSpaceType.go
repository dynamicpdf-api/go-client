package color

type ColorSpaceType int

const (
	// Represents a Monochrome color space.
	Monochrome ColorSpaceType = 0

	// Represents a Grayscale color space.
	Grayscale ColorSpaceType = 1

	// Represents an RGB color space.
	RGB ColorSpaceType = 2

	// Represents a CMYK color space.
	CMYK ColorSpaceType = 3

	// Represents an Indexed color space.
	Indexed ColorSpaceType = 4

	// Unknown color space.
	Unknown ColorSpaceType = 5
)
