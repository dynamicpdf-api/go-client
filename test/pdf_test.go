package test

import (
	"os"
	"testing"

	"github.com/dynamicpdf-api/go-client/color"
	"github.com/dynamicpdf-api/go-client/element"
	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/font"
	"github.com/dynamicpdf-api/go-client/input"
	"github.com/dynamicpdf-api/go-client/resource"
)

func getNewPdf() *endpoint.Pdf {
	pdfCl := endpoint.NewPdf()
	pdfCl.Endpoint.BaseUrl = TestArgs.BaseUrl
	pdfCl.Endpoint.ApiKey = TestArgs.ApiKey

	pdfCl.Author = "Jane Doe"
	pdfCl.Title = "Sample PDF"
	pdfCl.Subject = "topLevel document metadata"
	pdfCl.Creator = "John Creator"
	pdfCl.Keywords = "dynamicpdf api example pdf java instructions"

	return pdfCl
}

func TestPage_PdfInput(t *testing.T) {
	pdfCl := getNewPdf()

	pageInput := input.NewPage()
	pageInput.PageHeight = 612
	pageInput.PageWidth = 1008
	text := element.NewText("text", element.BottomCenter, 250, 0)
	pageInput.Elements = append(pageInput.Elements, text)
	pdfResource := resource.NewPdfResourceWithResourcePath("../InputResources/AllPageElements.pdf", "AllPageElements.pdf")
	pdfResource.ResourceName = "AllPageElements.pdf"
	pdfInput := input.NewPdfWithResource(pdfResource)
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput)
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
		os.Remove("../test_output/PageInput.pdf")
		os.WriteFile("../test_output/PageInput.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestMergeOptions(t *testing.T) {
	pdfCl := getNewPdf()
	pdfResource := resource.NewPdfResourceWithResourcePath("../InputResources/DocumentC.pdf", "DocumentC.pdf")
	pdfInput := input.NewPdfWithResource(pdfResource)
	mergeOption := input.NewMergeOptions()
	mergeOption.DocumentInfo = true
	pdfInput.MergeOption = mergeOption
	pdfInput.ResourceName = pdfResource.ResourceName
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput)

	pdfResource1 := resource.NewPdfResourceWithResourcePath("../InputResources/SinglePage.pdf", "SinglePage.pdf")
	pdfInput1 := input.NewPdfWithResource(pdfResource1)
	mergeOption1 := input.NewMergeOptions()
	mergeOption1.DocumentProperties = true
	pdfInput1.MergeOption = mergeOption1
	pdfInput1.ResourceName = pdfResource1.ResourceName
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput1)

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
		os.Remove("../test_output/MergeOptionsoutput.pdf")
		os.WriteFile("../test_output/MergeOptionsoutput.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestMergeOptions1(t *testing.T) {
	pdfCl := getNewPdf()
	pdfResource := resource.NewPdfResourceWithResourcePath("../InputResources/SinglePage.pdf", "SinglePage.pdf")
	pdfInput := input.NewPdfWithResource(pdfResource)
	mergeOption := input.NewMergeOptions()
	mergeOption.DocumentInfo = true
	pdfInput.MergeOption = mergeOption
	pdfInput.ResourceName = pdfResource.ResourceName
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput)

	pdfResource1 := resource.NewPdfResourceWithResourcePath("../InputResources/AllPageElements.pdf", "AllPageElements.pdf")
	pdfInput1 := input.NewPdfWithResource(pdfResource1)
	pdfInput1.StartPage = 1
	pdfInput1.PageCount = 1
	mergeOption1 := input.NewMergeOptions()
	mergeOption1.DocumentProperties = true
	pdfInput1.MergeOption = mergeOption1
	pdfInput1.ResourceName = pdfResource1.ResourceName
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput1)
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
		os.Remove("../test_output/MergeOptionsoutput1.pdf")
		os.WriteFile("../test_output/MergeOptionsoutput1.pdf", res.Content().Bytes(), os.ModeType)
	}
}
func TestPageInput(t *testing.T) {
	pdfCl := getNewPdf()
	textElement := element.NewText("Element", element.BottomCenter, 20, 10)
	pageInput := input.NewPage()
	pageInput.PageHeight = 612
	pageInput.PageWidth = 1008
	pageInput.Elements = append(pageInput.Elements, textElement)
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
		os.Remove("../test_output/Page_PdfInputoutput.pdf")
		os.WriteFile("../test_output/Page_PdfInputoutput.pdf", res.Content().Bytes(), os.ModeType)
	}
}
func TestFormFill(t *testing.T) {
	pdfCl := getNewPdf()
	pdfResource := resource.NewPdfResourceWithResourcePath("../InputResources/fw9AcroForm_14.pdf", "fw9AcroForm_14.pdf")
	pdfInput := input.NewPdfWithResource(pdfResource)
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput)
	field1 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].f1_1[0]", "Any Company, Inc.")
	pdfCl.FormFields = append(pdfCl.FormFields, *field1)
	field2 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].f1_2[0]", "Any Company")
	pdfCl.FormFields = append(pdfCl.FormFields, *field2)
	field3 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].FederalClassification[0].c1_1[0]", "1")
	pdfCl.FormFields = append(pdfCl.FormFields, *field3)
	field4 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].Address[0].f1_7[0]", "123 Main Street")
	pdfCl.FormFields = append(pdfCl.FormFields, *field4)
	field5 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].Address[0].f1_8[0]", "Washington, DC  22222")
	pdfCl.FormFields = append(pdfCl.FormFields, *field5)
	field6 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].f1_9[0]", "Any Requester")
	field6.Flatten = true
	pdfCl.FormFields = append(pdfCl.FormFields, *field6)
	field7 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].f1_10[0]", "17288825617")
	field7.Remove = true
	pdfCl.FormFields = append(pdfCl.FormFields, *field7)
	field8 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].EmployerID[0].f1_14[0]", "52")
	pdfCl.FormFields = append(pdfCl.FormFields, *field8)
	field9 := endpoint.NewFormFieldWithValue("topmostSubform[0].Page1[0].EmployerID[0].f1_15[0]", "1234567")
	pdfCl.FormFields = append(pdfCl.FormFields, *field9)
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
		os.Remove("../test_output/formFiledoutput.pdf")
		os.WriteFile("../test_output/formFiledoutput.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestPageInput_CoreFont(t *testing.T) {
	pdfCl := getNewPdf()

	pageInput := input.NewPage()
	pageInput.PageHeight = 612
	pageInput.PageWidth = 1008
	text := element.NewText("text", element.BottomCenter, 250, 0)
	font1 := font.NewFontStyle()
	text.SetFont(font1.TimesItalic())
	pageInput.Elements = append(pageInput.Elements, text)

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
		os.Remove("../test_output/Font.pdf")
		os.WriteFile("../test_output/Font.pdf", res.Content().Bytes(), os.ModeType)
	}
}
func TestPageCMYKColor(t *testing.T) {
	pdfCl := getNewPdf()
	pageInput := input.NewPage()
	pageInput.PageHeight = 612
	pageInput.PageWidth = 1008
	text := element.NewText("text", element.BottomCenter, 250, 0)
	Color := color.NewCmykColor(1, 0, 0, 0)
	text.SetColor(Color.Color)
	pageInput.Elements = append(pageInput.Elements, text)

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
		os.Remove("../test_output/Font.pdf")
		os.WriteFile("../test_output/Font.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestPageGrayColor(t *testing.T) {
	pdfCl := getNewPdf()

	pageInput := input.NewPage()
	pageInput.PageHeight = 612
	pageInput.PageWidth = 1008
	text := element.NewText("text", element.BottomCenter, 0, 0)
	text.SetColor(color.NewGrayscalecolor(0.8).Color)
	pageInput.Elements = append(pageInput.Elements, text)

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
		os.Remove("../test_output/ColorGrayScale.pdf")
		os.WriteFile("../test_output/ColorGrayScale.pdf", res.Content().Bytes(), os.ModeType)
	}
}
func TestPageGif(t *testing.T) {
	pdfCl := getNewPdf()
	pdfResource := resource.NewPdfResourceWithResourcePath("../InputResources/SinglePage.pdf", "SinglePage.pdf")
	pdfInput := input.NewPdfWithResource(pdfResource)
	temp1 := element.NewTemplate()
	temp1.Id = "TemplateA"
	image := resource.NewImageResourceWithResourcePath("../InputResources/Northwind Logo.gif", "Northwind Logo.gif")
	imageEle := element.NewImagewithImageResource(image, element.BottomCenter, 0, 0)
	temp1.Elements = append(temp1.Elements, imageEle)
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
		os.Remove("../test_output/PageGif.pdf")
		os.WriteFile("../test_output/PageGif.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestDlex(t *testing.T) {
	pdfCl := getNewPdf()
	temp1 := element.NewTemplate()
	temp1.Id = "TemplateA"
	dlex := resource.NewDlexResourceWithPath("../InputResources/SimpleReportWithCoverPage1.dlex", "SimpleReportWithCoverPage1.dlex")
	layout := resource.NewLayoutDataResource("../InputResources/SimpleReportData.json", "SimpleReportData.json")
	dlexInput := input.NewDlexWithDlexNLayoutResources(*dlex, layout)
	pdfCl.Inputs = append(pdfCl.Inputs, dlexInput)
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
		os.Remove("../test_output/Dlex.pdf")
		os.WriteFile("../test_output/Dlex.pdf", res.Content().Bytes(), os.ModeType)
	}
}
func TestOutline(t *testing.T) {
	pdfCl := getNewPdf()
	pageInput := pdfCl.AddPage()
	text := element.NewText("text", element.BottomCenter, 250, 0)
	pageInput.Elements = append(pageInput.Elements, text)
	pdfResource := resource.NewPdfResourceWithResourcePath("../InputResources/AllPageElements.pdf", "AllPageElements.pdf")
	pdfInput := input.NewPdfWithResource(pdfResource)
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput)
	pdfResource1 := resource.NewPdfResourceWithResourcePath("../InputResources/PdfOutlineInput.pdf", "PdfOutlineInput.pdf")
	pdfInput1 := input.NewPdfWithResource(pdfResource1)
	pdfCl.Inputs = append(pdfCl.Inputs, pdfInput1)
	outline := pdfCl.Outlines.Add("Root Outline")
	outline.Expanded = true
	out := outline.Children()
	out.AddPdfOutlines(*pdfInput)
	out.AddPdfOutlines(*pdfInput1)
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
		os.Remove("../test_output/Outline.pdf")
		os.WriteFile("../test_output/Outline.pdf", res.Content().Bytes(), os.ModeType)
	}
}
