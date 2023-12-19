package resource

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Represents the base class resource.
type Resource struct {
	// Gets or sets the resource name.
	ResourceName           string `json:"resourceName,omitempty"`
	LayoutDataResourceName string `json:"layoutDataResourceName,omitempty"`
	filepath1              string
	typeOfResource         ResourceType
	fileExtension1         string
	mimeType1              string
	data                   []byte
}

func NewResourceWithPath(input string, resourceName string) Resource {
	var ep Resource
	ep.ResourceName = resourceName
	_, err := os.Stat(input)
	if errors.Is(err, os.ErrNotExist) {
		panic(err)
	} else {
		ep.setFilepath(input)
		ep.data = ep.getFileData()
		ep.typeOfResource = ep.getResourceType(input)
	}
	return ep
}


func NewResourceWithByteValue(input []byte, resourceName string) *Resource {
	var ep Resource
	ep.ResourceName = resourceName
	ep.data = input
	ep.typeOfResource = ep.ResourceType()
	return &ep
}

func (p *Resource) Data() []byte {
	return p.data
}

func (p *Resource) setData(data []byte) {
	p.data = data
}

func (p *Resource) fileExtension() string {
	return p.fileExtension1
}

func (p *Resource) setFileExtension(fileExtension string) {
	p.fileExtension1 = fileExtension
}

func (p *Resource) mimeType() string {
	return p.mimeType1
}

func (p *Resource) setMimeType(mimeType string) {
	p.mimeType1 = mimeType
}

func (p *Resource) ResourceType() ResourceType {
	return p.typeOfResource
}

func (p *Resource) getResourceType(resourcePath string) ResourceType {
	typ := PdfResourceType
	fileExtension := filepath.Ext(resourcePath)
	switch fileExtension {
	case ".pdf":
		typ = PdfResourceType
		break
	case ".dlex":
		typ = DlexResourceType
		break
	case ".json":
		typ = LayoutDataResourceType
		break
	case ".ttf", ".otf":
		typ = FontResourceType
		break
	case ".tiff", ".tif", ".png", ".gif", ".jpeg", ".jpg", ".bmp":
		typ = ImageResourceType
		break
	case ".html":
		typ = HtmlResourceType
		break
	}
	return typ
}

func (p *Resource) setResourceType(mimeType ResourceType) {
	p.typeOfResource = mimeType
	p.getFileData()
}
func (p *Resource) filepath() string {
	return p.filepath1
}

func (p *Resource) setFilepath(filePath string) {
	p.filepath1 = filePath
}

func (p *Resource) getFileData() []byte {
	data, err := ioutil.ReadFile(p.filepath1)
	if err == nil {
		p.data = data
	}
	return data
}

func (p *Resource) MarshalJSON() ([]byte, error) {
	type Alias Resource
	return json.Marshal(&struct {
		Filepath     string       `json:"filepath,omitempty"`
		ResourceType ResourceType `json:"resourceType,omitempty"`
		*Alias
	}{
		Filepath:     p.filepath1,
		ResourceType: p.typeOfResource,
		Alias:        (*Alias)(p),
	})
}
