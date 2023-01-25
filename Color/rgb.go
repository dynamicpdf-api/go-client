package color

import "strconv"

/// <summary>
    //presents an RGB color.
    /// </summary>
	type RgbColor struct {
		Color
		red   float32
		green float32
		blue  float32
	}
	func newRgbColor(color string) *RgbColor {
		var ep RgbColor
		ep.colorString = color
		return &ep
	}
func NewRgbColorDefault() *RgbColor {
	var ep RgbColor
	return &ep
}

/// <summary>
        /// Initializes a new instance of the <see cref="RgbColor"/> class.
        /// </summary>
        /// <param name="red">The red intensity.</param>
        /// <param name="green">The green intensity.</param>
        /// <param name="blue">The blue intensity.</param>
///		</summary>
func NewRgbColor(red float32, green float32, blue float32) *RgbColor {
	var ep RgbColor
	ep.red = red
	ep.green = green
	ep.blue = blue
	ep.colorString = "rgb(" + strconv.FormatFloat(float64(ep.red), 'f', -1, 64) + "," + strconv.FormatFloat(float64(ep.green), 'f', -1, 64) + "," + strconv.FormatFloat(float64(blue), 'f', -1, 64) + ")"
	return &ep
}

func (p *RgbColor) RgbColorString() string {
	return p.colorString
}

func (p *RgbColor) setRgbColorString(value string) {
	p.colorString = value
}

/// <summary>Gets the color red.</summary>
func (p *RgbColor) Red() RgbColor {
	ep := newRgbColor("Red")
	return *ep
}

/// <summary>Gets the color blue.</summary>
func (p *RgbColor) Blue() RgbColor {
	ep := newRgbColor("Blue")
	return *ep
}
/// <summary>Gets the color green.</summary>
func (p *RgbColor) Green() RgbColor {
	ep := newRgbColor("Green")
	return *ep
}
/// <summary>Gets the color black.</summary>
func (p *RgbColor) Black() RgbColor {
	ep := newRgbColor("Black")
	return *ep
}
/// <summary>Gets the color silver.</summary>
func (p *RgbColor) Silver() RgbColor {
	ep := newRgbColor("Silver")
	return *ep
}
/// <summary>Gets the color dark gray.</summary>
func (p *RgbColor) DarkGray() RgbColor {
	ep := newRgbColor("DarkGray")
	return *ep
}
/// <summary>Gets the color gray.</summary>
func (p *RgbColor) Gray() RgbColor {
	ep := newRgbColor("Gray")
	return *ep
}

        /// <summary>Gets the color dim gray.</summary>
func (p *RgbColor) DimGray() RgbColor {
	ep := newRgbColor("DimGray")
	return *ep
}
/// <summary>Gets the color white.</summary>
func (p *RgbColor) White() RgbColor {
	ep := newRgbColor("White")
	return *ep
}
/// <summary>Gets the color lime.</summary>
func (p *RgbColor) Lime() RgbColor {
	ep := newRgbColor("Lime")
	return *ep
}
/// <summary>Gets the color aqua.</summary>
func (p *RgbColor) Aqua() RgbColor {
	ep := newRgbColor("Aqua")
	return *ep
}
/// <summary>Gets the color purple.</summary>
func (p *RgbColor) Purple() RgbColor {
	ep := newRgbColor("Purple")
	return *ep
}
/// <summary>Gets the color cyan.</summary>
func (p *RgbColor) Cyan() RgbColor {
	ep := newRgbColor("Cyan")
	return *ep
}
/// <summary>Gets the color magenta.</summary>
func (p *RgbColor) Magenta() RgbColor {
	ep := newRgbColor("Magenta")
	return *ep
}
 /// <summary>Gets the color yellow.</summary>
func (p *RgbColor) Yellow() RgbColor {
	ep := newRgbColor("Yellow")
	return *ep
}
/// <summary>Gets the color alice blue.</summary>
func (p *RgbColor) AliceBlue() RgbColor {
	ep := newRgbColor("AliceBlue")
	return *ep
}
/// <summary>Gets the color antique white.</summary>
func (p *RgbColor) AntiqueWhite() RgbColor {
	ep := newRgbColor("AntiqueWhite")
	return *ep
}
/// <summary>Gets the color aquamarine.</summary>
func (p *RgbColor) Aquamarine() RgbColor {
	ep := newRgbColor("Aquamarine")
	return *ep
}
/// <summary>Gets the color azure.</summary>
func (p *RgbColor) Azure() RgbColor {
	ep := newRgbColor("Azure")
	return *ep
}
/// <summary>Gets the color beige.</summary>
func (p *RgbColor) Beige() RgbColor {
	ep := newRgbColor("Beige")
	return *ep
}
/// <summary>Gets the color bisque.</summary>
func (p *RgbColor) Bisque() RgbColor {
	ep := newRgbColor("Bisque")
	return *ep
}
/// <summary>Gets the color blanched almond.</summary>
func (p *RgbColor) BlanchedAlmond() RgbColor {
	ep := newRgbColor("BlanchedAlmond")
	return *ep
}
 /// <summary>Gets the color blue violet.</summary>
