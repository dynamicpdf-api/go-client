package test

import (
	"os"
	"testing"

	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/resource"
)

func TestPdfXmp(t *testing.T) {
	resource := resource.NewPdfResourceWithResourcePath("../InputResources/bab6c782-2e85-4c6a-b248-9518a06549e900000.pdf", "bab6c782-2e85-4c6a-b248-9518a06549e900000.pdf")
	xmp := endpoint.NewPdfXmp(resource)
	xmp.Endpoint.BaseUrl = TestArgs.BaseUrl
	xmp.Endpoint.ApiKey = TestArgs.ApiKey
	resp := xmp.Process()
	res := <-resp
	t.Log("IsSuccessful: ", res.IsSuccessful())
	if res.IsSuccessful() == false {
		if res.ClientError() != nil {

			t.Error("Failed with error: " + res.ClientError().Error())
		} else {
			t.Error("Failed with error: " + res.ErrorJson())
		}
	} else if TestArgs.Logging {
		os.Remove("../test_output/Templateoutput.xml")
		os.WriteFile("../test_output/Templateoutput.xml", res.Content().Bytes(), os.ModeType)
	}
}
