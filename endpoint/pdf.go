package endpoint

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"

	"github.com/dynamicpdf-api/go-client/element"
	"github.com/dynamicpdf-api/go-client/font"
	"github.com/dynamicpdf-api/go-client/input"
	"github.com/dynamicpdf-api/go-client/resource"
)

/** Represents a pdf endpoint. */
type Pdf struct {
	Endpoint                  `json:"-"`             // This will be skipped in JSON
	Author                    string                 `json:"author"`
	Title                     string                 `json:"title"`
	Subject                   string                 `json:"subject"`
	Creator                   string                 `json:"creator"`
	Keywords                  string                 `json:"keywords"`
	Producer                  string                 `json:"producer"`
	Templates                 []element.Template     `json:"templates"`  /*TODO: Type*/
	Fonts                     []font.Font            `json:"fonts"`      /*TODO: Type*/
	FormFields                []FormField            `json:"formFields"` /*TODO: Type*/
	Outlines                  OutlineList            `json:"outlines"`   /*TODO: Type*/
	Inputs                    []input.InputCollector `json:"inputs"`
	Security                  Security               `json:"security"`
	FlattenAllFormFields      bool                   `json:"flattenAllFormFields"`
	RetainSignatureFormFields bool                   `json:"retainSignatureFormFields"`
	pdfInstruction            *PdfInstruction
}

var _ EndpointProcessor = (*Pdf)(nil)

/** Initializes a new instance of the `Pdf` class. */
func NewPdf() *Pdf {
	var ep Pdf

	ep.Endpoint.BaseUrl = DefaultBaseUrl
	ep.Endpoint.ApiKey = DefaultApiKey
	ep.Inputs = []input.InputCollector{}
	ep.Templates = []element.Template{}
	ep.pdfInstruction = NewPdfInstruction()
	return &ep
}

func (p *Pdf) GetInstructionsJson() {
	p.pdfInstruction.author = p.Author
	p.pdfInstruction.creator = p.Creator
	p.pdfInstruction.flattenAllFormFields = p.FlattenAllFormFields
	p.pdfInstruction.fonts = p.Fonts
	p.pdfInstruction.formFields = p.FormFields
	p.pdfInstruction.inputs = p.Inputs
	p.pdfInstruction.keywords = p.Keywords
	p.pdfInstruction.outlines = p.Outlines.outlinesList
	p.pdfInstruction.producer = p.Producer
	p.pdfInstruction.retainSignatureFormFields = p.RetainSignatureFormFields
	p.pdfInstruction.security = p.Security
	p.pdfInstruction.subject = p.Subject
	p.pdfInstruction.template = p.Templates
	p.pdfInstruction.title = p.Title
}

func (p *Pdf) EndpointName() string {
	return "pdf"
}

func (p *Pdf) BaseUrl() string {
	return p.Endpoint.BaseUrl
}

func (p *Pdf) ApiKey() string {
	return p.Endpoint.ApiKey
}

/**
 * Returns a `PdfInput` object containing the input pdf.
 * @param {resource} value The resource of type `PdfResource`
 * @param {MergeOptions} options The merge options for the pdf.
 * @returns PdfInput
 */
func (p *Pdf) AddPdf(resource resource.PdfResource, option input.MergeOptions) *input.Pdf {
	var input = input.NewPdfWithResourceWithMergerOption(resource, option)
	p.Inputs = append(p.Inputs, input)
	return input
}

/**
 * Returns a `PdfInput` object containing the input pdf.
 * @param {cloudResourcePath} The resource path in cloud resource manager.
 * @param {MergeOptions} options The merge options for the pdf.
 * @returns PdfInput
 */
func (p *Pdf) AddPdfCloudPath(cloudResourcePath string, option input.MergeOptions) *input.Pdf {
	var input = input.NewPdfWithCloudPath(cloudResourcePath, option)
	p.Inputs = append(p.Inputs, input)
	return input
}

/**
 * Returns a `ImageInput` object containing the input pdf.
 * @param {PdfResource} value The resource of type `ImageResource`.
 * @returns ImageInput
 */
func (p *Pdf) AddImage(resource resource.ImageResource) *input.Image {
	var input = input.NewImagewithImageResource(resource)
	p.Inputs = append(p.Inputs, input)
	return input
}

/**
 * Returns a `ImageInput` object containing the input pdf.
 * @param {string}The resource path in cloud resource manager.
 * @returns ImageInput
 */