func (p *RgbColor) BlueViolet() RgbColor {
	ep := newRgbColor("BlueViolet")
	return *ep
}
/// <summary>Gets the color brown.</summary>
func (p *RgbColor) Brown() RgbColor {
	ep := newRgbColor("Brown")
	return *ep
}
/// <summary>Gets the color burly wood.</summary>
func (p *RgbColor) BurlyWood() RgbColor {
	ep := newRgbColor("BurlyWood")
	return *ep
}
/// <summary>Gets the color cadet blue.</summary>
func (p *RgbColor) CadetBlue() RgbColor {
	ep := newRgbColor("CadetBlue")
	return *ep
}
/// <summary>Gets the color chartreuse.</summary>
func (p *RgbColor) Chartreuse() RgbColor {
	ep := newRgbColor("Chartreuse")
	return *ep
}
/// <summary>Gets the color chocolate.</summary>
func (p *RgbColor) Chocolate() RgbColor {
	ep := newRgbColor("Chocolate")
	return *ep
}
/// <summary>Gets the color coral.</summary>
func (p *RgbColor) Coral() RgbColor {
	ep := newRgbColor("Coral")
	return *ep
}
 /// <summary>Gets the color cornflower blue.</summary>
func (p *RgbColor) CornflowerBlue() RgbColor {
	ep := newRgbColor("CornflowerBlue")
	return *ep
}
/// <summary>Gets the color cornsilk.</summary>
func (p *RgbColor) Cornsilk() RgbColor {
	ep := newRgbColor("Cornsilk")
	return *ep
}
/// <summary>Gets the color crimson.</summary>
func (p *RgbColor) Crimson() RgbColor {
	ep := newRgbColor("Crimson")
	return *ep
}
/// <summary>Gets the color dark blue.</summary>
func (p *RgbColor) DarkBlue() RgbColor {
	ep := newRgbColor("DarkBlue")
	return *ep
}
/// <summary>Gets the color dark cyan.</summary>
func (p *RgbColor) DarkCyan() RgbColor {
	ep := newRgbColor("DarkCyan")
	return *ep
}
/// <summary>Gets the color dark golden rod.</summary>
func (p *RgbColor) DarkGoldenRod() RgbColor {
	ep := newRgbColor("DarkGoldenRod")
	return *ep
}
/// <summary>Gets the color dark green.</summary>
func (p *RgbColor) DarkGreen() RgbColor {
	ep := newRgbColor("DarkGreen")
	return *ep
}
/// <summary>Gets the color dark khaki.</summary>
func (p *RgbColor) DarkKhaki() RgbColor {
	ep := newRgbColor("DarkKhaki")
	return *ep
}
/// <summary>Gets the color dark magenta.</summary>
func (p *RgbColor) DarkMagenta() RgbColor {
	ep := newRgbColor("DarkMagenta")
	return *ep
}
/// <summary>Gets the color dark olive green.</summary>
func (p *RgbColor) DarkOliveGreen() RgbColor {
	ep := newRgbColor("DarkOliveGreen")
	return *ep
}
/// <summary>Gets the color dark orange.</summary>
func (p *RgbColor) DarkOrange() RgbColor {
	ep := newRgbColor("DarkOrange")
	return *ep
}
 /// <summary>Gets the color dark orchid.</summary>
func (p *RgbColor) DarkOrchid() RgbColor {
	ep := newRgbColor("DarkOrchid")
	return *ep
}
/// <summary>Gets the color dark red.</summary>
func (p *RgbColor) DarkRed() RgbColor {
	ep := newRgbColor("DarkRed")
	return *ep
}
 /// <summary>Gets the color dark salmon.</summary>
