package endpoint

// Specifies Encryption Type.
type EncryptionType string

const (

	// Represents a RC4 40 bit security.
	RC4_40 EncryptionType = "RC440"

	// Represents a RC4 128 bit security.
	RC4_128 EncryptionType = "RC4128"

	// Represents a AES 128 bit security with CBC cipher mode.
	Aes_128_Cbc EncryptionType = "Aes128Cbc"

	// Represents a AES 256 bit security with CBC cipher mode.
	Aes_256_Cbc EncryptionType = "Aes256Cbc"

	// Represents a No security.
	NONE EncryptionType = "None"
)
