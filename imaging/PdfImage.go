package imaging

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/dynamicpdf-api/go-client/v2/endpoint"
	"github.com/dynamicpdf-api/go-client/v2/resource"
	"github.com/google/uuid"
)

var (
	pkgLog log.Logger = *log.New(io.Discard, "PKG_LOG: ", 0)
)

// PdfImage represents a PDF Rasterizing endpoint that converts PDF to image.
type PdfImage struct {
	endpoint.Endpoint

	resource resource.PdfResource

	// Gets or sets the starting page number for rasterization.
	StartPageNumber int `json:"startPageNumber,omitempty"`
	// Gets or sets the number of pages to rasterize.
	PageCount int `json:"pageCount,omitempty"`

	// Gets or sets the image format for rasterization.
	ImageFormat ImageFormat `json:"imageFormat,omitempty"`

	// Gets or sets the size of the rasterized images.
	ImageSize ImageSize `json:"imageSize,omitempty"`
}

var _ endpoint.EndpointProcessor = (*PdfImage)(nil)

func (p *PdfImage) EndpointName() string {
	return "pdf-image"
}

func (p *PdfImage) BaseUrl() string {
	if p.Endpoint.BaseUrl != "" {
		return p.Endpoint.BaseUrl
	} else {
		return endpoint.DefaultBaseUrl
	}
}

func (p *PdfImage) ApiKey() string {
	return p.Endpoint.ApiKey
}

// Initializes a new instance of the PdfImage class with the specified PDF resource.
func NewPdfImage(resource resource.PdfResource) *PdfImage {
	var ep PdfImage
	ep.resource = resource
	return &ep
}

