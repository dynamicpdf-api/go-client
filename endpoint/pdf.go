package endpoint

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"

	"github.com/dynamicpdf-api/go-client/v2/element"
	"github.com/dynamicpdf-api/go-client/v2/font"
	"github.com/dynamicpdf-api/go-client/v2/input"
	"github.com/dynamicpdf-api/go-client/v2/resource"
)

// Represents a pdf endpoint.
type Pdf struct {
	Endpoint                  `json:"-"` // This will be skipped in JSON
	Author                    string
	Title                     string
	Subject                   string
	Creator                   string
	Keywords                  string
	Producer                  string
	Tag                       bool
	Templates                 []element.Template
	Fonts                     []font.Font
	FormFields                []FormField
	Outlines                  OutlineList
	Inputs                    []input.InputCollector
	Security                  Security
	FlattenAllFormFields      bool
	RetainSignatureFormFields bool
	pdfInstruction            *PdfInstruction
	finalResource             []resource.Resource
}

var _ EndpointProcessor = (*Pdf)(nil)

// Initializes a new instance of the `Pdf` class.
func NewPdf() *Pdf {
	var ep Pdf

	ep.Inputs = []input.InputCollector{}
	ep.Templates = []element.Template{}
	ep.pdfInstruction = NewPdfInstruction()
	return &ep
}

func (p *Pdf) initPdfInstruction() {
	p.pdfInstruction.author = p.Author
	p.pdfInstruction.creator = p.Creator
	p.pdfInstruction.flattenAllFormFields = p.FlattenAllFormFields
	p.pdfInstruction.fonts = p.Fonts
	p.pdfInstruction.formFields = p.FormFields
	p.pdfInstruction.inputs = p.Inputs
	p.pdfInstruction.keywords = p.Keywords
	p.pdfInstruction.outlines = p.Outlines.outlinesList
	p.pdfInstruction.producer = p.Producer
	p.pdfInstruction.tag = p.Tag
	p.pdfInstruction.retainSignatureFormFields = p.RetainSignatureFormFields
	p.pdfInstruction.security = p.Security
	p.pdfInstruction.subject = p.Subject
	p.pdfInstruction.template = p.Templates
	p.pdfInstruction.title = p.Title
}
func (p *Pdf) GetInstructionsJson(indent bool) *bytes.Buffer {
	var finalResource []resource.Resource
	p.initPdfInstruction()
	for _, input1 := range p.pdfInstruction.inputs {
		if input1.InputType() == "page" {
			for _, element1 := range input1.Element() {

				if element1.Resources().ResourceName != "" {
					finalResource = append(finalResource, element1.Resources())
				}

				if element1.TextFont().ResourceName != "" {
					p.pdfInstruction.fonts = append(p.pdfInstruction.fonts, element1.TextFont())
				}
			}
		}
		res := input1.Resources()
		finalResource = append(finalResource, res...)

		finalResource = append(finalResource, p.finalResource...)

		if len(input1.Template().Id) > 0 {
			p.pdfInstruction.template = append(p.pdfInstruction.template, input1.Template())
			if input1.Template().Elements != nil {
				for _, element1 := range input1.Template().Elements {
					if element1.Resources().ResourceName != "" {
						finalResource = append(finalResource, element1.Resources())
					}
					if element1.TextFont().ResourceName != "" {
						p.pdfInstruction.fonts = append(p.pdfInstruction.fonts, element1.TextFont())
					}
				}
			}
		}
	}
	p.finalResource = finalResource
	oBuf := new(bytes.Buffer)
	jsonInstruction, err := json.Marshal(p.pdfInstruction)
	if err == nil {
		if indent {
			json.Indent(oBuf, jsonInstruction, "", "  ")
		} else {
			oBuf.Write(jsonInstruction)
		}
		pkgLog.Println(oBuf)
	}
	return oBuf
}

func (p *Pdf) EndpointName() string {
	return "pdf"
}

func (p *Pdf) BaseUrl() string {
	if p.Endpoint.BaseUrl != "" {
		return p.Endpoint.BaseUrl
	} else {
		return DefaultBaseUrl
	}
}

func (p *Pdf) ApiKey() string {
	return p.Endpoint.ApiKey
}

