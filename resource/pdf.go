package resource

// Represents the pdf resource.
type PdfResource struct {
	Resource
	fileExtension string
}

/*
 Initializes a new instance of the `PdfResource` class.
  * @param {string} input The pdf file path.
  * @param {string} resourceName The name of the resource.
*/
func NewPdfResourceWithResourcePath(resource string, resourceName string) PdfResource {
	var ep PdfResource
	ep.typeOfResource = ep.ResourceType()
	ep = PdfResource{Resource: NewResourceWithPath(resource, resourceName)}
	return ep
}

/*
 Initializes a new instance of the `PdfResource` class.
  * @param {Byte[]}The byte array of the pdf file.
  * @param {string} resourceName The name of the resource.
*/
func NewPdfResourceWithByteValue(resource string, resourceName string) PdfResource {
	var ep PdfResource
	ep.typeOfResource = ep.ResourceType()
	ep = PdfResource{Resource: NewResourceWithPath(resource, resourceName)}
	return ep
}

func (p PdfResource) ResourceType() ResourceType {
	return PdfResourceType
}
