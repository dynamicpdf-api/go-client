package resource

import (
	"errors"
)

/*
 * Represents an image resource used to create an `ImageResource`
 * object to create PDF from images.
 */
type ImageResource struct {
	Resource
}

/**
 * Initializes a new instance of the `ImageResource` class.
 * @param {string} filePath The image file path.
 * @param {string} resourceName The name of the resource.
 */
func NewImageResourceWithResourcePath(resource string, resourceName string) ImageResource {
	var ep ImageResource
	ep.typeOfResource = ep.ResourceType()
	ep = ImageResource{Resource: NewResourceWithPath(resource, resourceName)}
	return ep
}

/**
 * Initializes a new instance of the `ImageResource` class.
 * @param {Byte[]} The byte array of the image file.
 * @param {string} resourceName The name of the resource.
 */
func NewImageResourceWithByteValue(resource []byte, resourceName string) ImageResource {
	var ep ImageResource
	ep.typeOfResource = ep.ResourceType()
	ep = ImageResource{Resource: *NewResourceWithByteValue(resource, resourceName)}
	return ep
}
func (p ImageResource) ResourceType() ResourceType {
	return ImageResourceType
}

func (p *ImageResource) fileExtension() string {

	fileHeader := []byte{}

	copy(fileHeader, "p.data")
	if IsPngImage(fileHeader) == true {
		p.Resource.mimeType1 = "image/png"
		p.Resource.fileExtension1 = ".png"
		return ".png"
	} else if IsJpegImage(fileHeader) == true {
		p.Resource.mimeType1 = "image/jpeg"
		p.Resource.fileExtension1 = ".jpeg"
		return ".jpeg"
	} else if IsGifImage(fileHeader) == true {
		p.Resource.mimeType1 = "image/gif"
		p.Resource.fileExtension1 = "gif"
		return ".gif"
	} else if IsTiffImage(fileHeader) == true {
		p.Resource.mimeType1 = "image/tiff"
		p.Resource.fileExtension1 = ".tiff"
		return ".tiff"
	} else if IsJpeg2000Image(fileHeader) == true {
		p.Resource.mimeType1 = "image/jpeg"
		p.Resource.fileExtension1 = ".jpeg"
		return ".jpeg"
	} else if IsValidBitmapImage(fileHeader) == true {
		p.Resource.mimeType1 = "image/bmp"
		p.Resource.fileExtension1 = ".bmp"
		return ".bmp"
	} else {
		errors.New("Not supported image type or invalid image.")
		return ""
	}
}
func IsJpeg2000Image(header []byte) bool {

	if (header[0] == 0x00 && header[1] == 0x00 && header[2] == 0x00 && header[3] == 0x0C && header[4] == 0x6A &&
		header[5] == 0x50 && (header[6] == 0x1A || header[6] == 0x20) && (header[7] == 0x1A || header[7] == 0x20) &&
		header[8] == 0x0D && header[9] == 0x0A && header[10] == 0x87 && header[11] == 0x0A) ||
		(header[0] == 0xFF && header[1] == 0x4F && header[2] == 0xFF && header[3] == 0x51 && header[6] == 0x00 && header[7] == 0x00) {
		return true
	}
	return false
}
func IsPngImage(header []byte) bool {
	if header[0] == 0x89 && header[1] == 0x50 && header[2] == 0x4E && header[3] == 0x47 &&
		header[4] == 0x0D && header[5] == 0x0A && header[6] == 0x1A && header[7] == 0x0A {
		return true
	}
	return false
}
func IsTiffImage(header []byte) bool {
	if (header[0] == 0x49 && header[1] == 0x49 && header[2] == 0x2A && header[3] == 0x00) ||
		(header[0] == 0x4D && header[1] == 0x4D && header[2] == 0x00 && header[3] == 0x2A) {
		return true
	}
	return false

}
func IsGifImage(header []byte) bool {
	if header[0] == 0x47 && header[1] == 0x49 && header[2] == 0x46 && header[3] == 0x38 && (header[4] == 0x37 || header[4] == 0x39) && header[5] == 0x61 {
		return true
	}
	return false
}
func IsJpegImage(header []byte) bool {
	if header[0] == 0xFF && header[1] == 0xD8 && header[2] == 0xFF {
		return true
	}
	return false
}
func IsValidBitmapImage(header []byte) bool {
	if header[0] == 0x42 && header[1] == 0x4D {
		return true
	}
	return false
}
