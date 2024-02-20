package input

import (
	"encoding/json"

	"github.com/dynamicpdf-api/go-client/v2/element"
	"github.com/dynamicpdf-api/go-client/v2/position"
	"github.com/dynamicpdf-api/go-client/v2/resource"
	"github.com/google/uuid"
)

// Represents the base class for inputs.
type InputCollector interface {
	InputType() InputType
	Template() element.Template
	Resources() []resource.Resource
	Element() []element.ElementCollector
}
type Input struct {
	InputCollector         `json:"-"`
	id                     string
	ResourceName           string `json:"resourceName,omitempty"`
	templateId             element.Template
	template_ID            string
	resources              []resource.Resource
	inputType              InputType
	Elements               []element.ElementCollector `json:"elements,omitempty"`
	LayoutDataResourceName string                     `json:"layoutDataResourceName,omitempty"`
	htmlString             string
	basePath               string
	topMargin              float32
	leftMargin             float32
	bottomMargin           float32
	rightMargin            float32
	pageWidth              float32
	pageHeight             float32
	pageSize               PageSize
	pageOrientation        Orientation
	TextReplace            []TextReplace `json:"textReplace,omitempty"`
	text                   string
	replaceText            string
	matchCase              bool
	ScaleX                 float32         `json:"scaleX,omitempty"`
	ScaleY                 float32         `json:"scaleY,omitempty"`
	VAlign                 position.VAlign `json:"vAlign,omitempty"`
	Align                  position.Align  `json:"align,omitempty"`
	ExpandToFit            bool            `json:"expandToFit,omitempty"`
	ShrinkToFit            bool            `json:"shrinkToFit,omitempty"`
	StartPage              int             `json:"startPage,omitempty"`
	PageCount              int             `json:"pageCount,omitempty"`    /*TODO:Nullability*/
	MergeOption            MergeOptions    `json:"mergeOptions,omitempty"` /*TODO:Type*/
}

func newInput(cloudResourcePath string) *Input {
	var p Input
	p.ResourceName = cloudResourcePath
	return &p
}

func newInputWithResource(resource resource.Resource) *Input {
	var p Input
	p.resources = append(p.resources, resource)
	p.ResourceName = resource.ResourceName
	return &p
}

// Gets the template.
func (p *Input) Template() element.Template {
	return p.templateId
}

// Sets the template.
func (p *Input) SetTemplate(temp element.Template) {
	p.templateId = temp
	p.template_ID = temp.Id
}

func (p *Input) Resources() []resource.Resource {
	return p.resources
}

func (p *Input) Element() []element.ElementCollector {
	return p.Elements
}

func (p *Input) Id() string {
	if p.id == "" {
		p.id = uuid.NewString()
	}

	return p.id
}

func (p *Input) SetId(value string) {
	p.id = value
}

func (p *Input) MarshalJSON() ([]byte, error) {
	type Alias Input
	return json.Marshal(&struct {
		TemplateId   string    `json:"templateId,omitempty"`
		InputType    InputType `json:"type,omitempty"`
		Id           string    `json:"id,omitempty"`
		BasePath     string    `json:"basePath,omitempty"`
		TopMargin    float32   `json:"topMargin,omitempty"`
		LeftMargin   float32   `json:"leftMargin,omitempty"`
		BottomMargin float32   `json:"bottomMargin,omitempty"`
		RightMargin  float32   `json:"rightMargin,omitempty"`
		PageWidth    float32   `json:"pageWidth,omitempty"`
		PageHeight   float32   `json:"pageHeight,omitempty"`
		Text         string    `json:"text,omitempty"`
		ReplaceText  string    `json:"replaceText,omitempty"`
		MatchCase    bool      `json:"matchCase,omitempty"`

		*Alias
	}{
		TemplateId:   p.template_ID,
		InputType:    p.inputType,
		Id:           p.id,
		BasePath:     p.basePath,
		TopMargin:    p.topMargin,
		LeftMargin:   p.leftMargin,
		BottomMargin: p.bottomMargin,
		RightMargin:  p.rightMargin,
		PageWidth:    p.pageWidth,
		PageHeight:   p.pageHeight,
		Text:         p.text,
		ReplaceText:  p.replaceText,
		MatchCase:    p.matchCase,
		Alias:        (*Alias)(p),
	})
}
