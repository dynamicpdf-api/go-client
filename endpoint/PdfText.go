package endpoint

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/dynamicpdf-api/go-client/v2/resource"
)

// Represents the pdf text endpoint.
type PdfText struct {
	Endpoint

	resource resource.PdfResource

	startPage int

	pageCount int

	textOrder TextOrder
}

/*
Initializes a new instance of the `PdfText` class.
  - @param { PdfResource } resource The image resource of type `PdfResource`.`
  - @param { int } startPage The start page.
  - @param { int } pageCount The page count.
*/
func NewPdfText(resource resource.PdfResource, startpage int, pagecount int) *PdfText {
	var ep PdfText
	ep.resource = resource
	ep.startPage = startpage
	ep.pageCount = pagecount
	return &ep
}

/*
Initializes a new instance of the `PdfText` class.
  - @param { PdfResource } resource The image resource of type `PdfResource`.`
  - @param { int } startPage The start page.
  - @param { int } pageCount The page count.
  - @param { TextOrder } textOrder The text extraction order.
*/
func NewPdfTextWithOrder(resource resource.PdfResource, startpage int, pagecount int, textOrder TextOrder) *PdfText {
	var ep PdfText
	ep.resource = resource
	ep.startPage = startpage
	ep.pageCount = pagecount
	ep.textOrder = textOrder
	return &ep
}

var _ EndpointProcessor = (*PdfText)(nil)

func (p *PdfText) BaseUrl() string {
	if p.Endpoint.BaseUrl != "" {
		return p.Endpoint.BaseUrl
	} else {
		return DefaultBaseUrl
	}
}

func (p *PdfText) ApiKey() string {
	return p.Endpoint.ApiKey
}

// Gets the start page.
func (p *PdfText) StartPage() int {
	return p.startPage
}

// Sets the start page.
func (p *PdfText) SetStartPage(startPage int) {
	p.startPage = startPage
}

// Gets the page count.
func (p *PdfText) PageCount() int {
	return p.pageCount
}

// Sets the page count.
func (p *PdfText) SetPageCount(pageCount int) {
	p.pageCount = pageCount
}

// Gets the text extraction order.
func (p *PdfText) TextOrder() TextOrder {
	return p.textOrder
}

// Sets the text extraction order.
func (p *PdfText) SetTextOrder(textOrder TextOrder) {
	p.textOrder = textOrder
}

func (p *PdfText) EndpointName() string {
	return "pdf-text"
}

/*
Process the pdf resource to get pdf's text.
  - @returns A Promise of PdfTextResponse callback.
*/
func (p *PdfText) Process() <-chan PdfTextResponse {
	restResponse := make(chan PdfTextResponse)
	postUrl := strings.TrimSuffix(p.BaseUrl(), "/") + "/v1.0/" + p.EndpointName()

	url := postUrl + "/?StartPage=" + strconv.Itoa(p.startPage) + "&PageCount=" + strconv.Itoa(p.pageCount)

	if p.textOrder != "" {
		url += "&TextOrder=" + string(p.textOrder)
	}

	postAuth := "Bearer " + p.Endpoint.ApiKey
	go func() {
		res, err := postHttpRequest(url, p.resource.Data(), postAuth, "application/pdf")
		if err != nil {
			res.clientError = err
		}

		bodyBytes := res.content.Bytes()

		var content []PdfContent
		err = json.Unmarshal(bodyBytes, &content)
		if err != nil {
			panic(err)
		}

		pdfRes := PdfTextResponse{
			TextContent:  content,
			JsonResponse: res,
		}
		restResponse <- pdfRes

	}()
	return restResponse
}

func (p PdfText) MarshalJSON() ([]byte, error) {
	type Alias PdfText
	return json.Marshal(&struct {
		StartPage    int                  `json:"startPage,omitempty"`
		PageCount    int                  `json:"pageCount,omitempty"`
		TextOrder    TextOrder            `json:"textOrder,omitempty"`
		ResourceName resource.PdfResource `json:"resourceName,omitempty"`
		Alias
	}{
		StartPage:    p.startPage,
		PageCount:    p.pageCount,
		TextOrder:    p.textOrder,
		ResourceName: p.resource,
		Alias:        (Alias)(p),
	})
}
