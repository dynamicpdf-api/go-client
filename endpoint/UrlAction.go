package endpoint

// Represents information of a signature field.
type UrlAction struct {

	// Gets or Sets the name of a signature field.
	Name string `json:"name,omitempty"`

	// Gets or Sets the boolean, indicating the field signed or not.
	Signed bool `json:"signed,omitempty"`
}
