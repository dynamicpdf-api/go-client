package endpoint

import (
	"encoding/json"

	"github.com/dynamicpdf-api/go-client/v2/element"
	"github.com/dynamicpdf-api/go-client/v2/font"
	"github.com/dynamicpdf-api/go-client/v2/input"
)

// Represents the pdf information.
type PdfInstruction struct {
	author                    string
	subject                   string
	keywords                  string
	creator                   string
	producer                  string
	tag                       bool
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
		Tag                       bool                   `json:"tag,omitempty"`
		Title                     string                 `json:"title,omitempty"`
		PageInformation           PageInformation        `json:"pageInformation,omitempty"`
		FlattenAllFormFields      bool                   `json:"flattenAllFormFields,omitempty"`
		RetainSignatureFormFields bool                   `json:"retainSignatureFormFields,omitempty"`
		Security                  Security               `json:"security,omitempty"`
		FormField                 []FormField            `json:"formFields,omitempty"`
		Fonts                     []font.Font            `json:"fonts,omitempty"`
		Template                  []element.Template     `json:"templates,omitempty"`
		Outlines                  []Outline              `json:"outlines,omitempty"`
		Input                     []input.InputCollector `json:"inputs,omitempty"`
		*Alias
	}{
		Author:                    p.author,
		Subject:                   p.subject,
		Keywords:                  p.keywords,
		Producer:                  p.producer,
		Tag:                       p.tag,
		Creator:                   p.creator,
		Title:                     p.title,
		PageInformation:           p.pageInformation,
		FlattenAllFormFields:      p.flattenAllFormFields,
		RetainSignatureFormFields: p.retainSignatureFormFields,
		Security:                  p.security,
		FormField:                 p.formFields,
		Fonts:                     p.fonts,
		Template:                  p.template,
		Outlines:                  p.outlines,
		Input:                     p.inputs,
		Alias:                     (*Alias)(p),
	})
}
