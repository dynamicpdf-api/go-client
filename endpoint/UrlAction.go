package endpoint

// Represents information of a signature field.
type UrlAction struct {

	// Gets or Sets the name of a signature field.
	Url string `json:"url,omitempty"`
}

func NewUrlAction(url string) *UrlAction {
	ua := new(UrlAction)
	ua.Url = url
	return ua
}
