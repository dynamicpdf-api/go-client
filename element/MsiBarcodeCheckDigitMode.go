package element

type MsiBarcodeCheckDigitMode string

const (

	// check digit of mod 10.
	Mod10 MsiBarcodeCheckDigitMode = "mod10"

	// check digit of mod 11.
	Mod11 MsiBarcodeCheckDigitMode = "mod11"

	// check digit of mod 1010.
	Mod1010 MsiBarcodeCheckDigitMode = "mod1010"

	// check digit of mod 1110.
	Mod1110 MsiBarcodeCheckDigitMode = "mod1110"
)
