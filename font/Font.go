package font

import (
	"github.com/dynamicpdf-api/go-client/resource"
	"github.com/google/uuid"
)

/** Represents font. */
type Font struct {
	Resource resource.Resource

	ResourceName string `json:"resource/name,omitempty"`

	Embed string `json:"embed,omitempty"`

	Subset string `json:"subset,omitempty"`

	Name string `json:"name,omitempty"`
}

type FontStyle struct {
	Font
	timesRoman Font

	timesBold Font

	timesItalic Font

	timesBoldItalic Font

	helvetica Font

	helveticaBold Font

	helveticaOblique Font

	helveticaBoldOblique Font

	courierBold Font

	courierOblique Font

	courierBoldOblique Font

	courier Font

	symbol Font

	zapfDingbats Font
}

/**
 * Initializes a new instance of the `Font` class
 * using the font name that is present in the cloud resource manager.
 */
func NewFontStyle() *FontStyle {
	var p FontStyle
	p.Name = uuid.NewString()
	return &p
}

/** Gets the Times Italic core font with Latin 1 encoding. */
func (p *FontStyle) TimesItalic() Font {
	p.timesItalic.Name = "timesItalic"
	return p.timesItalic
}

/** Gets the Times Roman core font with Latin 1 encoding. */
func (p *FontStyle) TimesRoman() Font {
	p.timesRoman.Name = "timesItalic"
	return p.timesRoman
}

/** Gets the Times Bold core font with Latin 1 encoding. */
func (p *FontStyle) TimesBold() Font {
	p.timesBold.Name = "timesItalic"
	return p.timesItalic
}

/** Gets the Times Bold Italic core font with Latin 1 encoding. */
func (p *FontStyle) TimesBoldItalic() Font {
	p.timesBoldItalic.Name = "timesItalic"
	return p.timesItalic
}

/** Gets the Helvetica core font with Latin 1 encoding. */
func (p *FontStyle) Helvetica() Font {
	p.helvetica.Name = "timesItalic"
	return p.timesItalic
}

/** Gets the Helvetica Bold core font with Latin 1 encoding. */
func (p *FontStyle) HelveticaBold() Font {
	p.helveticaBold.Name = "timesItalic"
	return p.timesItalic
}

/** Gets the Helvetica Oblique core font with Latin 1 encoding. */
func (p *FontStyle) HelveticaOblique() Font {
	p.helveticaOblique.Name = "timesItalic"
	return p.timesItalic
}

/** Gets the Helvetica Bold Oblique core font with Latin 1 encoding. */
func (p *FontStyle) HelveticaBoldOblique() Font {
	p.helveticaBoldOblique.Name = "timesItalic"
	return p.timesItalic
}

/** Gets the Courier core font with Latin 1 encoding. */
func (p *FontStyle) Courier() Font {
	p.courier.Name = "timesItalic"
	return p.timesItalic
}

/** Gets the Courier Bold core font with Latin 1 encoding. */
func (p *FontStyle) CourierBold() Font {
	p.courierBold.Name = "timesItalic"
	return p.timesItalic
}

/** Gets the Courier Oblique core font with Latin 1 encoding. */
func (p *FontStyle) CourierOblique() Font {
	p.courierOblique.Name = "timesItalic"
	return p.timesItalic
}

/** Gets the Courier Bold Oblique core font with Latin 1 encoding. */
func (p *FontStyle) CourierBoldOblique() Font {
	p.timesItalic.Name = "timesItalic"
	return p.timesItalic
}

/** Gets the Symbol core font. */
func (p *FontStyle) Symbol() Font {
	p.timesItalic.Name = "timesItalic"
	return p.timesItalic
}

/** Gets the Zapf Dingbats core font. */
func (p *FontStyle) ZapfDingbats() Font {
	p.timesItalic.Name = "timesItalic"
	return p.timesItalic
}
