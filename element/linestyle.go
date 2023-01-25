package element

import "fmt"

type LineStyle struct {
	lineStyleString string
}

func newLineStyle(dashArray string) *LineStyle {
	var p LineStyle
	p.lineStyleString = dashArray
	return &p
}

/**
 * Initializes a new instance of the `LineStyle` class.
 * @param {float []} dashArray The array specifying the line style.
 * @param {float} dashPhase The phase of the line style.
 */
func NewLineStyle(dashArray []float64, dashPhase float64) *LineStyle {
	var p LineStyle
	var strLineStyle = "["
	for i := 0; i < len(dashArray); i++ {
		var val = dashArray[i]
		if i == len(dashArray)-1 {
			strLineStyle += fmt.Sprint(val) + "0.00"
		} else {
			strLineStyle += fmt.Sprint(val) + "0.00" + ","
		}
		strLineStyle += "]"
	}
	if dashPhase != 0 {
		strLineStyle += fmt.Sprint(dashPhase) + "0.00"
	}
	p.lineStyleString = strLineStyle
	return &p
}

/** Gets a solid line.*/
func (p *LineStyle) Solid() *LineStyle {
	return newLineStyle("solid")
}

/** Gets a dotted line. */
func (p *LineStyle) Dots() *LineStyle {
	return newLineStyle("dots")
}

/** Gets a line with small dashes. */
func (p *LineStyle) DashSmall() *LineStyle {
	return newLineStyle("dashSmall")
}

/** Gets a dashed line. */
func (p *LineStyle) Dash() *LineStyle {
	return newLineStyle("dash")
}

/** Gets a line with large dashes. */
func (p *LineStyle) DashLarge() *LineStyle {
	return newLineStyle("dashLarge")
}

/** Gets a invisible line.  */
func (p *LineStyle) None() *LineStyle {
	return newLineStyle("none")
}
