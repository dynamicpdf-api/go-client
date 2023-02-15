package resource

// Represents the Layout data resource used to create PDF reports.
type LayoutDataResource struct {
	Resource
}

/*
 Initializes a new instance of the `LayoutDataResource` class
 using the layout data object and a resource name.
  * @param {string }The layout data JSON file path.
  * @param {string} layoutDataResourceName The name for layout data resource.
*/
func NewLayoutDataResource(filepath string, layoutdataresourcename string) LayoutDataResource {
	var ep LayoutDataResource
	ep = LayoutDataResource{Resource: NewResourceWithPath(filepath, layoutdataresourcename)}
	ep.typeOfResource = ep.ResourceType()
	ep.LayoutDataResourceName = layoutdataresourcename
	return ep
}
func (p LayoutDataResource) ResourceType() ResourceType {
	return LayoutDataResourceType
}
