package endpoint

/* Represents the pdf information. */
type PdfInformation struct {

	/* Gets the author. */
	Author string `json:"author,omitempty"`

	/* Gets the subject. */
	Subject string `json:"subject,omitempty"`

	/* Gets the keywords. */
	Keywords string `json:"keyword,omitempty"`

	/* Gets the creator. */
	Creator string `json:"creator,omitempty"`

	/* Gets the producer. */
	Producer string `json:"producer,omitempty"`

	/* Gets the title. */
	Title string `json:"title,omitempty"`

	/* Gets the collection of PageInformation. */
	Pages []PageInformation `json:"pages,omitempty"`

	/* Gets the form fields. */
	FormFields FormFieldInformation `json:"formFields,omitempty"`

	/* Gets the custom properties. */
	CustomProperties string `json:"customProperties,omitempty"`

	/* Gets the boolean representing xmp meta data. */
	XmpMetaData bool `json:"xmpMetaData,omitempty"`

	/* Gets the boolean, indicating whether the pdf is signed. */
	Signed bool `json:"signed,omitempty"`

	/* Gets the boolean, indicating whether the pdf is tagged. */
	Tagged bool `json:"tagged,omitempty"`
}
