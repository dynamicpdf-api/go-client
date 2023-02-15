package color

import "strconv"

// Represents an RGB color.
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

/*
  Initializes a new instance of the 'RgbColor' class.
   * @param {float} red. The red intensity.
   * @param {float} green. The green intensity.
   * @param {float} blue. The blue intensity.
*/
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

// Gets the color red.
func (p *RgbColor) Red() RgbColor {
	ep := newRgbColor("Red")
	return *ep
}

// Gets the color blue.
func (p *RgbColor) Blue() RgbColor {
	ep := newRgbColor("Blue")
	return *ep
}

// Gets the color green.
func (p *RgbColor) Green() RgbColor {
	ep := newRgbColor("Green")
	return *ep
}

// Gets the color black.
func (p *RgbColor) Black() RgbColor {
	ep := newRgbColor("Black")
	return *ep
}

// Gets the color silver.
func (p *RgbColor) Silver() RgbColor {
	ep := newRgbColor("Silver")
	return *ep
}

// Gets the color dark gray.
func (p *RgbColor) DarkGray() RgbColor {
	ep := newRgbColor("DarkGray")
	return *ep
}

// Gets the color gray.
func (p *RgbColor) Gray() RgbColor {
	ep := newRgbColor("Gray")
	return *ep
}

// Gets the color dim gray.
func (p *RgbColor) DimGray() RgbColor {
	ep := newRgbColor("DimGray")
	return *ep
}

// Gets the color white.
func (p *RgbColor) White() RgbColor {
	ep := newRgbColor("White")
	return *ep
}

// Gets the color lime.
func (p *RgbColor) Lime() RgbColor {
	ep := newRgbColor("Lime")
	return *ep
}

// Gets the color aqua.
func (p *RgbColor) Aqua() RgbColor {
	ep := newRgbColor("Aqua")
	return *ep
}

// Gets the color purple.
func (p *RgbColor) Purple() RgbColor {
	ep := newRgbColor("Purple")
	return *ep
}

// Gets the color cyan.
func (p *RgbColor) Cyan() RgbColor {
	ep := newRgbColor("Cyan")
	return *ep
}

// Gets the color magenta.
func (p *RgbColor) Magenta() RgbColor {
	ep := newRgbColor("Magenta")
	return *ep
}

// Gets the color yellow.
func (p *RgbColor) Yellow() RgbColor {
	ep := newRgbColor("Yellow")
	return *ep
}

// Gets the color alice blue.
func (p *RgbColor) AliceBlue() RgbColor {
	ep := newRgbColor("AliceBlue")
	return *ep
}

// Gets the color antique white.
func (p *RgbColor) AntiqueWhite() RgbColor {
	ep := newRgbColor("AntiqueWhite")
	return *ep
}

// Gets the color aquamarine.
func (p *RgbColor) Aquamarine() RgbColor {
	ep := newRgbColor("Aquamarine")
	return *ep
}

// Gets the color azure.
func (p *RgbColor) Azure() RgbColor {
	ep := newRgbColor("Azure")
	return *ep
}

// Gets the color beige.
func (p *RgbColor) Beige() RgbColor {
	ep := newRgbColor("Beige")
	return *ep
}

// Gets the color bisque.
func (p *RgbColor) Bisque() RgbColor {
	ep := newRgbColor("Bisque")
	return *ep
}

// Gets the color blanched almond.
func (p *RgbColor) BlanchedAlmond() RgbColor {
	ep := newRgbColor("BlanchedAlmond")
	return *ep
}

// Gets the color blue violet.
func (p *RgbColor) BlueViolet() RgbColor {
	ep := newRgbColor("BlueViolet")
	return *ep
}

// Gets the color brown.
func (p *RgbColor) Brown() RgbColor {
	ep := newRgbColor("Brown")
	return *ep
}

// Gets the color burly wood.
func (p *RgbColor) BurlyWood() RgbColor {
	ep := newRgbColor("BurlyWood")
	return *ep
}

// Gets the color cadet blue.
func (p *RgbColor) CadetBlue() RgbColor {
	ep := newRgbColor("CadetBlue")
	return *ep
}

// Gets the color chartreuse.
func (p *RgbColor) Chartreuse() RgbColor {
	ep := newRgbColor("Chartreuse")
	return *ep
}

// Gets the color chocolate.
func (p *RgbColor) Chocolate() RgbColor {
	ep := newRgbColor("Chocolate")
	return *ep
}

// Gets the color coral.
func (p *RgbColor) Coral() RgbColor {
	ep := newRgbColor("Coral")
	return *ep
}