func (p *Pdf) AddImageCloudPath(cloudResourcePath string) *input.Image {
	var input = input.NewImageWithResourcePath(cloudResourcePath)
	p.Inputs = append(p.Inputs, input)
	return input
}

/**
 * Returns a `DlexInput` object containing the input pdf.
 * @param {DlexResource} value The dlex resource of type `DlexResource`.
 * @param {LayoutDataResource} value1 The layout data resource of type `LayoutDataResource`
 * @returns DlexInput
 */
func (p *Pdf) AddNewDlexWithDlexNLayoutResources(resource resource.DlexResource, layoutData resource.LayoutDataResource) *input.Dlex {
	var input = input.NewDlexWithDlexNLayoutResources(resource, layoutData)
	p.Inputs = append(p.Inputs, input)
	return input
}

/**
 * Returns a `DlexInput` object containing the input pdf.
 * @param {DlexResource} value The dlex resource of type `DlexResource`.
 * @param {string} value1 The layout data resource path.
 * @returns DlexInput
 */
func (p *Pdf) AddDlexWithDlexResourceNLayoutDataPath(resource resource.DlexResource, layoutDataPath string) *input.Dlex {
	var input = input.NewDlexWithDlexResourceNLayoutDataPath(resource, layoutDataPath)
	p.Inputs = append(p.Inputs, input)
	return input
}

/**
 * Returns a `DlexInput` object containing the input pdf.
 * @param {String} The resource path in cloud resource manager.
 * @param {String} The layout data resource path.
 * @returns DlexInput
 */
func (p *Pdf) AddDlexWithCloudResourceNLayoutDataPath(cloudResourcePath string, layoutDataPath string) *input.Dlex {
	var input = input.NewDlexWithCloudResourceNLayoutDataPath(cloudResourcePath, layoutDataPath)
	p.Inputs = append(p.Inputs, input)
	return input
}

/**
 * Returns a `DlexInput` object containing the input pdf.
 * @param {String} The resource path in cloud resource manager.
 * @param {LayoutDataResource} value1 The layout data resource of type `LayoutDataResource`
 * @returns DlexInput
 */
func (p *Pdf) AddDlexWithCloudResourceNLayoutData(cloudResourcePath string, layoutData resource.LayoutDataResource) *input.Dlex {
	var input = input.NewDlexWithCloudResourceNLayoutData(cloudResourcePath, layoutData)
	p.Inputs = append(p.Inputs, input)
	return input
}

/**
 * Returns a `PageInput` object containing the input pdf.
 * @param {float} pageWidth The width of the page.
 * @param {float} pageHeight The height of the page.
 * @returns PageInput
 */
func (p *Pdf) AddPageWithDimension(width float32, height float32) *input.Page {
	var input = input.NewPageWithDimension(width, height)
	p.Inputs = append(p.Inputs, input)
	return input
}

/**
 * Returns a `PageInput` object containing the input pdf.
 * @returns PageInput
 */
func (p *Pdf) AddPage() *input.Page {
	var input = input.NewPage()
	p.Inputs = append(p.Inputs, input)
	return input
}

/**
 * Process to create pdf.
 * @returns A Promise of PdfResponse callback.
 */
func (p *Pdf) Process() <-chan PdfResponse {
	retResponse := make(chan PdfResponse)
	var finalResource []resource.Resource
	go func() {
		p.GetInstructionsJson()
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
			for _, r := range res {
				finalResource = append(finalResource, r)
			}
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
		var form formData
		form.content = bytes.NewBuffer(nil)
		formWriter := multipart.NewWriter(form.content)
		form.contentType = formWriter.FormDataContentType()

		part, err := formWriter.CreateFormFile("Instructions", "Instructions.json")
		if err != nil {
			return
		}

		jsonPdf, err := json.Marshal(p.pdfInstruction)
		if err == nil {
			oBuf := bytes.NewBuffer(nil)
			json.Indent(oBuf, jsonPdf, "", "\t")
			str := oBuf.String()
			pkgLog.Println("\n" + str)
		}
		jsonStr := bytes.NewBuffer(jsonPdf)
		io.Copy(part, jsonStr)

		for _, finalResource1 := range finalResource {
			formWriter.WriteField("Resource", string(finalResource1.Data()))
		}

		for _, resource1 := range finalResource {
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
