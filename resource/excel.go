package resource

import (
	"path"
)

/*
Represents a excel resource object that is created using the excel file and a name.
*/
type ExcelResource struct {
	Resource
	inputFileExtension string
	// resourceName       string
	// resource           string
}

/*
Initializes a new instance of the `excelResource` class with excel file path and resource name.
  - @param {string} excel The excel file path.
  - @param {string} resourceName The name of the resource.
*/
func NewExcelResourceWithResourcePath(resource string, resourceName string) ExcelResource {
	var ep ExcelResource
	ep.setFilepath(resource)
	ep.typeOfResource = ep.ResourceType()
	ep = ExcelResource{Resource: NewResourceWithPath(resource, resourceName)}
	ep.fileExtension()
	return ep
}

/*
Initializes a new instance of the `ExcelResource` class.
  - @param {Byte[]} The byte array of the Excel file.
  - @param {string} resourceName The name of the resource.
*/
func NewExcelResourceWithByteValue(resource []byte, resourceName string) ExcelResource {
	var ep ExcelResource
	ep.typeOfResource = ep.ResourceType()
	ep = ExcelResource{Resource: *NewResourceWithByteValue(resource, resourceName)}
	return ep
}
func (p ExcelResource) ResourceType() ResourceType {
	return ExcelResourceType
}

func (p *ExcelResource) fileExtension() string {

	p.inputFileExtension = ""
	if p.ResourceName != "" {
		p.inputFileExtension = path.Ext(p.ResourceName)
	} else if p.filepath1 != "" {
		p.inputFileExtension = path.Ext(p.filepath1)

	} else {
		panic("Invalid path or resource name.")
	}
	if p.inputFileExtension == ".xls" {
		p.setMimeType("application/vnd.ms-excel")
		return ".xls"
	} else if p.inputFileExtension == ".xlsx" && p.data[0] == 0x50 && p.data[1] == 0x4b {
		p.setMimeType("application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		return ".xlsx"
	} else {
		panic("Unsupported file type or invalid file extension.")
	}
}
