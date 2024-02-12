package endpoint

/*
 Represents the information of a choice field in interactive forms.
A choice field contains several text items,
one or more of which may be selected as the field value.
*/
type ChoiceFieldInformation struct {
	// Gets or Sets the name of the button field.
	Name string `json:"name,omitempty"`

	// Gets or sets the type of the button field, ex: RadioButton, CheckBox etc.
	Type ChoiceFieldType `json:"type,omitempty"`

	// Gets or sets the value of the button field.
	Value string `json:"value,omitempty"`

	// Gets or Sets the default value of the button field.
	DefaultValue string `json:"default,omitempty"`

	// Gets or Sets the export value.
	ExportValue string `json:"exportValue,omitempty"`

	// Gets the collection of items.
	Items []string `json:"items,omitempty"`

	ItemExportValues []string `json:"itemExportValues,omitempty"`
}
