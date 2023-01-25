package endpoint

/*
 * Represents information of a text field. */
type TextFieldInformation struct {

	/* Gets or Sets the name of the Text field. */
	Name string `json:"name,omitempty"`

	/* Gets or Sets the value of the Text field. */
	Value string `json:"value,omitempty"`

	/* Gets or Sets the default value of the Text field. */
	DefaultValue string `json:"defaultValue,omitempty"`
}
