package endpoint

import "encoding/json"

// PdfSecurityInfo represents the PDF security settings returned from the API.
type PdfSecurityInfo struct {
	EncryptionType EncryptionType `json:"-"`

	// Gets or sets the encryption type.
	EncryptionTypeString string `json:"encryptionType"`

	// Gets or sets a value indicating whether the PDF document has an user password set.
	HasUserPassword bool `json:"hasUserPassword"`

	// Gets or sets a value indicating whether the PDF document has an owner password set.
	HasOwnerPassword bool `json:"hasOwnerPassword"`

	// Gets or sets if the document can be printed by the user.
	AllowPrint *bool `json:"allowPrint"`

	// Gets or sets if the document can be edited by the user.
	AllowEdit *bool `json:"allowEdit"`

	// Gets or sets if text and images can be copied to the clipboard by the user.
	AllowCopy *bool `json:"allowCopy"`

	// Gets or sets if annotations and form fields can be added, edited and modified by the user.
	AllowUpdateAnnotsAndFields *bool `json:"allowUpdateAnnotsAndFields"`

	// Gets or sets if form filling should be allowed by the user.
	AllowFormFilling *bool `json:"allowFormFilling"`

	// Gets or sets if accessibility programs should be able to read the documents text and images for the user.
	AllowAccessibility *bool `json:"allowAccessibility"`

	// Gets or sets if the document can be assembled and manipulated by the user.
	AllowDocumentAssembly *bool `json:"allowDocumentAssembly"`

	// Gets or sets if the document can be printed at a high resolution by the user.
	AllowHighResolutionPrinting *bool `json:"allowHighResolutionPrinting"`

	// Gets or sets a value indicating whether all data should be encrypted except for metadata.
	EncryptAllExceptMetadata *bool `json:"encryptAllExceptMetadata"`

	// Gets or sets a value indicating whether only file attachments should be encrypted.
	EncryptOnlyFileAttachments *bool `json:"encryptOnlyFileAttachments"`
}

func (p *PdfSecurityInfo) GetEncryptionType() EncryptionType {
	switch p.EncryptionTypeString {
	case "rc4-40":
		return EncryptionType(RC4_40)
	case "rc4-128":
		return EncryptionType(RC4_128)
	case "aes-128-cbc":
		return EncryptionType(Aes_128_Cbc)
	case "aes-256-cbc":
		return EncryptionType(Aes_256_Cbc)
	default:
		return EncryptionType(NONE)
	}
}

func (p *PdfSecurityInfo) Bytes() []byte {
	if p == nil {
		return []byte{}
	}
	jsonBytes, _ := json.Marshal(p)
	return jsonBytes
}
