package test

import (
	"os"
	"strconv"
	"testing"

	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/resource"
)

func TestImageInfoEndpoint(t *testing.T) {

	images := [5]string{"../InputResources/3_rescale_indexed.bmp", "../InputResources/Northwind Logo.gif", "../InputResources/fw9_13.tif", "../InputResources/DPDFLogo.png", "../InputResources/image.jpg"}
	for i, input1 := range images {
		resource := resource.NewImageResourceWithResourcePath(input1, "")

		image := endpoint.NewImageInfo(resource)
		image.Endpoint.BaseUrl = TestArgs.BaseUrl
		image.Endpoint.ApiKey = TestArgs.ApiKey

		resp := image.Process()
		res := <-resp
		t.Log("IsSuccessful: ", res.IsSuccessful())
		if res.IsSuccessful() == false {
			if res.ClientError() != nil {
				t.Error("Failed with error: " + res.ClientError().Error())
			} else {
				t.Error("Failed with error: " + res.ErrorJson())
			}
			break
		} else if TestArgs.Logging {
			os.Remove("../test_output/Templateoutput" + strconv.Itoa(i) + ".json")
			os.WriteFile("../test_output/Templateoutput"+strconv.Itoa(i)+".json", res.Content().Bytes(), os.ModeType)
		}
	}
}
