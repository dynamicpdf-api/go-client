package element

type ElementType string

const (
	/*Text element.*/
	TextElement ElementType = "Text"

	/*Image element.*/
	ImageElement ElementType = "Image"

	/*Code 128 barcode element.*/
	Code128BarcodeElement ElementType = "Code128Barcode"

	/*Code 39 barcode element.*/
	Code39BarcodeElement ElementType = "Code39Barcode"

	/*Code 2 of 5 barcode element.*/
	Code25BarcodeElement ElementType = "Code25Barcode"

	/*Code 93 barcode element.*/
	Code93BarcodeElement ElementType = "Code93Barcode"

	/*Code 11 barcode element.*/
	Code11BarcodeElement ElementType = "Code11Barcode"

	/*GS1 data bar barcode element.*/
	Gs1DataBarBarcodeElement ElementType = "Gs1DataBarBarcode"

	/*IATA 25 barcode element.*/
	Iata25BarcodeElement ElementType = "Iata25Barcode"

	/*MSI barcode element.*/
	MsiBarcodeElement ElementType = "MsiBarcode"

	/*Stacked GS1 data bar barcode element.*/
	StackedGs1DataBarBarcodeElement ElementType = "StackedGs1DataBarBarcode"

	/*Aztec barcode element.*/
	AztecBarcodeElement ElementType = "AztecBarcode"

	/*Data Matrix barcode element.*/
	DataMatrixBarcodeElement ElementType = "DataMatrixBarcode"

	/*Pdf417 barcode*/
	Pdf417BarcodeElement ElementType = "Pdf417Barcode"

	/*Qrcode barcode*/
	QrCodeElement ElementType = "QrCode"

	/*PageNumbering Element*/
	PageNumberingElement ElementType = "PageNumbering"

	/*Line Element*/
	LineElement ElementType = "Line"

	/*Rectangle Element*/
	RectangleElement ElementType = "Rectangle"
)
