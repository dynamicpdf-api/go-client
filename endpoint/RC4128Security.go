package endpoint

/**
 * Represents RC4 128 bit PDF document security.
 * RC4 128 bit PDF security, with UseCryptFilter property set to false is compatible with PDF version 1.4 or higher and can be read
 * with Adobe Acrobat Reader version 5 or higher. By default UseCryptFilter property is false. RC4 128 bit PDF security with crypt filter
 * is compatible with PDF version 1.5 or higher and can be read with Adobe Acrobat Reader version 6 and higher.
 * Older readers will not be able to read document encrypted with this security.
 */
type RC4128Security struct {
	Security
}

/**
 * Initializes a new instance of the `RC4128Security` class.
 */
func NewRC4128Security() *RC4128Security {
	var p RC4128Security
	p.securityType = p.typeOfSecurity()
	return &p
}

func (p *RC4128Security) typeOfSecurity() SecurityType {
	return RC4128
}

/** Gets the documents components to be encrypted. */
func (p *RC4128Security) EncryptMetadata() bool {
	return p.encryptMetadata
}

/** Sets the documents components to be encrypted. */
func (p *RC4128Security) SetEncryptMetadata(value bool) {
	p.encryptMetadata = value
}
