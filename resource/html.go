package resource

/* Represents the pdf resource. */
type HtmlResource struct {
	Resource
}

/**
 * Initializes a new instance of the `HtmlResource` class
 * with html string and resource name.
 * @param {string} html The Html string.
 * @param {string} resource The name of the resource.
 */
func NewHtmlResource(resource string, resourceName string) HtmlResource {
	var ep HtmlResource
	ep = HtmlResource{Resource: NewResourceWithPath(resource, resourceName)}
	ep.typeOfResource = ep.ResourceType()
	return ep
}

func (p HtmlResource) ResourceType() ResourceType {
	return HtmlResourceType
}