/*
Adds additional resource to the endpoint.
  - @param {resourcePath} The resource file path.
  - @param {resourceName} The name of the resource.
*/
func (p *Pdf) AddAdditionalResource(resourcePath string, resourceName string) resource.Resource {
	additionalResource := resource.NewAdditionalResource(resourcePath, resourceName)
	if additionalResource.ResourceType() == resource.LayoutDataResourceType {
		errors.New("Layout data resources cannot be added to a DlexLayout object.")
	} else if additionalResource.ResourceType() == resource.DlexResourceType {
		errors.New("Dlex resources cannot be added to a DlexLayout object.")
	} else {
		p.finalResource = append(p.finalResource, additionalResource.Resource)
	}
	return additionalResource.Resource
}

/*
Adds additional resource to the endpoint.
  - @param {resourceData} The resource data.
  - @param {additionalResourceType} The type of the additional resource.
  - @param {resourceName} The name of the resource.
*/
func (p *Pdf) AddAdditionalResourceWithBytes(resourceData []byte, additionalResourceType resource.AdditionalResourceType, resourceName string) resource.Resource {
	typ := resource.PdfResourceType
	switch additionalResourceType {
	case resource.Font:
		typ = resource.FontResourceType
	case resource.Image:
		typ = resource.ImageResourceType
	case resource.Pdf:
		typ = resource.PdfResourceType
	case resource.Html:
		typ = resource.HtmlResourceType
	default:
		errors.New("This type of resource not allowed")
	}
	additionalResource := resource.NewAdditionalResourceWithByteValue(resourceData, resourceName, typ)
	p.finalResource = append(p.finalResource, additionalResource.Resource)
	return additionalResource.Resource
}

/*
Returns a `PdfInput` object containing the input pdf.
  - @param {resource} value The resource of type `PdfResource`
  - @param {MergeOptions} options The merge options for the pdf.
  - @returns PdfInput
*/
func (p *Pdf) AddPdf(resource resource.PdfResource, option input.MergeOptions) *input.Pdf {
	var input = input.NewPdfWithResourceWithMergerOption(resource, option)
	p.Inputs = append(p.Inputs, input)
	return input
}

/*
Returns a `PdfInput` object containing the input pdf.
  - @param {cloudResourcePath} The resource path in cloud resource manager.
  - @param {MergeOptions} options The merge options for the pdf.
  - @returns PdfInput
*/
func (p *Pdf) AddPdfCloudPath(cloudResourcePath string, option input.MergeOptions) *input.Pdf {
	var input = input.NewPdfWithCloudPath(cloudResourcePath, option)
	p.Inputs = append(p.Inputs, input)
	return input
}

/*
Returns a `ImageInput` object containing the input pdf.
  - @param {ImageResource} value The resource of type `ImageResource`.
  - @returns ImageInput
*/
func (p *Pdf) AddImage(resource resource.ImageResource) *input.Image {
	var input = input.NewImagewithImageResource(resource)
	p.Inputs = append(p.Inputs, input)
	return input
}

/*
Returns a `ImageInput` object containing the input pdf.
  - @param {string}The resource path in cloud resource manager.
  - @returns ImageInput
*/
func (p *Pdf) AddImageCloudPath(cloudResourcePath string) *input.Image {
	var input = input.NewImageWithResourcePath(cloudResourcePath)
	p.Inputs = append(p.Inputs, input)
	return input
}

/*
Returns a `HtmlInput` object containing the input pdf.
  - @param {HtmlResource} value The resource of type `HtmlResource`.
  - @returns HtmlInput
*/
func (p *Pdf) AddHtml(resource resource.HtmlResource) *input.Html {
	var input = input.NewHtmlInputWithResource(resource)
	p.Inputs = append(p.Inputs, input)
	return input
}

/*
Returns a `HtmlInput` object containing the input pdf.
  - @param {string} html The HTML input string.
  - @param {string} resource The name of the resource
  - @returns HtmlInput
*/
func (p *Pdf) AddHtmlString(html string, resourceName string) *input.Html {
	return p.AddHtml(resource.NewHtmlResource(html, resourceName))
}

/*
Returns a `WordInput` object containing the input pdf.
  - @param {WordResource} value The resource of type `WordResource`.
  - @returns WordInput
*/
func (p *Pdf) AddWord(resource resource.WordResource) *input.Word {
	var input = input.NewWordInputWithResource(resource)
	p.Inputs = append(p.Inputs, input)
	return input
}

/*
Returns a `ExcelInput` object containing the input pdf.
  - @param {ExcelResource} value The resource of type `ExcelResource`.
  - @returns ExcelInput
*/
func (p *Pdf) AddExcel(resource resource.ExcelResource) *input.Excel {
	var input = input.NewExcelInputWithResource(resource)
	p.Inputs = append(p.Inputs, input)
	return input
}

