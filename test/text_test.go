package test

import (
	"os"
	"testing"

	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/resource"
)

func TestTextExtraction(t *testing.T) {
	resource := resource.NewPdfResourceWithResourcePath("../InputResources/AllPageElements.pdf", "AllPageElements.pdf")

	text := endpoint.NewPdfText(resource, 1, 3)
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
		os.Remove("../test_output/Templateoutput.json")
		os.WriteFile("../test_output/Templateoutput.json", res.Content().Bytes(), os.ModeType)
	}
}
