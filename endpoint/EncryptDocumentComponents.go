package endpoint

/** Specifies the document components to be encrypted. */
type EncryptDocumentComponents int

const (
	/** Encrypts all document contents. */
	All EncryptDocumentComponents = iota
	/** Encrypts all document contents except metadata. */
	AllExceptMetadata EncryptDocumentComponents = iota
)
