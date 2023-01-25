package endpoint

import (
	"encoding/json"

	"github.com/dynamicpdf-api/go-client/input"
)

/**Represents an outlineList. */
type OutlineList struct {
	outlinesList []Outline
}

/**
 * Adds an `Outline` object to the outline list.
 *
 * @param {string} text Text of the outline.
 * @returns outline
 */
func (p *OutlineList) Add(value string) Outline {
	var outline = newOutlineWithText(value)
	p.outlinesList = append(p.outlinesList, *outline)
	return *outline
}

/**
 * Adds an `Outline` object to the outline list.
 *
 * @param {string} text Text of the outline.
 * @param {string} input URL the action launches.
 * @returns outline
 */
func (p *OutlineList) AddWithURL(value string, url Action) Outline {
	var outline = newOutlineWithAction(value, url)
	p.outlinesList = append(p.outlinesList, *outline)
	return *outline
}

/**
 * Adds an `Outline` object to the outline list.
 *
 * @param {string} text Text of the outline.
 * @param {Input} Any of the `ImageInput`, `DlexInput`, `PdfInput`,
 * @param {*} pgOffset Page number to navigate.
 * @param {*} pgZoom Page Zoom to display the destination.
 * @returns outline
 */
func (p *OutlineList) AddWithInputValue(text string, input input.Input, pageOffSet int, pageZoom PageZoom) *Outline {
	var linkTo = NewGoToAction(input, pageOffSet, pageZoom)
	var outline = newOutlineWithAction(text, linkTo.Action)
	p.outlinesList = append(p.outlinesList, *outline)
	return outline
}

func (p *OutlineList) AddPdfOutlines(value input.Pdf) {
	p.outlinesList = append(p.outlinesList, *newOutline(value))
}

func (p *OutlineList) outlines() []Outline {
	return p.outlinesList
}

func (p OutlineList) MarshalJSON() ([]byte, error) {
	type Alias OutlineList
	return json.Marshal(&struct {
		OutlinesList []Outline `json:"Outlines,omitempty"`
		Alias
	}{
		OutlinesList: p.outlinesList,
		Alias:        (Alias)(p),
	})
}