func (p *RgbColor) DarkSalmon() RgbColor {
	ep := newRgbColor("DarkSalmon")
	return *ep
}
/// <summary>Gets the color dark sea green.</summary>
func (p *RgbColor) DarkSeaGreen() RgbColor {
	ep := newRgbColor("DarkSeaGreen")
	return *ep
}
/// <summary>Gets the color dark slate blue.</summary>
func (p *RgbColor) DarkSlateBlue() RgbColor {
	ep := newRgbColor("DarkSlateBlue")
	return *ep
}
 /// <summary>Gets the color dark slate gray.</summary>
func (p *RgbColor) DarkSlateGray() RgbColor {
	ep := newRgbColor("DarkSlateGray")
	return *ep
}
/// <summary>Gets the color dark turquoise.</summary>
func (p *RgbColor) DarkTurquoise() RgbColor {
	ep := newRgbColor("DarkTurquoise")
	return *ep
}
/// <summary>Gets the color dark violet.</summary>
func (p *RgbColor) DarkViolet() RgbColor {
	ep := newRgbColor("DarkViolet")
	return *ep
}
/// <summary>Gets the color deep pink.</summary>
func (p *RgbColor) DeepPink() RgbColor {
	ep := newRgbColor("DeepPink")
	return *ep
}
/// <summary>Gets the color deep sky blue.</summary>
func (p *RgbColor) DeepSkyBlue() RgbColor {
	ep := newRgbColor("DeepSkyBlue")
	return *ep
}
/// <summary>Gets the color dodger blue.</summary>
func (p *RgbColor) DodgerBlue() RgbColor {
	ep := newRgbColor("DodgerBlue")
	return *ep
}
/// <summary>Gets the color feldspar.</summary>
func (p *RgbColor) Feldspar() RgbColor {
	ep := newRgbColor("Feldspar")
	return *ep
}
/// <summary>Gets the color fire brick.</summary>
func (p *RgbColor) FireBrick() RgbColor {
	ep := newRgbColor("FireBrick")
	return *ep
}
/// <summary>Gets the color floral white.</summary>
func (p *RgbColor) FloralWhite() RgbColor {
	ep := newRgbColor("FloralWhite")
	return *ep
}
/// <summary>Gets the color forest green.</summary>
func (p *RgbColor) ForestGreen() RgbColor {
	ep := newRgbColor("ForestGreen")
	return *ep
}
/// <summary>Gets the color fuchsia.</summary>
func (p *RgbColor) Fuchsia() RgbColor {
	ep := newRgbColor("Fuchsia")
	return *ep
}
/// <summary>Gets the color ghost white.</summary>
func (p *RgbColor) GhostWhite() RgbColor {
	ep := newRgbColor("GhostWhite")
	return *ep
}
/// <summary>Gets the color gold.</summary>
func (p *RgbColor) Gold() RgbColor {
	ep := newRgbColor("Gold")
	return *ep
}
/// <summary>Gets the color golden rod.</summary>
func (p *RgbColor) GoldenRod() RgbColor {
	ep := newRgbColor("GoldenRod")
	return *ep
}
/// <summary>Gets the color green yellow.</summary>
func (p *RgbColor) GreenYellow() RgbColor {
	ep := newRgbColor("GreenYellow")
	return *ep
}
/// <summary>Gets the color honey dew.</summary>
func (p *RgbColor) HoneyDew() RgbColor {
	ep := newRgbColor("HoneyDew")
	return *ep
}
/// <summary>Gets the color hot pink.</summary>
func (p *RgbColor) HotPink() RgbColor {
	ep := newRgbColor("HotPink")
	return *ep
}
/// <summary>Gets the color indian red.</summary>
func (p *RgbColor) IndianRed() RgbColor {
	ep := newRgbColor("IndianRed")
	return *ep
}
/// <summary>Gets the color indigo.</summary>
func (p *RgbColor) Indigo() RgbColor {
	ep := newRgbColor("Indigo")
	return *ep
}
/// <summary>Gets the color ivory.</summary>
func (p *RgbColor) Ivory() RgbColor {
	ep := newRgbColor("Ivory")
	return *ep
}
/// <summary>Gets the color khaki.</summary>
func (p *RgbColor) Khaki() RgbColor {
	ep := newRgbColor("Khaki")
	return *ep
}
/// <summary>Gets the color lavender.</summary>
func (p *RgbColor) Lavender() RgbColor {
	ep := newRgbColor("Lavender")
	return *ep
}
/// <summary>Gets the color lavender blush.</summary>
func (p *RgbColor) LavenderBlush() RgbColor {
	ep := newRgbColor("LavenderBlush")
	return *ep
}
/// <summary>Gets the color lawn Green.</summary>
func (p *RgbColor) LawnGreen() RgbColor {
	ep := newRgbColor("LawnGreen")
	return *ep
}
/// <summary>Gets the color lemon chiffon.</summary>
func (p *RgbColor) LemonChiffon() RgbColor {
	ep := newRgbColor("LemonChiffon")
	return *ep
}
/// <summary>Gets the color light blue.</summary>
func (p *RgbColor) LightBlue() RgbColor {
	ep := newRgbColor("LightBlue")
	return *ep
}
/// <summary>Gets the color light coral.</summary>
func (p *RgbColor) LightCoral() RgbColor {
	ep := newRgbColor("LightCoral")
	return *ep
}
/// <summary>Gets the color light cyan.</summary>
func (p *RgbColor) LightCyan() RgbColor {
	ep := newRgbColor("LightCyan")
	return *ep
}
/// <summary>Gets the color light golden rod yellow.</summary>
func (p *RgbColor) LightGoldenRodYellow() RgbColor {
	ep := newRgbColor("LightGoldenRodYellow")
	return *ep
}
/// <summary>Gets the color light grey.</summary>
func (p *RgbColor) LightGrey() RgbColor {
	ep := newRgbColor("LightGrey")
	return *ep
}
/// <summary>Gets the color light green.</summary>
func (p *RgbColor) LightGreen() RgbColor {
	ep := newRgbColor("LightGreen")
	return *ep
}
/// <summary>Gets the color light pink.</summary>
func (p *RgbColor) LightPink() RgbColor {
	ep := newRgbColor("LightPink")
	return *ep
}
/// <summary>Gets the color light salmon.</summary>
func (p *RgbColor) LightSalmon() RgbColor {
	ep := newRgbColor("LightSalmon")
	return *ep
}
/// <summary>Gets the color light sea green.</summary>
func (p *RgbColor) LightSeaGreen() RgbColor {
	ep := newRgbColor("LightSeaGreen")
	return *ep
}
/// <summary>Gets the color light sky blue.</summary>
func (p *RgbColor) LightSkyBlue() RgbColor {
	ep := newRgbColor("LightSkyBlue")
	return *ep
}
/// <summary>Gets the color light slate blue.</summary>
func (p *RgbColor) LightSlateBlue() RgbColor {
	ep := newRgbColor("LightSlateBlue")
	return *ep
}
/// <summary>Gets the color light slate gray.</summary>
func (p *RgbColor) LightSlateGray() RgbColor {
	ep := newRgbColor("LightSlateGray")
	return *ep
}
 /// <summary>Gets the color light steel blue.</summary>
