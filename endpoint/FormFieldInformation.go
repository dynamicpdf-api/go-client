package endpoint

import "github.com/dynamicpdf-api/go-client/v2/resource"

/*
 Represents the form field information containing the collection
 of different types of field information.
*/
type FormFieldInformation struct {
	resource.Resource
	// Gets or sets a collection of `SignatureFieldInformation`
	SignatureFields []string `json:"signatureFields,omitempty"`
	// Gets or sets a collection of `TextFieldInformation`
	TextFields []string `json:"textFields,omitempty"`
	//  Gets or sets a collection of `ChoiceFieldInformation` .
	ChoiceFields []string `json:"choiceFields,omitempty"`
	// Gets or sets a collection of `ButtonFieldInformation`
	ButtonFields []string `json:"buttonFields,omitempty"`
	// Gets or sets a collection of `PushButtonInformation`
	PushButtons []string `json:"pushButtons,omitempty"`
	// Gets or sets a collection of `MultiSelectListBoxInformation`
	MultiSelectListBoxFields []string `json:"multiSelectListBoxFields,omitempty"`
}