// Gets the color cornflower blue.
func (p *RgbColor) CornflowerBlue() RgbColor {
	ep := newRgbColor("CornflowerBlue")
	return *ep
}

// Gets the color cornsilk.
func (p *RgbColor) Cornsilk() RgbColor {
	ep := newRgbColor("Cornsilk")
	return *ep
}

// Gets the color crimson.
func (p *RgbColor) Crimson() RgbColor {
	ep := newRgbColor("Crimson")
	return *ep
}

// Gets the color dark blue.
func (p *RgbColor) DarkBlue() RgbColor {
	ep := newRgbColor("DarkBlue")
	return *ep
}

// Gets the color dark cyan.
func (p *RgbColor) DarkCyan() RgbColor {
	ep := newRgbColor("DarkCyan")
	return *ep
}

// Gets the color dark golden rod.
func (p *RgbColor) DarkGoldenRod() RgbColor {
	ep := newRgbColor("DarkGoldenRod")
	return *ep
}

// Gets the color dark green.
func (p *RgbColor) DarkGreen() RgbColor {
	ep := newRgbColor("DarkGreen")
	return *ep
}

// Gets the color dark khaki.
func (p *RgbColor) DarkKhaki() RgbColor {
	ep := newRgbColor("DarkKhaki")
	return *ep
}

// Gets the color dark magenta.
func (p *RgbColor) DarkMagenta() RgbColor {
	ep := newRgbColor("DarkMagenta")
	return *ep
}

// Gets the color dark olive green.
func (p *RgbColor) DarkOliveGreen() RgbColor {
	ep := newRgbColor("DarkOliveGreen")
	return *ep
}

// Gets the color dark orange.
func (p *RgbColor) DarkOrange() RgbColor {
	ep := newRgbColor("DarkOrange")
	return *ep
}

// Gets the color dark orchid.
func (p *RgbColor) DarkOrchid() RgbColor {
	ep := newRgbColor("DarkOrchid")
	return *ep
}

// Gets the color dark red.
func (p *RgbColor) DarkRed() RgbColor {
	ep := newRgbColor("DarkRed")
	return *ep
}

// Gets the color dark salmon.
func (p *RgbColor) DarkSalmon() RgbColor {
	ep := newRgbColor("DarkSalmon")
	return *ep
}

// Gets the color dark sea green.
func (p *RgbColor) DarkSeaGreen() RgbColor {
	ep := newRgbColor("DarkSeaGreen")
	return *ep
}

// Gets the color dark slate blue.
func (p *RgbColor) DarkSlateBlue() RgbColor {
	ep := newRgbColor("DarkSlateBlue")
	return *ep
}

// Gets the color dark slate gray.
func (p *RgbColor) DarkSlateGray() RgbColor {
	ep := newRgbColor("DarkSlateGray")
	return *ep
}

// Gets the color dark turquoise.
func (p *RgbColor) DarkTurquoise() RgbColor {
	ep := newRgbColor("DarkTurquoise")
	return *ep
}

// Gets the color dark violet.
func (p *RgbColor) DarkViolet() RgbColor {
	ep := newRgbColor("DarkViolet")
	return *ep
}

// Gets the color deep pink.
func (p *RgbColor) DeepPink() RgbColor {
	ep := newRgbColor("DeepPink")
	return *ep
}

// Gets the color deep sky blue.
func (p *RgbColor) DeepSkyBlue() RgbColor {
	ep := newRgbColor("DeepSkyBlue")
	return *ep
}

// Gets the color dodger blue.
func (p *RgbColor) DodgerBlue() RgbColor {
	ep := newRgbColor("DodgerBlue")
	return *ep
}

// Gets the color feldspar.
func (p *RgbColor) Feldspar() RgbColor {
	ep := newRgbColor("Feldspar")
	return *ep
}

// Gets the color fire brick.
func (p *RgbColor) FireBrick() RgbColor {
	ep := newRgbColor("FireBrick")
	return *ep
}

// Gets the color floral white.
func (p *RgbColor) FloralWhite() RgbColor {
	ep := newRgbColor("FloralWhite")
	return *ep
}

// Gets the color forest green.
func (p *RgbColor) ForestGreen() RgbColor {
	ep := newRgbColor("ForestGreen")
	return *ep
}

// Gets the color fuchsia.
func (p *RgbColor) Fuchsia() RgbColor {
	ep := newRgbColor("Fuchsia")
	return *ep
}

