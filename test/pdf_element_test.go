package test

import (
	"os"
	"testing"

	"github.com/dynamicpdf-api/go-client/color"
	"github.com/dynamicpdf-api/go-client/element"
	"github.com/dynamicpdf-api/go-client/input"
	"github.com/dynamicpdf-api/go-client/resource"
)

func TestTemplate(t *testing.T) {
	pdfCl := getNewPdf()

	pageInput := input.NewPage()
	lineElement := element.NewLine(element.BottomCenter, 150, 200)
	pageInput.Elements = append(pageInput.Elements, lineElement)
	pageInput.PageHeight = 612
	pageInput.PageWidth = 1008
	pdfCl.Inputs = append(pdfCl.Inputs, pageInput)
	t.Log("Endpoint Name: ", pdfCl.EndpointName())
	t.Log("Endpoint Url : ", pdfCl.BaseUrl())
	t.Log("Endpoint Key : ", pdfCl.ApiKey())
	resp := pdfCl.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/Templateoutput.pdf")
		os.WriteFile("../test_output/Templateoutput.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestTextElement(t *testing.T) {
	pdfCl := getNewPdf()

	textElement := element.NewText("element", element.TopCenter, 50, 100)
	textElement.SetText("Element")
	pageInput := input.NewPage()
	pageInput.Elements = append(pageInput.Elements, textElement)
	pdfResource := resource.NewPdfResourceWithResourcePath("../InputResources/DocumentC.pdf", "DocumentC.pdf")
	pageInput.ResourceName = pdfResource.ResourceName
	pdfCl.Inputs = append(pdfCl.Inputs, pageInput)
	t.Log("Endpoint Name: ", pdfCl.EndpointName())
	t.Log("Endpoint Url : ", pdfCl.BaseUrl())
	t.Log("Endpoint Key : ", pdfCl.ApiKey())
	resp := pdfCl.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/TextElementoutput.pdf")
		os.WriteFile("../test_output/TextElementoutput.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestLineElement(t *testing.T) {
	pdfCl := getNewPdf()

	lineElement := element.NewLine(element.BottomRight, 30, 100)
	lineElement.SetX2Offset(200)
	lineElement.SetY2Offset(500)
	lineElement.SetWidth(50)
	lineElement.SetColor(color.NewRgbColor(255, 0, 0).Color)
	pageInput := input.NewPage()
	pageInput.PageHeight = 5000
	pageInput.PageWidth = 2000
	pageInput.Elements = append(pageInput.Elements, lineElement)
	pdfCl.Inputs = append(pdfCl.Inputs, pageInput)
	t.Log("Endpoint Name: ", pdfCl.EndpointName())
	t.Log("Endpoint Url : ", pdfCl.BaseUrl())
	t.Log("Endpoint Key : ", pdfCl.ApiKey())
	resp := pdfCl.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/LineElementoutput.pdf")
		os.WriteFile("../test_output/LineElementoutput.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestPageNumberElement(t *testing.T) {
	pdfCl := getNewPdf()

	pdfResource := resource.NewPdfResourceWithResourcePath("../InputResources/Invoice.pdf", "Invoice.pdf")
	temp1 := element.NewTemplate()
	temp1.Id = "TemplateA"
	pagenumberElement := element.NewPageNumberingElement("%%CP%% of %%TP%%", element.TopLeft, 100, 100)
	temp1.Elements = append(temp1.Elements, pagenumberElement)
	mergerOptions := input.NewMergeOptions()
	mergerOptions.DocumentInfo = true
	pdfInput := input.NewPdfWithResourceWithMergerOption(pdfResource, mergerOptions)
	pdfInput.SetTemplate(temp1)
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput)
	t.Log("Endpoint Name: ", pdfCl.EndpointName())
	t.Log("Endpoint Url : ", pdfCl.BaseUrl())
	t.Log("Endpoint Key : ", pdfCl.ApiKey())
	resp := pdfCl.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/PageNumberingElementoutput.pdf")
		os.WriteFile("../test_output/PageNumberingElementoutput.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestRectangleElement(t *testing.T) {
	pdfCl := getNewPdf()
	rectangle := element.NewRectangle(element.BottomRight, 100, 600)
	rectangle.SetCornerRadius(10)
	rectangle.SetBorderWidth(5)
	color1 := color.NewRgbColorDefault()
	rectangle.SetBorderColor(color1.AliceBlue().Color)
	pdfInput := input.NewPage()
	pdfInput.Elements = append(pdfInput.Elements, rectangle)
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput)
	t.Log("Endpoint Name: ", pdfCl.EndpointName())
	t.Log("Endpoint Url : ", pdfCl.BaseUrl())
	t.Log("Endpoint Key : ", pdfCl.ApiKey())
	resp := pdfCl.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/RectangleElementoutput.pdf")
		os.WriteFile("../test_output/RectangleElementoutput.pdf", res.Content().Bytes(), os.ModeType)
	}
}
func TestAztecBarcodeElement(t *testing.T) {
	pdfCl := getNewPdf()
	aztecBarcode := element.NewAztecBarcodeElement("Hello World", element.BottomRight, 50.0, 50.0)
	aztecBarcode.SetSymbolSize(element.R105xC105)
	aztecBarcode.SetXDimension(2)
	aztecBarcode.XOffset = 50
	aztecBarcode.YOffset = 50
	pdfInput := input.NewPage()
	pdfInput.Elements = append(pdfInput.Elements, aztecBarcode)
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput)
	t.Log("Endpoint Name: ", pdfCl.EndpointName())
	t.Log("Endpoint Url : ", pdfCl.BaseUrl())
	t.Log("Endpoint Key : ", pdfCl.ApiKey())
	resp := pdfCl.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/AztecElementoutput.pdf")
		os.WriteFile("../test_output/AztecElementoutput.pdf", res.Content().Bytes(), os.ModeType)
	}
}
func TestCode11BarcodeElement(t *testing.T) {
	pdfCl := getNewPdf()
	barcodeElement := element.NewCode11Barcode("12345678", element.TopCenter, 2.3, 100.0, 200.0)
	pdfInput := input.NewPage()
	pdfInput.Elements = append(pdfInput.Elements, barcodeElement)
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput)
	t.Log("Endpoint Name: ", pdfCl.EndpointName())
	t.Log("Endpoint Url : ", pdfCl.BaseUrl())
	t.Log("Endpoint Key : ", pdfCl.ApiKey())
	resp := pdfCl.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/Code11Elementoutput.pdf")
		os.WriteFile("../test_output/Code11Elementoutput.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestCode25BarcodeElement(t *testing.T) {
	pdfCl := getNewPdf()
	barcodeElement := element.NewCode25Barcode("12345678", element.TopCenter, 2.3, 50.0, 200.0)
	pdfInput := input.NewPage()
	pdfInput.Elements = append(pdfInput.Elements, barcodeElement)
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput)
	t.Log("Endpoint Name: ", pdfCl.EndpointName())
	t.Log("Endpoint Url : ", pdfCl.BaseUrl())
	t.Log("Endpoint Key : ", pdfCl.ApiKey())
	resp := pdfCl.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/Code25Elementoutput.pdf")
		os.WriteFile("../test_output/Code25Elementoutput.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestCode39BarcodeElement(t *testing.T) {
	pdfCl := getNewPdf()
	barcodeElement := element.NewCode39Barcode("Code 39", element.TopCenter, 2.3, 100.0, 500.0)
	pdfInput := input.NewPage()
	pdfInput.Elements = append(pdfInput.Elements, barcodeElement)
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput)
	t.Log("Endpoint Name: ", pdfCl.EndpointName())
	t.Log("Endpoint Url : ", pdfCl.BaseUrl())
	t.Log("Endpoint Key : ", pdfCl.ApiKey())
	resp := pdfCl.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/Code39Elementoutput.pdf")
		os.WriteFile("../test_output/Code39Elementoutput.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestCode93BarcodeElement(t *testing.T) {
	pdfCl := getNewPdf()
	barcodeElement := element.NewCode93Barcode("CODE 93", element.TopCenter, 2.3, 10.0, 200.0)
	pdfInput := input.NewPage()
	pdfInput.Elements = append(pdfInput.Elements, barcodeElement)
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput)
	t.Log("Endpoint Name: ", pdfCl.EndpointName())
	t.Log("Endpoint Url : ", pdfCl.BaseUrl())
	t.Log("Endpoint Key : ", pdfCl.ApiKey())
	resp := pdfCl.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/Code93Elementoutput.pdf")
		os.WriteFile("../test_output/Code93Elementoutput.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestCode128BarcodeElement(t *testing.T) {
	pdfCl := getNewPdf()
	barcodeElement := element.NewCode128Barcode("CODE 128", element.TopCenter, 2.3, 100.0, 50.0)
	pdfInput := input.NewPage()
	pdfInput.Elements = append(pdfInput.Elements, barcodeElement)
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput)
	t.Log("Endpoint Name: ", pdfCl.EndpointName())
	t.Log("Endpoint Url : ", pdfCl.BaseUrl())
	t.Log("Endpoint Key : ", pdfCl.ApiKey())
	resp := pdfCl.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/Code128Elementoutput.pdf")
		os.WriteFile("../test_output/Code128Elementoutput.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestPdf417BarcodeElement(t *testing.T) {
	pdfCl := getNewPdf()
	barcodeElement := element.NewPdf417Barcode("Pdf417", element.TopCenter, 2, 100.0, 200.0)
	pdfInput := input.NewPage()
	pdfInput.Elements = append(pdfInput.Elements, barcodeElement)
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput)
	t.Log("Endpoint Name: ", pdfCl.EndpointName())
	t.Log("Endpoint Url : ", pdfCl.BaseUrl())
	t.Log("Endpoint Key : ", pdfCl.ApiKey())
	resp := pdfCl.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/Pdf417Elementoutput.pdf")
		os.WriteFile("../test_output/Pdf417Elementoutput.pdf", res.Content().Bytes(), os.ModeType)
	}
}
func TestMsiBarcodeElement(t *testing.T) {
	pdfCl := getNewPdf()
	barcodeElement := element.NewMsiBarcode("12345678", element.TopCenter, 10, 100.0, 200.0)
	pdfInput := input.NewPage()
	pdfInput.Elements = append(pdfInput.Elements, barcodeElement)
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput)
	t.Log("Endpoint Name: ", pdfCl.EndpointName())
	t.Log("Endpoint Url : ", pdfCl.BaseUrl())
	t.Log("Endpoint Key : ", pdfCl.ApiKey())
	resp := pdfCl.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/MsiElementoutput.pdf")
		os.WriteFile("../test_output/MsiElementoutput.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestQrCodeElement(t *testing.T) {
	pdfCl := getNewPdf()
	barcodeElement := element.NewQrCodeElement("12345678", element.TopLeft, 300.0, 200.0)
	pdfInput := input.NewPage()
	pdfInput.Elements = append(pdfInput.Elements, barcodeElement)
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput)
	t.Log("Endpoint Name: ", pdfCl.EndpointName())
	t.Log("Endpoint Url : ", pdfCl.BaseUrl())
	t.Log("Endpoint Key : ", pdfCl.ApiKey())
	resp := pdfCl.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/QrCodeElementoutput.pdf")
		os.WriteFile("../test_output/QrCodeElementoutput.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestGs1DataElement(t *testing.T) {
	pdfCl := getNewPdf()
	barcodeElement := element.NewGs1DataBarBarcode("12345678", element.TopLeft, 20, element.Expanded, 200.0, 200.0)
	pdfInput := input.NewPage()
	pdfInput.Elements = append(pdfInput.Elements, barcodeElement)
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput)
	t.Log("Endpoint Name: ", pdfCl.EndpointName())
	t.Log("Endpoint Url : ", pdfCl.BaseUrl())
	t.Log("Endpoint Key : ", pdfCl.ApiKey())
	resp := pdfCl.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/Gs1Dataoutput.pdf")
		os.WriteFile("../test_output/Gs1Dataoutput.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestIata25Element(t *testing.T) {
	pdfCl := getNewPdf()
	barcodeElement := element.NewIata25Barcode("12345678", element.BottomLeft, 20, 50.0, 200.0)
	pdfInput := input.NewPage()
	pdfInput.Elements = append(pdfInput.Elements, barcodeElement)
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput)
	t.Log("Endpoint Name: ", pdfCl.EndpointName())
	t.Log("Endpoint Url : ", pdfCl.BaseUrl())
	t.Log("Endpoint Key : ", pdfCl.ApiKey())
	resp := pdfCl.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/Iata25output.pdf")
		os.WriteFile("../test_output/Iata25output.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestStackedGs1DataElement(t *testing.T) {
	pdfCl := getNewPdf()
	barcodeElement := element.NewStackedGs1DataBarBarcode("12345678", element.TopLeft, element.StackedOmnidirectional, 20, 200.0, 200.0)
	pdfInput := input.NewPage()
	pdfInput.Elements = append(pdfInput.Elements, barcodeElement)
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput)
	t.Log("Endpoint Name: ", pdfCl.EndpointName())
	t.Log("Endpoint Url : ", pdfCl.BaseUrl())
	t.Log("Endpoint Key : ", pdfCl.ApiKey())
	resp := pdfCl.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {
			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/StackedGs1Dataoutput.pdf")
		os.WriteFile("../test_output/StackedGs1Dataoutput.pdf", res.Content().Bytes(), os.ModeType)
	}
}
