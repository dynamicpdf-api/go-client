package endpoint

/*Represents information of a MultiSelectListBox. */
type MultiSelectListBoxInformation struct {
	/**Gets or Sets the name of a MultiSelectListBox. */
	Name string `json:"name,omitempty"`

	/** Gets or sets a collection of values of the MultiSelectListBox. */
	Values []string `json:"values,omitempty"`

	/** Gets or sets a collection of default values of the MultiSelectListBox. */
	DefaultValues []string `json:"defaultValues,omitempty"`

	/** Gets or sets a collection of export values of the MultiSelectListBox. */
	ExportValues []string `json:"exportValues,omitempty"`

	/** Gets or sets a collection of items of the MultiSelectListBox. */
	Items []string `json:"items,omitempty"`

	/** Gets or sets a collection of export values of the MultiSelectListBox. */
	ItemsExportValues []string `json:"itemsExportValues,omitempty"`
}
