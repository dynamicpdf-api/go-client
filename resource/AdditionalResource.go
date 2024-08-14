package resource

import (
	"errors"
	"path/filepath"
)

type AdditionalResource struct {
	Resource
}

func NewAdditionalResource(resourcePath string, resourceName string) AdditionalResource {
	var p = AdditionalResource{Resource: NewResourceWithPath(resourcePath, resourceName)}
	p.typeOfResource = p.getResourceType(resourcePath)
	return p
}

func NewAdditionalResourceWithByteValue(resourceData []byte, resourceName string, resourceType ResourceType) AdditionalResource {
	var p = AdditionalResource{Resource: *NewResourceWithByteValue(resourceData, resourceName)}
	p.typeOfResource = resourceType
	return p
}

func (p *AdditionalResource) getResourceType(resourcePath string) ResourceType {
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

func (p *AdditionalResource) fileExtension() string {
	// func (p *AdditionalResource) determineFileExtension() (string, error) {
	switch p.typeOfResource {
	case ImageResourceType:
		fileHeader := []byte{}
		copy(fileHeader, "p.data")
		if IsPngImage(fileHeader) {
			p.setMimeType("image/png")
			return ".png"
		} else if IsJpegImage(fileHeader) {
			p.setMimeType("image/jpeg")
			return ".jpeg"
		} else if IsGifImage(fileHeader) {
			p.setMimeType("image/gif")
			return ".gif"
		} else if IsTiffImage(fileHeader) {
			p.setMimeType("image/tiff")
			return ".tiff"
		} else if IsJpeg2000Image(fileHeader) {
			p.setMimeType("image/jpeg")
			return ".jpeg"
		} else if IsValidBitmapImage(fileHeader) {
			p.setMimeType("image/bmp")
			return ".bmp"
		} else {
			errors.New("Not supported image type or invalid image.")
			return ""
		}
	case PdfResourceType:
		p.setMimeType("application/pdf")
		return ".pdf"
	case LayoutDataResourceType:
		p.setMimeType("application/json")
		return ".json"
	case DlexResourceType:
		p.setMimeType("application/xml")
		return ".dlex"
	case FontResourceType:
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
	case HtmlResourceType:
		p.setMimeType("text/html")
		return ".html"
	default:
		errors.New("unknown resource type")
		return ""
	}
}
