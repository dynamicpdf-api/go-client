package resource

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/uuid"
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

func newResource() Resource {
	var ep Resource
	return ep
}

func NewResourceWithPath(filePath string, resourceName string) Resource {
	var ep Resource
	_, err := os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		panic(err)
	} else {
		ep.setFilepath(filePath)
		ep.data = ep.getFileData()
	}
	if resourceName == "" {
		ep.ResourceName = uuid.New().String() + ep.fileExtension()
	} else {
		ep.ResourceName = resourceName
	}
	return ep
}

func NewResourceWithByteValue(value []byte, resourceName string) *Resource {
	var ep Resource
	ep.data = value
	//ep.typeOfResource = ep.ResourceType()
	if resourceName == "" {
		ep.ResourceName = uuid.New().String() + ep.fileExtension()
	} else {
		ep.ResourceName = resourceName
	}
	return &ep
}

func newResourceWithStream(stream io.Reader, resourceName string) Resource {
	var ep Resource
	ep.data = ep.getStreamData(stream)
	if resourceName == "" {
		ep.ResourceName = uuid.New().String() + ep.fileExtension()
	} else {
		ep.ResourceName = resourceName
	}
	return ep
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

func (p *Resource) setResourceType(mimeType ResourceType) {
	p.typeOfResource = mimeType
	p.getFileData()
}
func (p *Resource) filepath() string {
	return p.filepath1
}

func (p *Resource) setFilepath(path string) {
	p.filepath1 = path
	ext := filepath.Ext(path)
	p.setFileExtension(ext)
}

func (p *Resource) getFileData() []byte {
	data, err := ioutil.ReadFile(p.filepath1)
	if err == nil {
		p.data = data
	}
	return data
}

func (p *Resource) getStreamData(input io.Reader) []byte {
	if input == nil {
		panic("stream is nil")
	}

	data, err := io.ReadAll(input)
	if err != nil {
		panic(fmt.Sprintf("failed to read stream: %v", err))
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