// Gets the color ghost white.
func (p *RgbColor) GhostWhite() RgbColor {
	ep := newRgbColor("GhostWhite")
	return *ep
}

// Gets the color gold.
func (p *RgbColor) Gold() RgbColor {
	ep := newRgbColor("Gold")
	return *ep
}

// Gets the color golden rod.
func (p *RgbColor) GoldenRod() RgbColor {
	ep := newRgbColor("GoldenRod")
	return *ep
}

// Gets the color green yellow.
func (p *RgbColor) GreenYellow() RgbColor {
	ep := newRgbColor("GreenYellow")
	return *ep
}

// Gets the color honey dew.
func (p *RgbColor) HoneyDew() RgbColor {
	ep := newRgbColor("HoneyDew")
	return *ep
}

// Gets the color hot pink.
func (p *RgbColor) HotPink() RgbColor {
	ep := newRgbColor("HotPink")
	return *ep
}

// Gets the color indian red.
func (p *RgbColor) IndianRed() RgbColor {
	ep := newRgbColor("IndianRed")
	return *ep
}

// Gets the color indigo.
func (p *RgbColor) Indigo() RgbColor {
	ep := newRgbColor("Indigo")
	return *ep
}

// Gets the color ivory.
func (p *RgbColor) Ivory() RgbColor {
	ep := newRgbColor("Ivory")
	return *ep
}

// Gets the color khaki.
func (p *RgbColor) Khaki() RgbColor {
	ep := newRgbColor("Khaki")
	return *ep
}

// Gets the color lavender.
func (p *RgbColor) Lavender() RgbColor {
	ep := newRgbColor("Lavender")
	return *ep
}

// Gets the color lavender blush.
func (p *RgbColor) LavenderBlush() RgbColor {
	ep := newRgbColor("LavenderBlush")
	return *ep
}

// Gets the color lawn Green.
func (p *RgbColor) LawnGreen() RgbColor {
	ep := newRgbColor("LawnGreen")
	return *ep
}

// Gets the color lemon chiffon.
func (p *RgbColor) LemonChiffon() RgbColor {
	ep := newRgbColor("LemonChiffon")
	return *ep
}

// Gets the color light blue.
func (p *RgbColor) LightBlue() RgbColor {
	ep := newRgbColor("LightBlue")
	return *ep
}

// Gets the color light coral.
func (p *RgbColor) LightCoral() RgbColor {
	ep := newRgbColor("LightCoral")
	return *ep
}

// Gets the color light cyan.
func (p *RgbColor) LightCyan() RgbColor {
	ep := newRgbColor("LightCyan")
	return *ep
}

// Gets the color light golden rod yellow.
func (p *RgbColor) LightGoldenRodYellow() RgbColor {
	ep := newRgbColor("LightGoldenRodYellow")
	return *ep
}

// Gets the color light grey.
func (p *RgbColor) LightGrey() RgbColor {
	ep := newRgbColor("LightGrey")
	return *ep
}

// Gets the color light green.
func (p *RgbColor) LightGreen() RgbColor {
	ep := newRgbColor("LightGreen")
	return *ep
}

// Gets the color light pink.
func (p *RgbColor) LightPink() RgbColor {
	ep := newRgbColor("LightPink")
	return *ep
}

// Gets the color light salmon.
func (p *RgbColor) LightSalmon() RgbColor {
	ep := newRgbColor("LightSalmon")
	return *ep
}

// Gets the color light sea green.
func (p *RgbColor) LightSeaGreen() RgbColor {
	ep := newRgbColor("LightSeaGreen")
	return *ep
}

// Gets the color light sky blue.
func (p *RgbColor) LightSkyBlue() RgbColor {
	ep := newRgbColor("LightSkyBlue")
	return *ep
}

// Gets the color light slate blue.
func (p *RgbColor) LightSlateBlue() RgbColor {
	ep := newRgbColor("LightSlateBlue")
	return *ep
}

// Gets the color light slate gray.
func (p *RgbColor) LightSlateGray() RgbColor {
	ep := newRgbColor("LightSlateGray")
	return *ep
}

// Gets the color light steel blue.
func (p *RgbColor) LightSteelBlue() RgbColor {
	ep := newRgbColor("LightSteelBlue")
	return *ep
}

// Gets the color light yellow.
func (p *RgbColor) LightYellow() RgbColor {
	ep := newRgbColor("LightYellow")
	return *ep
}

// Gets the color lime green.
func (p *RgbColor) LimeGreen() RgbColor {
	ep := newRgbColor("LimeGreen")
	return *ep
}

