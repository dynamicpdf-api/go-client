package endpoint

import (
	"encoding/json"

	"github.com/dynamicpdf-api/go-client/element"
	"github.com/dynamicpdf-api/go-client/font"
	"github.com/dynamicpdf-api/go-client/input"
)

// Represents the pdf information.
type PdfInstruction struct {
	author                    string
	subject                   string
	keywords                  string
	creator                   string
	producer                  string
	title                     string
	pageInformation           PageInformation
	flattenAllFormFields      bool
	retainSignatureFormFields bool
	security                  Security
	formFields                []FormField
	fonts                     []font.Font
	template                  []element.Template
	outlines                  []Outline
	inputs                    []input.InputCollector
}

func NewPdfInstruction() *PdfInstruction {
	var ep PdfInstruction
	ep.author = "Cete Software"
	ep.creator = "DynamicPDF Cloud Api"
	return &ep
}

func (p *PdfInstruction) GetOutlines() []Outline {
	return p.outlines
}

func (p *PdfInstruction) MarshalJSON() ([]byte, error) {
	type Alias PdfInstruction
	return json.Marshal(&struct {
		Author                    string                 `json:"author,omitempty"`
		Subject                   string                 `json:"subject,omitempty"`
		Keywords                  string                 `json:"keywords,omitempty"`
		Creator                   string                 `json:"creator,omitempty"`
		Producer                  string                 `json:"producer,omitempty"`
		Title                     string                 `json:"title,omitempty"`
		PageInformation           PageInformation        `json:"pageInformation,omitempty"`
		FlattenAllFormFields      bool                   `json:"flattenAllFormFields,omitempty"`
		RetainSignatureFormFields bool                   `json:"retainSignatureFormFields,omitempty"`
		Security                  Security               `json:"security,omitempty"`
		FormField                 []FormField            `json:"formFields,omitempty"`
		Font                      []font.Font            `json:"font,omitempty"`
		Template                  []element.Template     `json:"templates,omitempty"`
		Outlines                  []Outline              `json:"outlines,omitempty"`
		Input                     []input.InputCollector `json:"inputs,omitempty"`
		*Alias
	}{
		Author:                    p.author,
		Subject:                   p.subject,
		Keywords:                  p.keywords,
		Creator:                   p.creator,
		Title:                     p.title,
		PageInformation:           p.pageInformation,
		FlattenAllFormFields:      p.flattenAllFormFields,
		RetainSignatureFormFields: p.retainSignatureFormFields,
		Security:                  p.security,
		FormField:                 p.formFields,
		Font:                      p.fonts,
		Template:                  p.template,
		Outlines:                  p.outlines,
		Input:                     p.inputs,
		Alias:                     (*Alias)(p),
	})
}