/*
Returns a `DlexInput` object containing the input pdf.
  - @param {DlexResource} value The dlex resource of type `DlexResource`.
  - @param {LayoutDataResource} value1 The layout data resource of type `LayoutDataResource`
  - @returns DlexInput
*/
func (p *Pdf) AddNewDlexWithDlexNLayoutResources(resource resource.DlexResource, layoutData resource.LayoutDataResource) *input.Dlex {
	var input = input.NewDlexWithDlexNLayoutResources(resource, layoutData)
	p.Inputs = append(p.Inputs, input)
	return input
}

/*
Returns a `DlexInput` object containing the input pdf.
  - @param {DlexResource} value The dlex resource of type `DlexResource`.
  - @param {string} value1 The layout data resource path.
  - @returns DlexInput
*/
func (p *Pdf) AddDlexWithDlexResourceNLayoutDataPath(resource resource.DlexResource, layoutDataPath string) *input.Dlex {
	var input = input.NewDlexWithDlexResourceNLayoutDataPath(resource, layoutDataPath)
	p.Inputs = append(p.Inputs, input)
	return input
}

/*
Returns a `DlexInput` object containing the input pdf.
  - @param {String} The resource path in cloud resource manager.
  - @param {String} The layout data resource path.
  - @returns DlexInput
*/
func (p *Pdf) AddDlexWithCloudResourceNLayoutDataPath(cloudResourcePath string, layoutDataPath string) *input.Dlex {
	var input = input.NewDlexWithCloudResourceNLayoutDataPath(cloudResourcePath, layoutDataPath)
	p.Inputs = append(p.Inputs, input)
	return input
}

/*
Returns a `DlexInput` object containing the input pdf.
  - @param {String} The resource path in cloud resource manager.
  - @param {LayoutDataResource} value1 The layout data resource of type `LayoutDataResource`
  - @returns DlexInput
*/
func (p *Pdf) AddDlexWithCloudResourceNLayoutData(cloudResourcePath string, layoutData resource.LayoutDataResource) *input.Dlex {
	var input = input.NewDlexWithCloudResourceNLayoutData(cloudResourcePath, layoutData)
	p.Inputs = append(p.Inputs, input)
	return input
}

/*
Returns a `PageInput` object containing the input pdf.
  - @param {float} pageWidth The width of the page.
  - @param {float} pageHeight The height of the page.
  - @returns PageInput
*/
func (p *Pdf) AddPageWithDimension(width float32, height float32) *input.Page {
	var input = input.NewPageWithDimension(width, height)
	p.Inputs = append(p.Inputs, input)
	return input
}

/*
Returns a `PageInput` object containing the input pdf.
  - @returns PageInput
*/
func (p *Pdf) AddPage() *input.Page {
	var input = input.NewPage()
	p.Inputs = append(p.Inputs, input)
	return input
}

/*
Process to create pdf.
  - @returns A Promise of PdfResponse callback.
*/
func (p *Pdf) Process() <-chan PdfResponse {
	retResponse := make(chan PdfResponse)
	go func() {
		var form formData
		form.content = bytes.NewBuffer(nil)
		formWriter := multipart.NewWriter(form.content)
		form.contentType = formWriter.FormDataContentType()

		part, err := formWriter.CreateFormFile("Instructions", "Instructions.json")
		if err != nil {
			return
		}

		io.Copy(part, p.GetInstructionsJson(false))

		for _, finalResource1 := range p.finalResource {
			formWriter.WriteField("Resource", string(finalResource1.Data()))
		}

		for _, resource1 := range p.finalResource {
			if resource1.ResourceType() == resource.LayoutDataResourceType {
				part, err := formWriter.CreateFormFile("Resource", resource1.LayoutDataResourceName)
				if err != nil {
					return
				}
				resourceData := bytes.NewBuffer(resource1.Data())
				io.Copy(part, resourceData)
			} else {
				part, err := formWriter.CreateFormFile("Resource", resource1.ResourceName)
				if err != nil {
					return
				}
				resourceData := bytes.NewBuffer(resource1.Data())
				io.Copy(part, resourceData)
			}
		}
		err = formWriter.Close()
		if err != nil {
			return
		}
		res, err := postForm(form, p)
		if err != nil {
			res.clientError = err
		}
		pdfRes := PdfResponse{response: res}
		retResponse <- pdfRes
	}()

	return retResponse
}
