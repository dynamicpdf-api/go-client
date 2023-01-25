package color

/**
 * Represents information of a signature field.
 */
type WebColor struct {
	RgbColor
}

/**
 * Initializes a new instance of the `WebColor` class.
 * @param {string} webHexString The hexadecimal string representing the color.
 */
func NewWebColor(color string) *WebColor {
	var ep WebColor
	ep.colorString = color
	return &ep
}
