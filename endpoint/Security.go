package endpoint

import "encoding/json"

type Securiter interface {
	typeOfSecurity() SecurityType
}

/*
 * Represents AES 128 bit PDF document security.
 * AES 128 bit PDF security is compatible with PDF version 1.5 and higher and,
 * Adobe Acrobat Reader version 7 or higher is needed to open these documents.
 * Older readers will not be able to read documents encrypted with this security.
 */
type Security struct {
	Securiter `json:"-"`

	securityType SecurityType
	/**Gets or sets the user password. */
	UserPassword string `json:"userPassword,omitempty"`

	/**Gets or sets the owner password. */
	OwnerPassword string `json:"ownerPassword,omitempty"`

	/** Gets or sets if text and images can be copied to the clipboard by the user. */
	AllowCopy string `json:"allowCopy,omitempty"`

	/** Gets or sets if the document can be edited by the user. */
	AllowEdit string `json:"allowEdit,omitempty"`

	/** Gets or sets if the document can be printed by the user. */
	AllowPrint string `json:"allowPrint,omitempty"`

	/**  Gets or sets if annotations and form fields can be added, edited
	 * and modified by the user. */
	AllowUpdateAnnotationsAndFields string `json:"allowUpdateAnnotationsAndFields,omitempty"`

	/** Gets or sets if accessibility programs should be able to read
	 * the documents text and images for the user. */
	AllowAccessibility string `json:"allowAccessibility,omitempty"`

	/** Gets or sets if form filling should be allowed by the user. */
	AllowFormFilling string `json:"allowFormFilling,omitempty"`

	/** Gets or sets if the document can be printed at a high resolution by the user. */
	AllowHighResolutionPrinting string `json:"allowHighResolutionPrinting,omitempty"`

	/** Gets or sets if the document can be assembled and manipulated by the user. */
	AllowDocumentAssembly string `json:"allowDocumentAssembly,omitempty"`

	documentComponents EncryptDocumentComponents

	encryptMetadata bool
}

func newSecurity(userPwd string, ownerPwd string) *Security {
	var p Security
	p.UserPassword = userPwd
	p.OwnerPassword = ownerPwd
	return &p
}

func (p *Security) typeOfSecurity() SecurityType {
	return p.securityType
}

func (p Security) MarshalJSON() ([]byte, error) {
	type Alias Security
	return json.Marshal(&struct {
		DocumentComponents EncryptDocumentComponents `json:"documentComponents,omitempty"`
		EncryptMetadata    bool                      `json:"encryptMetadata,omitempty"`
		Alias
	}{
		DocumentComponents: p.documentComponents,
		EncryptMetadata:    p.encryptMetadata,
		Alias:              (Alias)(p),
	})
}
