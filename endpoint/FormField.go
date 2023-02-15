package endpoint

//Represents a form field in the PDF document.
type FormField struct {
	// Gets or sets name of the form field.
	Name string `json:"name,omitempty"`
	// Gets or sets value of the form field.
	Value string `json:"value,omitempty"`
	// Gets or sets a boolean indicating whether to flatten the form field.
	Flatten bool `json:"flatten,omitempty"`
	// Gets or sets a boolean indicating whether to remove the form field.
	Remove bool `json:"remove,omitempty"`
}

/*
  Initializes a new instance of the `FormField` class
   * using the name of the form field as a parameter.
   * @param {string} name The name of the form field.
*/
func NewFormField(name string) *FormField {
	var ep FormField
	ep.Name = name
	return &ep
}

/*
 Initializes a new instance of the `FormField` class
  * using the name of the form field as a parameter.
  * @param {string} name The name of the form field.
  * @param {string} value The value of the form field.
*/
func NewFormFieldWithValue(name string, value string) *FormField {
	var ep FormField
	ep.Name = name
	ep.Value = value
	return &ep
}
