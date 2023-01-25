package endpoint

/*
 * Represents AES 256 bit PDF document security.
 * AES 128 bit PDF security is compatible with PDF version 1.5 and higher.
 */
type Aes256Security struct {
	Security
}

/**
 * Initializes a new instance of the `Aes256Security` class by
 * taking the owner and user passwords as parameters to create PDF.
 * @param {string} userPassword The owner password to open the document.
 * @param {string} ownerPassword The user password to open the document.
 */
func NewAes256Security(userPwd string, ownerPwd string) *Aes256Security {
	var p = Aes256Security{Security: *newSecurity(userPwd, ownerPwd)}
	p.securityType = p.typeOfSecurity()
	return &p
}

func (p *Aes256Security) typeOfSecurity() SecurityType {
	return Aes256
}

/**
 * Gets the `EncryptDocumentComponents`, components of the document to be encrypted.
 * We can encrypt all the PDF content or the content, excluding the metadata.
 */
func (p *Aes256Security) DocumentComponents() EncryptDocumentComponents {
	return p.documentComponents
}

/**
 * Sets the `EncryptDocumentComponents`, components of the document to be encrypted.
 * We can encrypt all the PDF content or the content, excluding the metadata.
 */
func (p *Aes256Security) SetDocumentComponents(value EncryptDocumentComponents) {
	p.documentComponents = value
}
