package resource

import (
	"path"
)

/*
Represents a word resource object that is created using the word file and a name.
*/
type WordResource struct {
	Resource
	inputFileExtension string
	resourceName       string
	resource           string
}

/*
Initializes a new instance of the `wordResource` class with word file path and resource name.
  - @param {string} word The word file path.
  - @param {string} resourceName The name of the resource.
*/
func NewWordResourceWithResourcePath(resource string, resourceName string) WordResource {
	var ep WordResource
	ep.setFilepath(resource)
	ep.typeOfResource = ep.ResourceType()
	ep = WordResource{Resource: NewResourceWithPath(resource, resourceName)}
	ep.fileExtension()
	return ep
}

/*
Initializes a new instance of the `WordResource` class.
  - @param {Byte[]} The byte array of the Word file.
  - @param {string} resourceName The name of the resource.
*/
func NewWordResourceWithByteValue(resource []byte, resourceName string) WordResource {
	var ep WordResource
	ep.typeOfResource = ep.ResourceType()
	ep = WordResource{Resource: *NewResourceWithByteValue(resource, resourceName)}
	return ep
}
func (p WordResource) ResourceType() ResourceType {
	return WordResourceType
}

func (p *WordResource) fileExtension() string {

	p.inputFileExtension = ""
	if p.ResourceName != "" {
		p.inputFileExtension = path.Ext(p.ResourceName)
	} else if p.filepath1 != "" {
		p.inputFileExtension = path.Ext(p.filepath1)

	} else {
		panic("Invalid path or resource name.")
	}
	if p.inputFileExtension == ".doc" {
		p.setMimeType("application/msword")
		return ".doc"
	} else if p.inputFileExtension == ".docx" && p.data[0] == 0x50 && p.data[1] == 0x4b {
		p.setMimeType("application/vnd.openxmlformats-officedocument.wordprocessingml.document")
		return ".docx"
	} else if p.inputFileExtension == ".odt" && p.data[0] == 0x50 && p.data[1] == 0x4b {
		p.setMimeType("application/vnd.oasis.opendocument.text")
		return ".odt"
	} else {
		panic("Unsupported file type or invalid file extension.")
	}
}