// Gets the color linen.
func (p *RgbColor) Linen() RgbColor {
	ep := newRgbColor("Linen")
	return *ep
}

// Gets the color maroon.
func (p *RgbColor) Maroon() RgbColor {
	ep := newRgbColor("Maroon")
	return *ep
}

// Gets the color medium aqua marine.
func (p *RgbColor) MediumAquaMarine() RgbColor {
	ep := newRgbColor("MediumAquaMarine")
	return *ep
}

// Gets the color medium blue.
func (p *RgbColor) MediumBlue() RgbColor {
	ep := newRgbColor("MediumBlue")
	return *ep
}

// Gets the color medium orchid.
func (p *RgbColor) MediumOrchid() RgbColor {
	ep := newRgbColor("MediumOrchid")
	return *ep
}

// Gets the color medium purple.
func (p *RgbColor) MediumPurple() RgbColor {
	ep := newRgbColor("MediumPurple")
	return *ep
}

// Gets the color medium sea green.
func (p *RgbColor) MediumSeaGreen() RgbColor {
	ep := newRgbColor("MediumSeaGreen")
	return *ep
}

// Gets the color medium slate blue.
func (p *RgbColor) MediumSlateBlue() RgbColor {
	ep := newRgbColor("MediumSlateBlue")
	return *ep
}

// Gets the color medium spring green.
func (p *RgbColor) MediumSpringGreen() RgbColor {
	ep := newRgbColor("MediumSpringGreen")
	return *ep
}

// Gets the color medium turquoise.
func (p *RgbColor) MediumTurquoise() RgbColor {
	ep := newRgbColor("MediumTurquoise")
	return *ep
}

// Gets the color medium violet red.
func (p *RgbColor) MediumVioletRed() RgbColor {
	ep := newRgbColor("MediumVioletRed")
	return *ep
}

// Gets the color midnight blue.
func (p *RgbColor) MidnightBlue() RgbColor {
	ep := newRgbColor("MidnightBlue")
	return *ep
}

// Gets the color mint cream.
func (p *RgbColor) MintCream() RgbColor {
	ep := newRgbColor("MintCream")
	return *ep
}

// Gets the color misty rose.
func (p *RgbColor) MistyRose() RgbColor {
	ep := newRgbColor("MistyRose")
	return *ep
}

// Gets the color moccasin.
func (p *RgbColor) Moccasin() RgbColor {
	ep := newRgbColor("Moccasin")
	return *ep
}

// Gets the color navajo white.
func (p *RgbColor) NavajoWhite() RgbColor {
	ep := newRgbColor("NavajoWhite")
	return *ep
}

// Gets the color navy.
func (p *RgbColor) Navy() RgbColor {
	ep := newRgbColor("Navy")
	return *ep
}

// Gets the color old lace.
func (p *RgbColor) OldLace() RgbColor {
	ep := newRgbColor("OldLace")
	return *ep
}

// Gets the color olive.
func (p *RgbColor) Olive() RgbColor {
	ep := newRgbColor("Olive")
	return *ep
}

// Gets the color olive drab.
func (p *RgbColor) OliveDrab() RgbColor {
	ep := newRgbColor("OliveDrab")
	return *ep
}

// Gets the color gainsboro.
func (p *RgbColor) Gainsboro() RgbColor {
	ep := newRgbColor("Gainsboro")
	return *ep
}

// Gets the color orange.
func (p *RgbColor) Orange() RgbColor {
	ep := newRgbColor("Orange")
	return *ep
}

// Gets the color orange red.
func (p *RgbColor) OrangeRed() RgbColor {
	ep := newRgbColor("OrangeRed")
	return *ep
}

// Gets the color orchid.
func (p *RgbColor) Orchid() RgbColor {
	ep := newRgbColor("Orchid")
	return *ep
}

// Gets the color pale golden rod.
func (p *RgbColor) PaleGoldenRod() RgbColor {
	ep := newRgbColor("PaleGoldenRod")
	return *ep
}

// Gets the color pale green.
func (p *RgbColor) PaleGreen() RgbColor {
	ep := newRgbColor("PaleGreen")
	return *ep
}

// Gets the color pale turquoise.
func (p *RgbColor) PaleTurquoise() RgbColor {
	ep := newRgbColor("PaleTurquoise")
	return *ep
}

// Gets the color pale violet red.
func (p *RgbColor) PaleVioletRed() RgbColor {
	ep := newRgbColor("PaleVioletRed")
	return *ep
}

