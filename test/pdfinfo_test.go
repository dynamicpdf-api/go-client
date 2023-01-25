package test

import (
	"os"
	"testing"

	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/resource"
)

func TestPdfInfo(t *testing.T) {

	resource := resource.NewPdfResourceWithResourcePath("../InputResources/AllPageElements.pdf", "AllPageElements.pdf")
	text := endpoint.NewPdfInfoResource(resource)
	text.Endpoint.BaseUrl = TestArgs.BaseUrl
	text.Endpoint.ApiKey = TestArgs.ApiKey
	resp := text.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {

			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/TestPdfInfo.json")
		os.WriteFile("../test_output/TestPdfInfo.json", res.Content().Bytes(), os.ModeType)
	}
}