func (p *RgbColor) LightSteelBlue() RgbColor {
	ep := newRgbColor("LightSteelBlue")
	return *ep
}
/// <summary>Gets the color light yellow.</summary>
func (p *RgbColor) LightYellow() RgbColor {
	ep := newRgbColor("LightYellow")
	return *ep
}
/// <summary>Gets the color lime green.</summary>
func (p *RgbColor) LimeGreen() RgbColor {
	ep := newRgbColor("LimeGreen")
	return *ep
}
/// <summary>Gets the color linen.</summary>
func (p *RgbColor) Linen() RgbColor {
	ep := newRgbColor("Linen")
	return *ep
}
/// <summary>Gets the color maroon.</summary>
func (p *RgbColor) Maroon() RgbColor {
	ep := newRgbColor("Maroon")
	return *ep
}
/// <summary>Gets the color medium aqua marine.</summary>
func (p *RgbColor) MediumAquaMarine() RgbColor {
	ep := newRgbColor("MediumAquaMarine")
	return *ep
}

        /// <summary>Gets the color medium blue.</summary>
func (p *RgbColor) MediumBlue() RgbColor {
	ep := newRgbColor("MediumBlue")
	return *ep
}
/// <summary>Gets the color medium orchid.</summary>
func (p *RgbColor) MediumOrchid() RgbColor {
	ep := newRgbColor("MediumOrchid")
	return *ep
}
/// <summary>Gets the color medium purple.</summary>
func (p *RgbColor) MediumPurple() RgbColor {
	ep := newRgbColor("MediumPurple")
	return *ep
}
/// <summary>Gets the color medium sea green.</summary>
func (p *RgbColor) MediumSeaGreen() RgbColor {
	ep := newRgbColor("MediumSeaGreen")
	return *ep
}
/// <summary>Gets the color medium slate blue.</summary>
func (p *RgbColor) MediumSlateBlue() RgbColor {
	ep := newRgbColor("MediumSlateBlue")
	return *ep
}
/// <summary>Gets the color medium spring green.</summary>
func (p *RgbColor) MediumSpringGreen() RgbColor {
	ep := newRgbColor("MediumSpringGreen")
	return *ep
}
/// <summary>Gets the color medium turquoise.</summary>
func (p *RgbColor) MediumTurquoise() RgbColor {
	ep := newRgbColor("MediumTurquoise")
	return *ep
}
/// <summary>Gets the color medium violet red.</summary>
func (p *RgbColor) MediumVioletRed() RgbColor {
	ep := newRgbColor("MediumVioletRed")
	return *ep
}
/// <summary>Gets the color midnight blue.</summary>
func (p *RgbColor) MidnightBlue() RgbColor {
	ep := newRgbColor("MidnightBlue")
	return *ep
}
/// <summary>Gets the color mint cream.</summary>
func (p *RgbColor) MintCream() RgbColor {
	ep := newRgbColor("MintCream")
	return *ep
}
/// <summary>Gets the color misty rose.</summary>
func (p *RgbColor) MistyRose() RgbColor {
	ep := newRgbColor("MistyRose")
	return *ep
}
/// <summary>Gets the color moccasin.</summary>
func (p *RgbColor) Moccasin() RgbColor {
	ep := newRgbColor("Moccasin")
	return *ep
}
/// <summary>Gets the color navajo white.</summary>
func (p *RgbColor) NavajoWhite() RgbColor {
	ep := newRgbColor("NavajoWhite")
	return *ep
}
/// <summary>Gets the color navy.</summary>
func (p *RgbColor) Navy() RgbColor {
	ep := newRgbColor("Navy")
	return *ep
}
/// <summary>Gets the color old lace.</summary>
func (p *RgbColor) OldLace() RgbColor {
	ep := newRgbColor("OldLace")
	return *ep
}
/// <summary>Gets the color olive.</summary>
func (p *RgbColor) Olive() RgbColor {
	ep := newRgbColor("Olive")
	return *ep
}
/// <summary>Gets the color olive drab.</summary>
func (p *RgbColor) OliveDrab() RgbColor {
	ep := newRgbColor("OliveDrab")
	return *ep
}
/// <summary>Gets the color gainsboro.</summary>
func (p *RgbColor) Gainsboro() RgbColor {
	ep := newRgbColor("Gainsboro")
	return *ep
}
/// <summary>Gets the color orange.</summary>
func (p *RgbColor) Orange() RgbColor {
	ep := newRgbColor("Orange")
	return *ep
}
/// <summary>Gets the color orange red.</summary>
func (p *RgbColor) OrangeRed() RgbColor {
	ep := newRgbColor("OrangeRed")
	return *ep
}
/// <summary>Gets the color orchid.</summary>
func (p *RgbColor) Orchid() RgbColor {
	ep := newRgbColor("Orchid")
	return *ep
}
/// <summary>Gets the color pale golden rod.</summary>
func (p *RgbColor) PaleGoldenRod() RgbColor {
	ep := newRgbColor("PaleGoldenRod")
	return *ep
}
/// <summary>Gets the color pale green.</summary>
func (p *RgbColor) PaleGreen() RgbColor {
	ep := newRgbColor("PaleGreen")
	return *ep
}
/// <summary>Gets the color pale turquoise.</summary>
func (p *RgbColor) PaleTurquoise() RgbColor {
	ep := newRgbColor("PaleTurquoise")
	return *ep
}
/// <summary>Gets the color pale violet red.</summary>
func (p *RgbColor) PaleVioletRed() RgbColor {
	ep := newRgbColor("PaleVioletRed")
	return *ep
}
/// <summary>Gets the color papaya whip.</summary>
func (p *RgbColor) PapayaWhip() RgbColor {
	ep := newRgbColor("PapayaWhip")
	return *ep
}
/// <summary>Gets the color peach puff.</summary>
func (p *RgbColor) PeachPuff() RgbColor {
	ep := newRgbColor("PeachPuff")
	return *ep
}

        /// <summary>Gets the color peru.</summary>
