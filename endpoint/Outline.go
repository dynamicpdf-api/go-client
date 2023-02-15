package endpoint

import (
	"encoding/json"

	"github.com/dynamicpdf-api/go-client/color"
	"github.com/dynamicpdf-api/go-client/input"
)

// Represents an outline.
type Outline struct {
	// Gets or sets the text of the outline.
	Text string `json:"text,omitempty"`

	// Gets or sets a value specifying if the outline is expanded.
	Expanded bool `json:"expanded,omitempty"`

	// Gets or sets the style of the outline.
	Style OutlineStyle `json:"style,omitempty"`

	// Gets or sets the Action of the outline.
	linkTo Action

	fromInputID string

	color color.Color

	colorName string

	children OutlineList
}

func newOutlineDefault() *OutlineList {
	var ep OutlineList
	return &ep
}
func newOutline(input input.Pdf) *Outline {
	var ep Outline
	ep.fromInputID = input.Id()
	return &ep
}
func newOutlineWithText(text string) *Outline {
	var ep Outline
	ep.Text = text
	return &ep
}

func newOutlineWithAction(text string, action Action) *Outline {
	var ep Outline
	ep.Text = text
	ep.setAction(action)
	return &ep
}

// Gets the Action of the outline.
func (p *Outline) Action() Action {
	return p.linkTo
}

// Sets the color of the outline.
func (p *Outline) setAction(value Action) {
	p.linkTo = value
}

// Gets the color of the outline.
func (p *Outline) Color() color.Color {
	return p.color
}

// sets the color of the outline.
func (p *Outline) SetColor(value color.Color) {
	p.color = value
	p.colorName = p.color.ColorString()
}

func (p *Outline) Children() OutlineList {
	p.children = *newOutlineDefault()
	return p.children
}

// Gets a collection of child outlines.
func (p *Outline) GetChildren() []Outline {
	if len(p.children.outlinesList) > 0 {
		return p.children.outlinesList
	}
	return nil
}

func (p Outline) MarshalJSON() ([]byte, error) {
	type Alias Outline
	return json.Marshal(&struct {
		Color       string    `json:"color,omitempty"`
		LinkTo      Action    `json:"linkTo,omitempty"`
		FromInputID string    `json:"fromInputID,omitempty"`
		GetChildren []Outline `json:"children,omitempty"`
		Alias
	}{
		Color:       p.colorName,
		LinkTo:      p.linkTo,
		FromInputID: p.fromInputID,
		GetChildren: p.children.outlinesList,
		Alias:       (Alias)(p),
	})
}
