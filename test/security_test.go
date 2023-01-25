package test

import (
	"os"
	"testing"

	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/input"
	"github.com/dynamicpdf-api/go-client/resource"
)

func TestAes128Security(t *testing.T) {
	pdfcl := getNewPdf()
	resource := resource.NewPdfResourceWithResourcePath("../InputResources/XmpAndOtherSample.pdf", "XmpAndOtherSample.pdf")
	input := input.NewPdfWithResource(resource)

	security := endpoint.NewAes128Security("user", "owner")
	pdfcl.Security = endpoint.Security(security.Security)
	pdfcl.Inputs = append(pdfcl.Inputs, input)
	resp := pdfcl.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {

			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/Aes128Security.pdf")
		os.WriteFile("../test_output/Aes128Security.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestAes128SecurityEncryptDocumentComponents(t *testing.T) {
	pdfcl := getNewPdf()
	pdfcl.Endpoint.BaseUrl = TestArgs.BaseUrl
	pdfcl.Endpoint.ApiKey = TestArgs.ApiKey
	resource := resource.NewPdfResourceWithResourcePath("../InputResources/XmpAndOtherSample.pdf", "XmpAndOtherSample.pdf")
	input := input.NewPdfWithResource(resource)
	security := endpoint.NewAes128Security("user", "owner")
	security.SetDocumentComponents(endpoint.All) //EncryptDocumentComponents.all
	pdfcl.Security = endpoint.Security(security.Security)
	pdfcl.Inputs = append(pdfcl.Inputs, input)
	resp := pdfcl.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {

			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/Aes128Security.pdf")
		os.WriteFile("../test_output/Aes128Security.pdf", res.Content().Bytes(), os.ModeType)
	}
}

func TestRC4128Security(t *testing.T) {
	pdfcl := getNewPdf()
	resource := resource.NewPdfResourceWithResourcePath("../InputResources/XmpAndOtherSample.pdf", "XmpAndOtherSample.pdf")
	input := input.NewPdfWithResource(resource)
	security := endpoint.NewRC4128Security()
	security.UserPassword = "user"
	security.OwnerPassword = "owner"
	security.SetEncryptMetadata(true)
	pdfcl.Security = security.Security
	pdfcl.Inputs = append(pdfcl.Inputs, input)
	resp := pdfcl.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {

			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/RC4128Security.pdf")
		os.WriteFile("../test_output/RC4128Security.pdf", res.Content().Bytes(), os.ModeType)
	}
}
