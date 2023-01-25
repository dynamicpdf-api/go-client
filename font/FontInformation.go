package font

import "github.com/dynamicpdf-api/go-client/resource"

type FontInformation struct {
	resource.Resource
	fontName string
	filePath string
}

func NewFontInformation() *FontInformation {
	var ep FontInformation
	return &ep
}

func (p *FontInformation) FontName() string {
	return p.fontName
}

func (p *FontInformation) FilePath() string {
	return p.filePath
}
