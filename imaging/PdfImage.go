package imaging

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/dynamicpdf-api/go-client/v2/endpoint"
	"github.com/dynamicpdf-api/go-client/v2/resource"
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
	restResponse := make(chan PdfImageResponse)
	go func() {
		var form formData
		form.content = bytes.NewBuffer(nil)
		formWriter := multipart.NewWriter(form.content)
		form.contentType = formWriter.FormDataContentType()

		part, err := formWriter.CreateFormFile("pdf", p.resource.ResourceName)
		if err != nil {
			panic(err)
		}
		resourceData := bytes.NewBuffer(p.resource.Data())
		io.Copy(part, resourceData)

		err = formWriter.Close()
		if err != nil {
			panic(err)
		}

		// var response response = response{isSuccessful: false}

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

		ctx := context.TODO() // TODO: Need to add a context property for user to be able to cancel/timeout the request.
		req, err := http.NewRequestWithContext(ctx, "POST", u.String(), form.content)
		if err != nil {
			pkgLog.Printf("%s\n", err)
			panic(err)
		}

		req.Header.Add("Authorization", postAuth)
		req.Header.Add("Content-Type", form.contentType)
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil || resp.Status != "200 OK" {
			pkgLog.Printf("\nError: %s\n", err)
			panic(resp.Status)
		}

		rasterizerResponse := PdfImageResponse{}

		if resp.Status == "200 OK" {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			decoder := json.NewDecoder(bytes.NewReader([]byte(bodyBytes)))

			var pdfImageResponse PdfImageResponse
			err = decoder.Decode(&pdfImageResponse)
			if err != nil {
				panic(err)
			}

			imageType := pdfImageResponse.ContentType
			rasterizerResponse.ContentType = imageType
			rasterizerResponse.HorizontalDpi = pdfImageResponse.HorizontalDpi
			rasterizerResponse.VerticalDpi = pdfImageResponse.VerticalDpi
			rasterizerResponse.Images = pdfImageResponse.Images
			rasterizerResponse.ImageFormat = strings.Split(imageType, "/")[1]

			rasterizerResponse.isSuccessful = true
			rasterizerResponse.statusCode = resp.StatusCode
		} else {
			rasterizerResponse.isSuccessful = false
			rasterizerResponse.statusCode = resp.StatusCode
			rasterizerResponse.errorJson = resp.Status
			rasterizerResponse.clientError = err
		}
		restResponse <- rasterizerResponse
	}()
	return restResponse
}

type formData struct {
	content     *bytes.Buffer
	contentType string
}
