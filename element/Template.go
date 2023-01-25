package element

/**
 * Represents a document template. */
type Template struct {

	/** Gets or Sets the name of a signature field. */
	Id string `json:"id"`

	/** Gets or Sets the boolean, indicating the field signed or not. */
	Elements []ElementCollector `json:"elements"`
}

func NewTemplate() Template {
	var ep Template
	return ep
}
