package endpoint

/*
  Represents AES 128 bit PDF document security.
AES 128 bit PDF security is compatible with PDF version 1.5 and higher and,
Adobe Acrobat Reader version 7 or higher is needed to open these documents.
Older readers will not be able to read documents encrypted with this security.
*/
type Aes128Security struct {
	Security
}

/*
Initializes a new instance of the `Aes128Security` class by
taking the owner and user passwords as parameters to create PDF.
 * @param {string} userPassword The owner password to open the document.
 * @param {string} ownerPassword The user password to open the document.
*/
func NewAes128Security(userPwd string, ownerPwd string) *Aes128Security {
	var p = Aes128Security{Security: *newSecurity(userPwd, ownerPwd)}
	p.securityType = p.typeOfSecurity()
	return &p
}

/*
 Gets the `EncryptDocumentComponents` components of the document to be encrypted.
We can encrypt all the PDF content or the content, excluding the metadata.
*/
func (p *Aes128Security) DocumentComponents() EncryptDocumentComponents {
	return p.documentComponents
}

/*
 sets the `EncryptDocumentComponents` components of the document to be encrypted.
We can encrypt all the PDF content or the content, excluding the metadata.
*/
func (p *Aes128Security) SetDocumentComponents(value EncryptDocumentComponents) {
	p.documentComponents = value
}
