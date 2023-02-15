package element

type DataMatrixEncodingType string

const (

	// Calculates Encoding based on Value.
	AutoDataMatrixEncodingType DataMatrixEncodingType = "Auto"

	// Calculates ASCII Encoding based on Value.
	AutoAscii DataMatrixEncodingType = "AutoAscii"

	// ASCII Encoding.
	Ascii DataMatrixEncodingType = "Ascii"

	// Extended ASCII Encoding.
	ExtendedAscii DataMatrixEncodingType = "ExtendedAscii"

	// Double digit Encoding.
	DoubleDigitDataMatrixEncodingType = "DoubleDigit"

	// C40 Encoding.
	C40 DataMatrixEncodingType = "C40"

	// Text Encoding.
	TextDataMatrixEncodingType DataMatrixEncodingType = "Text"

	// ANSI X12 Encoding.
	AnsiX12 DataMatrixEncodingType = "AnsiX12"

	// EDIFACT Encoding.
	Edifact DataMatrixEncodingType = "Edifact"

	// Base256 Encoding.
	Base256 DataMatrixEncodingType = "Base256"
)
