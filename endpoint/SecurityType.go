package endpoint

type SecurityType int

/**
 * Represents information of a PushButton field.
 */
const (

	/** AES 128 bit security. */
	Aes128 SecurityType = 0

	/** AES 256 bit security. */
	Aes256 SecurityType = 1

	/** RC4 128 bit security. */
	RC4128 SecurityType = 2
)
