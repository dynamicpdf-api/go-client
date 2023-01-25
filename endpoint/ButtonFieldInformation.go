package endpoint

type ButtonFieldInformation struct {
	/* Gets or Sets the name of the button field.*/
	Name string `json:"name,omitempty"`

	/* Gets or sets the type of the button field, ex: RadioButton, CheckBox etc. */
	Type int `json:"type,omitempty"`

	/* Gets or sets the value of the button field. */
	Value string `json:"value,omitempty"`

	/*Gets or Sets the default value of the button field. */
	DefaultValue string `json:"defaultValue,omitempty"`

	/*Gets or Sets the export value. These values will be exported when submitting the form.
	 * To create a set of mutually exclusive radio buttons
	 * (i.e., where only one can be selected at a time),
	 * create the fields with the same name but different export values.
	 */
	ExportValue string `json:"exportValue,omitempty"`

	/* Gets the collection of export value.*/
	ExportValues []string `json:"exportValues,omitempty"`
}
