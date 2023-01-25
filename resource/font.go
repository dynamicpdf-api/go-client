package resource

import (
	"errors"
)

type FontResource struct {
	Resource
}

func NewFontResource() *FontResource {
	var ep FontResource
	ep.typeOfResource = ep.ResourceType()
	return &ep
}
func (p FontResource) ResourceType() ResourceType {
	return FontResourceType
}
func (p FontResource) FileExtension() string {
	if p.Resource.data[0] == 0x4f && p.Resource.data[1] == 0x54 && p.Resource.data[2] == 0x54 && p.Resource.data[3] == 0x4f {
		p.Resource.mimeType1 = "font/otf"
		return ".otf"
	} else if p.Resource.data[0] == 0x00 && p.Resource.data[1] == 0x01 && p.Resource.data[2] == 0x00 && p.Resource.data[3] == 0x00 {
		p.Resource.mimeType1 = "font/ttf"
		return ".ttf"
	} else {
		errors.New("Unsupported font")
		return ""
	}
}
