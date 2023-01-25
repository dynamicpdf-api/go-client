package test

import (
	"os"
	"testing"

	"github.com/dynamicpdf-api/go-client/input"
	"github.com/dynamicpdf-api/go-client/resource"
)

func TestHtmlInput(t *testing.T) {
	pdfCl := getNewPdf()
	resource := resource.NewHtmlResource("../InputResources/htmlSample.html", "abc")
	input1 := input.NewHtmlInputWithResource(resource)
	input1.SetBasePath("https://www.google.com/images/branding/googlelogo/1x/")
	input1.SetPageSize(input.A5)
	input1.SetPageOrientation(input.Landscape)
	pdfCl.Inputs = append(pdfCl.Inputs, input1)
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
		os.Remove("../test_output/html.pdf")
		os.WriteFile("../test_output/html.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestHtmlInputWithOptions(t *testing.T) {

	pdfCl := getNewPdf()
	input := input.NewHtmlInputWithString("<html><body>hello</body></html>")
	input.SetPageWidth(600)
	input.SetPageHeight(200)
	input.SetLeftMargin(200)

	pdfCl.Inputs = append(pdfCl.Inputs, input)
	pdfCl.Endpoint.BaseUrl = TestArgs.BaseUrl
	pdfCl.Endpoint.ApiKey = TestArgs.ApiKey
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
		os.Remove("../test_output/htmlOption.pdf")
		os.WriteFile("../test_output/htmlOption.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestHtmlInputWithBasePath(t *testing.T) {

	pdfCl := getNewPdf()
	resource := resource.NewHtmlResource("../InputResources/htmlSample.html", "abc")
	input := input.NewHtmlInputWithResource(resource)
	input.SetBasePath("https://www.google.com/images/branding/googlelogo/1x/")
	input.SetPageWidth(600)
	input.SetPageHeight(200)
	input.SetLeftMargin(200)

	pdfCl.Inputs = append(pdfCl.Inputs, input)
	pdfCl.Endpoint.BaseUrl = TestArgs.BaseUrl
	pdfCl.Endpoint.ApiKey = TestArgs.ApiKey
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
		os.Remove("../test_output/htmlOptionWithBasepath.pdf")
		os.WriteFile("../test_output/htmlOptionwithBasepath.pdf", res.Content().Bytes(), os.ModeType)
	}
}
