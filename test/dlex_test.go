package test

import (
	"os"
	"testing"

	"github.com/dynamicpdf-api/go-client/endpoint"
	"github.com/dynamicpdf-api/go-client/resource"
)

func TestDlexInput(t *testing.T) {

	layoutDataResource := resource.NewLayoutDataResource("../InputResources/ContactList.json", "ContactList.json")
	var Filepath = "ContactList.dlex"
	layoutData := endpoint.NewDlexEndpoint(Filepath, layoutDataResource)
	layoutData.Endpoint.BaseUrl = TestArgs.BaseUrl
	layoutData.Endpoint.ApiKey = TestArgs.ApiKey
	t.Log("Endpoint Name: ", layoutData.EndpointName())
	t.Log("Endpoint Url : ", layoutData.BaseUrl())
	t.Log("Endpoint Key : ", layoutData.ApiKey())
	resp := layoutData.Process()
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
