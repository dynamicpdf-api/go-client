package font

import (
	"strconv"

	"github.com/dynamicpdf-api/go-client/resource"
	"github.com/google/uuid"
)

// Represents font.
type Font struct {
	Resource resource.Resource `json:"-"`

	ResourceName string `json:"resourceName,omitempty"`

	Embed bool `json:"embed,omitempty"`

	Subset bool `json:"subset,omitempty"`

	Name string `json:"name,omitempty"`
}

// Initializes a new instance of the `Font` class
func NewFont() *Font {
	var p Font
	p.Name = uuid.NewString()
	return &p
}

// NewFontResource initializes a new Font with given font file provided with path.
func NewFontResource(path string, resourceName string) *Font {
	var p Font
	p.Resource = resource.NewResourceWithPath(path, resourceName)
	p.Name = uuid.NewString()
	p.ResourceName = resourceName
	return &p
}

// Gets the Times Italic core font with Latin 1 encoding.
func TimesItalic() Font {
	return Font{Name: "timesItalic"}
}

// Gets the Times Roman core font with Latin 1 encoding.
func TimesRoman() Font {
	return Font{Name: "timesRoman"}
}

// Gets the Times Bold core font with Latin 1 encoding.
func TimesBold() Font {
	return Font{Name: "timesBold"}
}

// Gets the Times Bold Italic core font with Latin 1 encoding.
func TimesBoldItalic() Font {
	return Font{Name: "timesBoldItalic"}
}

// Gets the Helvetica core font with Latin 1 encoding.
func Helvetica() Font {
	return Font{Name: "helvetica"}
}

// Gets the Helvetica Bold core font with Latin 1 encoding.
func HelveticaBold() Font {
	return Font{Name: "helveticaBold"}
}

// Gets the Helvetica Oblique core font with Latin 1 encoding.
func HelveticaOblique() Font {
	return Font{Name: "helveticaOblique"}
}

/** Gets the Helvetica Bold Oblique core font with Latin 1 encoding. */
func HelveticaBoldOblique() Font {
	return Font{Name: "helveticaBoldOblique"}
}

// Gets the Courier core font with Latin 1 encoding.
func Courier() Font {
	return Font{Name: "courier"}
}

// Gets the Courier Bold core font with Latin 1 encoding.
func CourierBold() Font {
	return Font{Name: "courierBold"}
}

// Gets the Courier Oblique core font with Latin 1 encoding.
func CourierOblique() Font {
	return Font{Name: "courierOblique"}
}

// Gets the Courier Bold Oblique core font with Latin 1 encoding.
func CourierBoldOblique() Font {
	return Font{Name: "courierBoldOblique"}
}

// Gets the Symbol core font.
func Symbol() Font {
	return Font{Name: "symbol"}
}

// Gets the Zapf Dingbats core font.
func ZapfDingbats() Font {
	return Font{Name: "zapfDingbats"}
}

func getGoogleFontText(name string, weight int, italic bool) string {
	if italic == true {
		return name + ":" + strconv.Itoa(weight) + "italic"
	} else {
		return name + ":" + strconv.Itoa(weight)
	}
}

func Google(fontName string) *Font {
	font := NewFont()
	font.Name = fontName
	return font
}

func GoogleFontWithWeight(fontName string, weight int, italic bool) *Font {
	font := NewFont()
	font.Name = getGoogleFontText(fontName, weight, italic)
	return font
}

func Global(fontName string) *Font {
	font := NewFont()
	font.Name = fontName
	return font
}