func (p *RgbColor) Peru() RgbColor {
	ep := newRgbColor("Peru")
	return *ep
}
/// <summary>Gets the color pink.</summary>
func (p *RgbColor) Pink() RgbColor {
	ep := newRgbColor("Pink")
	return *ep
}
/// <summary>Gets the color plum.</summary>
func (p *RgbColor) Plum() RgbColor {
	ep := newRgbColor("Plum")
	return *ep
}
/// <summary>Gets the color powder blue.</summary>
func (p *RgbColor) PowderBlue() RgbColor {
	ep := newRgbColor("PowderBlue")
	return *ep
}
/// <summary>Gets the color rosy brown.</summary>
func (p *RgbColor) RosyBrown() RgbColor {
	ep := newRgbColor("RosyBrown")
	return *ep
}
/// <summary>Gets the color royal blue.</summary>
func (p *RgbColor) RoyalBlue() RgbColor {
	ep := newRgbColor("RoyalBlue")
	return *ep
}
/// <summary>Gets the color saddle brown.</summary>
func (p *RgbColor) SaddleBrown() RgbColor {
	ep := newRgbColor("SaddleBrown")
	return *ep
}
/// <summary>Gets the color salmon.</summary>
func (p *RgbColor) Salmon() RgbColor {
	ep := newRgbColor("Salmon")
	return *ep
}
/// <summary>Gets the color sandy brown.</summary>
func (p *RgbColor) SandyBrown() RgbColor {
	ep := newRgbColor("SandyBrown")
	return *ep
}
/// <summary>Gets the color sea green.</summary>
func (p *RgbColor) SeaGreen() RgbColor {
	ep := newRgbColor("SeaGreen")
	return *ep
}
/// <summary>Gets the color sea shell.</summary>
func (p *RgbColor) SeaShell() RgbColor {
	ep := newRgbColor("SeaShell")
	return *ep
}
/// <summary>Gets the color sienna.</summary>
func (p *RgbColor) Sienna() RgbColor {
	ep := newRgbColor("Sienna")
	return *ep
}
/// <summary>Gets the color sky blue.</summary>
func (p *RgbColor) SkyBlue() RgbColor {
	ep := newRgbColor("SkyBlue")
	return *ep
}
/// <summary>Gets the color slate blue.</summary>
func (p *RgbColor) SlateBlue() RgbColor {
	ep := newRgbColor("SlateBlue")
	return *ep
}
/// <summary>Gets the color slate gray.</summary>
func (p *RgbColor) SlateGray() RgbColor {
	ep := newRgbColor("SlateGray")
	return *ep
}
/// <summary>Gets the color snow.</summary>
func (p *RgbColor) Snow() RgbColor {
	ep := newRgbColor("Snow")
	return *ep
}
/// <summary>Gets the color spring green.</summary>
func (p *RgbColor) SpringGreen() RgbColor {
	ep := newRgbColor("SpringGreen")
	return *ep
}
/// <summary>Gets the color steel blue.</summary>
func (p *RgbColor) SteelBlue() RgbColor {
	ep := newRgbColor("SteelBlue")
	return *ep
}
/// <summary>Gets the color Tan.</summary>
func (p *RgbColor) Tan() RgbColor {
	ep := newRgbColor("Tan")
	return *ep
}
/// <summary>Gets the color teal.</summary>
func (p *RgbColor) Teal() RgbColor {
	ep := newRgbColor("Teal")
	return *ep
}
/// <summary>Gets the color thistle.</summary>
func (p *RgbColor) Thistle() RgbColor {
	ep := newRgbColor("Thistle")
	return *ep
}
/// <summary>Gets the color tomato.</summary>
func (p *RgbColor) Tomato() RgbColor {
	ep := newRgbColor("Tomato")
	return *ep
}
/// <summary>Gets the color turquoise.</summary>
func (p *RgbColor) Turquoise() RgbColor {
	ep := newRgbColor("Turquoise")
	return *ep
}
/// <summary>Gets the color violet.</summary>
func (p *RgbColor) Violet() RgbColor {
	ep := newRgbColor("Violet")
	return *ep
}
/// <summary>Gets the color violet red.</summary>
func (p *RgbColor) VioletRed() RgbColor {
	ep := newRgbColor("VioletRed")
	return *ep
}

/// <summary>Gets the color wheat.</summary>
func (p *RgbColor) Wheat() RgbColor {
	ep := newRgbColor("Wheat")
	return *ep
}

/// <summary>Gets the color white smoke.</summary>
func (p *RgbColor) WhiteSmoke() RgbColor {
	ep := newRgbColor("WhiteSmoke")
	return *ep
}

/// <summary>Gets the color yellow green.</summary>
func (p *RgbColor) YellowGreen() RgbColor {
	ep := newRgbColor("YellowGreen")
	return *ep
}