func (p *PdfImage) Process() <-chan PdfImageResponse {
	rasterizerResponse := make(chan PdfImageResponse)
	go func() {
		var form formData
		form.content = bytes.NewBuffer(nil)
		formWriter := multipart.NewWriter(form.content)
		form.contentType = formWriter.FormDataContentType()

		part, err := formWriter.CreateFormFile("pdf", p.resource.ResourceName)
		if err != nil {
			return
		}
		resourceData := bytes.NewBuffer(p.resource.Data())
		io.Copy(part, resourceData)

		err = formWriter.Close()
		if err != nil {
			return
		}

		postUrl := strings.TrimSuffix(p.BaseUrl(), "/") + "/v1.0/" + p.EndpointName()

		u, err := url.Parse(postUrl)
		if err != nil {
			panic(err)
		}

		// Define query parameters
		url := u.Query()

		if p.StartPageNumber > 0 {
			url.Set("sp", strconv.Itoa(p.StartPageNumber))
		}
		if p.StartPageNumber > 0 {
			url.Set("pc", strconv.Itoa(p.PageCount))
		}

		typ := p.ImageSize.Type
		imageSize := p.ImageSize
		switch typ {
		case "Dpi":
			url.Set("is", string(imageSize.Type))
			if imageSize.horizontalDpi > 0 {
				url.Set("hd", strconv.Itoa(imageSize.horizontalDpi))
			}
			if imageSize.verticalDpi > 0 {
				url.Set("vd", strconv.Itoa(imageSize.verticalDpi))
			}
		case "Fixed":
			url.Set("is", string(imageSize.Type))
			if imageSize.height > 0 {
				url.Set("ht", strconv.Itoa(imageSize.height))
			}
			if imageSize.width > 0 {
				url.Set("wd", strconv.Itoa(imageSize.width))
			}
			if imageSize.unit != "" {
				url.Set("ut", string(imageSize.unit))
			}
		case "Max":
			url.Set("is", string(imageSize.Type))
			if imageSize.maxHeight > 0 {
				url.Set("mh", strconv.Itoa(imageSize.maxHeight))
			}
			if imageSize.maxWidth > 0 {
				url.Set("mw", strconv.Itoa(imageSize.maxWidth))
			}
			if imageSize.unit != "" {
				url.Set("ut", string(imageSize.unit))
			}
		case "Percentage":
			url.Set("is", string(imageSize.Type))
			if imageSize.horizontalPercentage > 0 {
				url.Set("hp", strconv.Itoa(imageSize.horizontalPercentage))
			}
			if imageSize.verticalPercentage > 0 {
				url.Set("vp", strconv.Itoa(imageSize.verticalPercentage))
			}
		}

		format := p.ImageFormat.Type
		imageFormat := p.ImageFormat

		switch format {
		case "Gif":
			url.Set("if", string(imageFormat.Type))
			if imageFormat.ditheringPercent > 0 {
				url.Set("dp", strconv.Itoa(imageFormat.ditheringPercent))
			}
			if imageFormat.ditheringAlgorithm != "" {
				url.Set("da", string(imageFormat.ditheringAlgorithm))
			}
		case "Jpeg":
			url.Set("if", string(imageFormat.Type))
			if imageFormat.quality > 0 {
				url.Set("qt", strconv.Itoa(imageFormat.quality))
			}
		case "Png":
			url.Set("if", string(imageFormat.Type))

			cfmt := p.ImageFormat.colorFormat
			fmt := cfmt.Type
			switch fmt {
			case "Indexed":
				url.Set("cf", string(cfmt.Type))
				if cfmt.ditheringAlgorithm != "" {
					url.Set("da", string(cfmt.ditheringAlgorithm))
				}
				if cfmt.ditheringPercent > 0 {
					url.Set("dp", strconv.Itoa(cfmt.ditheringPercent))
				}
				if cfmt.quantizationAlgorithm != "" {
					url.Set("qa", string(cfmt.quantizationAlgorithm))
				}
			case "Monochrome":
				url.Set("cf", string(cfmt.Type))
				if cfmt.blackThreshold > 0 {
					url.Set("bt", strconv.Itoa(cfmt.blackThreshold))
				}
				if cfmt.ditheringAlgorithm != "" {
					url.Set("da", string(cfmt.ditheringAlgorithm))
				}
				if cfmt.ditheringPercent > 0 {
					url.Set("dp", strconv.Itoa(cfmt.ditheringPercent))
				}
			}
		case "Tiff":
			url.Set("if", string(imageFormat.Type))
			if p.ImageFormat.multiPage {
				url.Set("mp", "true")
			}
			cfmt := p.ImageFormat.colorFormat
			fmt := cfmt.Type
			switch fmt {
			case "Indexed":
				url.Set("cf", string(cfmt.Type))
				if cfmt.ditheringAlgorithm != "" {
					url.Set("da", string(cfmt.ditheringAlgorithm))
				}
				if cfmt.ditheringPercent > 0 {
					url.Set("dp", strconv.Itoa(cfmt.ditheringPercent))
				}
				if cfmt.quantizationAlgorithm != "" {
					url.Set("qa", string(cfmt.quantizationAlgorithm))
				}
			case "Monochrome":
				url.Set("cf", string(cfmt.Type))
				if cfmt.compressionType != "" {
					url.Set("ct", string(cfmt.compressionType))
				}
				if cfmt.blackThreshold > 0 {
					url.Set("bt", strconv.Itoa(cfmt.blackThreshold))
				}
				if cfmt.ditheringAlgorithm != "" {
					url.Set("da", string(cfmt.ditheringAlgorithm))
				}
				if cfmt.ditheringPercent > 0 {
					url.Set("dp", strconv.Itoa(cfmt.ditheringPercent))
				}
			}
		case "Bmp":
			url.Set("if", string(imageFormat.Type))
			Bcfmt := p.ImageFormat.colorFormat
			fmt := Bcfmt.Type
			switch fmt {
			case "Monochrome":
				url.Set("cf", string(Bcfmt.Type))
				if Bcfmt.blackThreshold > 0 {
					url.Set("bt", strconv.Itoa(Bcfmt.blackThreshold))
				}
				if Bcfmt.ditheringPercent > 0 {
					url.Set("dp", strconv.Itoa(Bcfmt.ditheringPercent))
				}
				if Bcfmt.ditheringAlgorithm != "" {
					url.Set("da", string(Bcfmt.ditheringAlgorithm))
				}
			}
		}

		u.RawQuery = url.Encode()

		postAuth := "Bearer " + p.Endpoint.ApiKey

		res, err := postFormRast(form, postAuth, u.String())
		if err != nil {
			res.clientError = err
		}
		rastResponse := PdfImageResponse{response: res}

		if res.statusCode == 200 {

			bodyBytes, _ := io.ReadAll(res.content)

			decoder := json.NewDecoder(bytes.NewReader([]byte(bodyBytes)))
			var pdfImgResp PdfImageResponse
			err = decoder.Decode(&pdfImgResp)
			if err != nil {
				panic(err)
			}

			imageType := pdfImgResp.ContentType
			rastResponse.ContentType = imageType
			rastResponse.HorizontalDpi = pdfImgResp.HorizontalDpi
			rastResponse.VerticalDpi = pdfImgResp.VerticalDpi
			rastResponse.Images = pdfImgResp.Images
			rastResponse.ImageFormat = strings.Split(imageType, "/")[1]

			rastResponse.isSuccessful = true
			rastResponse.statusCode = res.statusCode
		}

		rasterizerResponse <- rastResponse
	}()

	return rasterizerResponse
}

func postFormRast(form formData, postAuth string, postUrl string) (response, error) {
	var response response = response{isSuccessful: false}

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
		body, _ := io.ReadAll(resp.Body)
		response.errorJson = string(body)
		var errorId string
		var errorJson map[string]string
		json.Unmarshal([]byte(response.errorJson), &errorJson)
		errorId = errorJson["id"]
		if errorId != "" {
			response.errorId, _ = uuid.Parse(errorJson["id"])
		}
		response.errorMessage = errorJson["message"]
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
		response.statusCode = resp.StatusCode
		response.content = bytes.NewBuffer(body)
	}
	return response, err
}

type formData struct {
	content     *bytes.Buffer
	contentType string
}