// Gets the color papaya whip.
func (p *RgbColor) PapayaWhip() RgbColor {
	ep := newRgbColor("PapayaWhip")
	return *ep
}

// Gets the color peach puff.
func (p *RgbColor) PeachPuff() RgbColor {
	ep := newRgbColor("PeachPuff")
	return *ep
}

// Gets the color peru.
func (p *RgbColor) Peru() RgbColor {
	ep := newRgbColor("Peru")
	return *ep
}

// Gets the color pink.
func (p *RgbColor) Pink() RgbColor {
	ep := newRgbColor("Pink")
	return *ep
}

// Gets the color plum.
func (p *RgbColor) Plum() RgbColor {
	ep := newRgbColor("Plum")
	return *ep
}

// Gets the color powder blue.
func (p *RgbColor) PowderBlue() RgbColor {
	ep := newRgbColor("PowderBlue")
	return *ep
}

// Gets the color rosy brown.
func (p *RgbColor) RosyBrown() RgbColor {
	ep := newRgbColor("RosyBrown")
	return *ep
}

// Gets the color royal blue.
func (p *RgbColor) RoyalBlue() RgbColor {
	ep := newRgbColor("RoyalBlue")
	return *ep
}

// Gets the color saddle brown.
func (p *RgbColor) SaddleBrown() RgbColor {
	ep := newRgbColor("SaddleBrown")
	return *ep
}

// Gets the color salmon.
func (p *RgbColor) Salmon() RgbColor {
	ep := newRgbColor("Salmon")
	return *ep
}

// Gets the color sandy brown.
func (p *RgbColor) SandyBrown() RgbColor {
	ep := newRgbColor("SandyBrown")
	return *ep
}

// Gets the color sea green.
func (p *RgbColor) SeaGreen() RgbColor {
	ep := newRgbColor("SeaGreen")
	return *ep
}

// Gets the color sea shell.
func (p *RgbColor) SeaShell() RgbColor {
	ep := newRgbColor("SeaShell")
	return *ep
}

// Gets the color sienna.
func (p *RgbColor) Sienna() RgbColor {
	ep := newRgbColor("Sienna")
	return *ep
}

// Gets the color sky blue.
func (p *RgbColor) SkyBlue() RgbColor {
	ep := newRgbColor("SkyBlue")
	return *ep
}

// Gets the color slate blue.
func (p *RgbColor) SlateBlue() RgbColor {
	ep := newRgbColor("SlateBlue")
	return *ep
}

// Gets the color slate gray.
func (p *RgbColor) SlateGray() RgbColor {
	ep := newRgbColor("SlateGray")
	return *ep
}

// Gets the color snow.
func (p *RgbColor) Snow() RgbColor {
	ep := newRgbColor("Snow")
	return *ep
}

// Gets the color spring green.
func (p *RgbColor) SpringGreen() RgbColor {
	ep := newRgbColor("SpringGreen")
	return *ep
}

// Gets the color steel blue.
func (p *RgbColor) SteelBlue() RgbColor {
	ep := newRgbColor("SteelBlue")
	return *ep
}

// Gets the color Tan.
func (p *RgbColor) Tan() RgbColor {
	ep := newRgbColor("Tan")
	return *ep
}

// Gets the color teal.
func (p *RgbColor) Teal() RgbColor {
	ep := newRgbColor("Teal")
	return *ep
}

// Gets the color thistle.
func (p *RgbColor) Thistle() RgbColor {
	ep := newRgbColor("Thistle")
	return *ep
}

// Gets the color tomato.
func (p *RgbColor) Tomato() RgbColor {
	ep := newRgbColor("Tomato")
	return *ep
}

// Gets the color turquoise.
func (p *RgbColor) Turquoise() RgbColor {
	ep := newRgbColor("Turquoise")
	return *ep
}

// Gets the color violet.
func (p *RgbColor) Violet() RgbColor {
	ep := newRgbColor("Violet")
	return *ep
}

// Gets the color violet red.
func (p *RgbColor) VioletRed() RgbColor {
	ep := newRgbColor("VioletRed")
	return *ep
}

// Gets the color wheat.
func (p *RgbColor) Wheat() RgbColor {
	ep := newRgbColor("Wheat")
	return *ep
}

// Gets the color white smoke.
func (p *RgbColor) WhiteSmoke() RgbColor {
	ep := newRgbColor("WhiteSmoke")
	return *ep
}

// Gets the color yellow green.
func (p *RgbColor) YellowGreen() RgbColor {
	ep := newRgbColor("YellowGreen")
	return *ep
}
