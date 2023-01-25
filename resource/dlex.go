package resource

/**
 * Represents a Dlex resource object that is created using the DLEX file and a name.
 */
type DlexResource struct {
	Resource
	fileExtension          string
	layoutDataResourceName string
}

/**
 * Initializes a new instance of the `DlexResource` class
 * with DLEX file path and resource name or
 * byte data of the DLEX file and resource name as parameters.
 * @param {string} dlex The dlex file path.
 * @param {string} resource The name of the resource.
 */
func NewDlexResourceWithPath(dlexPath string, resourceName string) *DlexResource {
	var ep DlexResource
	ep = DlexResource{Resource: NewResourceWithPath(dlexPath, resourceName)}
	ep.typeOfResource = ep.ResourceType()
	return &ep
}

/**
 * Initializes a new instance of the `DlexResource` class
 * with DLEX file path and resource name or
 * byte data of the DLEX file and resource name as parameters.
 * @param {byte[]} The byte array of the dlex file.
 * @param {string} resource The name of the resource.
 */
func NewDlexResourceWithByteValue(value []byte, resourceName string) *DlexResource {
	var ep DlexResource
	ep = DlexResource{Resource: *NewResourceWithByteValue(value, resourceName)}
	ep.typeOfResource = ep.ResourceType()
	return &ep
}

/** Gets name for layout data resource. */
func (p DlexResource) LayoutDataResourceName() string {
	return p.layoutDataResourceName
}

/** Sets name for layout data resource. */
func (p DlexResource) SetLayoutDataResourceName(layoutData string) {
	p.layoutDataResourceName = layoutData
}

func (p DlexResource) ResourceType() ResourceType {
	return DlexResourceType
}
