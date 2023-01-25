package input

import (
	"encoding/json"

	"github.com/dynamicpdf-api/go-client/element"
	"github.com/dynamicpdf-api/go-client/resource"
	"github.com/google/uuid"
)

/** Represents the base class for inputs. */
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

/** Gets the template. */
func (p *Input) Template() element.Template {
	return p.templateId
}

/** Sets the template. */
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

/** Gets the top margin. */
func (p *Input) TopMargin() float32 {
	return p.topMargin
}

/** sets the top margin. */
func (p *Input) SetTopMargin(value float32) {
	p.topMargin = value
}

/** Gets the bottom margin. */
func (p *Input) BottomMargin() float32 {
	return p.bottomMargin
}

/** sets the bottom margin. */
func (p *Input) SetBottomMargin(value float32) {
	p.bottomMargin = value
}

/** Gets the left margin. */
func (p *Input) LeftMargin() float32 {
	return p.leftMargin
}

/** sets the left margin. */
func (p *Input) SetLeftMargin(value float32) {
	p.leftMargin = value
}

/** Gets the right margin. */
func (p *Input) RightMargin() float32 {
	return p.rightMargin
}

/** sets the right margin. */
func (p *Input) SetRightMargin(value float32) {
	p.rightMargin = value
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
		Alias:        (*Alias)(p),
	})
}
