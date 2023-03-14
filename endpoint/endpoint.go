package endpoint

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	// Default ApiKey for all endpoints.
	DefaultApiKey string = ""

	// Default base URL
	DefaultBaseUrl string = "https://api.dynamicpdf.com/"

	pkgLog log.Logger = *log.New(io.Discard, "PKG_LOG: ", 0)
)

/*
Represents the base class for endpoint and has settings for base url,
api key and creates a rest request object.
*/
type EndpointProcessor interface {
	EndpointName() string
	BaseUrl() string
	ApiKey() string
}

type Endpoint struct {
	EndpointProcessor
	BaseUrl, ApiKey string
}

type formData struct {
	content     *bytes.Buffer
	contentType string
}

func SetLogger(logger log.Logger) {
	pkgLog = logger
}

func postForm(form formData, endPoint EndpointProcessor) (response, error) {
	var response response = response{isSuccessful: false}

	postUrl := strings.TrimSuffix(endPoint.BaseUrl(), "/") + "/v1.0/" + endPoint.EndpointName()
	postAuth := "Bearer " + endPoint.ApiKey()

	var httpClient http.Client
	ctx := context.TODO() // TODO: Need to add a context property for user to be able to cancel/timeout the request.
	req, err := http.NewRequestWithContext(ctx, "POST", postUrl, form.content)
	if err != nil {
		pkgLog.Printf("%s\n", err)
		return response, err
	}

	req.Header.Add("Authorization", postAuth)
	req.Header.Add("Content-Type", form.contentType)
	resp, err := httpClient.Do(req)
	if err != nil || resp.Status != "200 OK" {
		pkgLog.Printf("\nError: %s\n", err)
		response.statusCode = resp.StatusCode
		response.errorMessage = resp.Status
		body, _ := io.ReadAll(resp.Body)
		response.errorJson = string(body)
		response.clientError = err
		return response, err
	}
	pkgLog.Println("response Status:", resp.Status)
	headers := ""
	for k, v := range resp.Header {
		headers += fmt.Sprintln("\t", k, "\t:", v)
	}
	pkgLog.Println("response Headers:\n", headers)

	body, e := ioutil.ReadAll(resp.Body)
	if e == nil && err == nil {
		response.clientError = err
		response.isSuccessful = true
		response.content = bytes.NewBuffer(body)
	}
	return response, err
}

func postHttpRequest(url string, data []byte, api string, contentType string) (JsonResponse, error) {
	var response JsonResponse = JsonResponse{isSuccessful: false}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Add("Authorization", api)
	req.Header.Add("Content-Type", contentType)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	httpResponse, httpErr := client.Do(req)
	if httpErr != nil || httpResponse.StatusCode != 200 {
		pkgLog.Printf("\nError: %sn", httpErr)
		response.statusCode = httpResponse.StatusCode
		response.errorJson = httpResponse.Status
		response.clientError = httpErr
		return response, httpErr
	}
	pkgLog.Println("response Status:", httpResponse.Status)
	pkgLog.Println("response Headers:", httpResponse.Header)
	body, readErr := ioutil.ReadAll(httpResponse.Body)
	if readErr == nil {
		response.clientError = readErr
		response.content = bytes.NewBuffer(body)
		response.errorJson = httpResponse.Status
	}
	response.isSuccessful = (httpResponse.StatusCode == 200)
	return response, readErr
}
